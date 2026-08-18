# Domain Invariant — account

> File: `docs/spec/account/invariants.md`
> Status: draft
> Last updated: 2026-08-05

## Domain summary

`account` owns user identity, authentication (`email_password` +
`google`), session lifecycle (access/refresh tokens), MFA (TOTP +
backup codes), platform role assignment (`admin`, `kurator`), and the
account-scoped audit trail (`user_logs`). It does **not** own
notification delivery (`notification` domain) or organization
representation (`organization` domain), though one invariant below
(INV-account-10) has a mutation trigger point here that a later
domain must respect.

## Invariants

### INV-account-01: Auth identity uniqueness is per-provider, not global

- **Statement**: For any two rows in `auth_identities`, if
  `provider_type` is equal AND `identifier_hash` is equal, they must
  be the same row. Uniqueness is scoped to `(provider_type,
  identifier_hash)`, not to `identifier_hash` alone — the same email
  may appear once under `email_password` and once under `google`,
  each belonging to a different `user_id`, without violating this
  invariant.
- **Holds after operations**: register (Fitur 1), Google login/register
  (Fitur 1B), account linking / set-password (Fitur 4).
- **Verification**: DB-level — `ux_auth_identities_provider_identifier`
  unique index (already in ERD). Test: concurrent-insert race test
  asserting the second of two simultaneous registrations with the
  identical `(provider_type, identifier)` fails cleanly, not silently
  overwrites or duplicates.

### INV-account-02: Every user retains at least one auth identity at all times

- **Statement**: For every row in `users`, `COUNT(*) FROM
  auth_identities WHERE user_id = users.id` must be **≥ 1**, at every
  point in time after the user's `AuthIdentity` is first created.
- **Holds after operations**: unlink Google (Fitur 4) — this is the
  only operation that removes an `auth_identities` row. The unlink
  operation must be rejected (409) if it would bring the count to 0
  for that user.
- **Verification**: Test asserting unlink is rejected when it is the
  user's sole `auth_identities` row, and a concurrency test for the
  case where a user attempts to unlink the same identity twice
  simultaneously (must not race past the count check).

### INV-account-03: Refresh token rotation is single-use and one-way

- **Statement**: For any row in `refresh_tokens`, `replaced_by_id` can
  transition from `NULL` to a non-null value **at most once**. No two
  distinct rows may ever reference the same parent via
  `replaced_by_id`.
- **Holds after operations**: token refresh (Fitur 2, "Flow — Refresh").
- **Verification**: Concurrency test — two simultaneous refresh
  requests using the same refresh token; exactly one must succeed
  (atomic `WHERE replaced_by_id IS NULL AND revoked_at IS NULL` guard),
  the other must fail without creating a second child token.

### INV-account-04: Reuse of a rotated refresh token revokes its entire family

- **Statement**: If a `refresh_tokens` row with `replaced_by_id IS NOT
  NULL` is presented again at the refresh endpoint, then after that
  request, every row in `refresh_tokens` sharing the same `family_id`
  must have `revoked_at IS NOT NULL` (including rows that were already
  rotated further down the chain).
- **Holds after operations**: token refresh, specifically the reuse-
  detection branch (Fitur 2).
- **Verification**: Test — rotate a token twice (`A → B → C`), then
  replay `A`; assert `A`, `B`, and `C` all end up revoked, and that a
  subsequent refresh attempt using `C` (the last legitimately-issued
  token) is also rejected.

### INV-account-05: Successful password reset revokes all of the user's existing sessions

- **Statement**: After `AuthIdentity.credential_secret` is updated via
  the reset-password flow, every row in `refresh_tokens` for that
  `user_id` that had `revoked_at IS NULL` at the start of the request
  must have `revoked_at IS NOT NULL` by the end of it.
- **Holds after operations**: reset password (Fitur 2B).
- **Verification**: Test — user has 2+ active refresh tokens (e.g. two
  devices), completes password reset, assert both are revoked in the
  same transaction as the credential update (not a separate,
  best-effort follow-up step that could partially fail).

### INV-account-06: MFA backup codes are single-use, and only meaningful while MFA is enabled

- **Statement**: A row in `mfa_backup_codes` may transition `used_at`
  from `NULL` to a timestamp **at most once**. Additionally, a backup
  code is only accepted at login if `mfa_totp_secrets.enabled_at IS
  NOT NULL` for that `user_id` at the moment of verification —
  **implicit invalidation** (see decision below): disabling MFA does
  **not** delete or mark existing `mfa_backup_codes` rows, it makes
  them permanently unusable via this enabled-check, since the
  `enabled_at` toggle to `null` on disable is one-way per cycle (a
  fresh `enabled_at` after re-enrollment does **not** revive old
  backup codes — see INV-account-07 and the state machine below,
  which requires a full new set of 10 to be generated on every
  enrollment).
- **Decision — implicit invalidation, not hard-delete** **[RESOLVED —
  2026-08-05]**: chosen over deleting rows on disable, to avoid an
  extra write on every disable-MFA request; the trade-off (unused
  `mfa_backup_codes` rows accumulate indefinitely across disable/
  re-enable cycles, with no housekeeping) is accepted as a known,
  low-severity cost for a sandbox project — see "Knowingly accepted
  residual risk" in `docs/spec/account/threat-model.md` once written.
- **Holds after operations**: MFA enrollment (generates a fresh batch
  of 10), disable MFA, login via backup code.
- **Verification**: Test — enroll, disable, then attempt to use a
  backup code issued from the disabled enrollment; must fail even
  though the row's `used_at` is still `NULL`. Separate test: using the
  same backup code twice while MFA is enabled must fail on the second
  attempt.

### INV-account-07: MFA can never be "enabled" without a verified TOTP confirmation

- **Statement**: `mfa_totp_secrets.enabled_at` may only transition from
  `NULL` to non-null immediately after a TOTP code generated from that
  same `secret_encrypted` has been successfully verified in the same
  enrollment flow. There is no code path that sets `enabled_at` from a
  freshly-generated secret without an intervening successful
  verification.
- **Holds after operations**: MFA enrollment (Fitur 3).
- **Verification**: Test asserting enrollment cannot complete (i.e.
  `enabled_at` stays `NULL`) if the confirmation code check fails or
  is skipped.

### INV-account-08: Email-verification and password-reset tokens are single-use and time-bound

- **Statement**: A row in `auth_tokens` may only be successfully
  redeemed (email verified / password updated) if, at redemption time,
  `used_at IS NULL AND revoked_at IS NULL AND expires_at > now()`.
  Once redeemed, `used_at` is set and no further redemption of the
  same row may succeed.
- **Holds after operations**: email verification (Fitur 1), password
  reset (Fitur 2B), and any resend flow that sets `revoked_at` on a
  superseded token.
- **Verification**: Test — double-submit the same reset link
  concurrently; exactly one request succeeds, guarded by `WHERE
  used_at IS NULL AND expires_at > now()` at the `UPDATE`.

### INV-account-09: Admin role is mutually exclusive with Kurator role

- **Statement**: For any `user_id`, `user_roles` may not simultaneously
  contain a row with `role = 'admin'` and a row with `role = 'kurator'`.
- **Holds after operations**: role assignment / revoke (Fitur 5) — both
  directions (assigning `kurator` to an existing Admin, and assigning
  `admin` to an existing Kurator, must both be rejected).
- **Verification**: Test covering both assignment directions, plus a
  concurrency test for two simultaneous assignment requests
  (`admin` and `kurator`) targeting the same user.

### INV-account-10: Admin role is mutually exclusive with being an Organisasi Representative

- **Statement**: For any `user_id`, `user_roles` may not contain a row
  with `role = 'admin'` while that same `user_id` also has any row in
  `organization_representatives` (any `level`). This is checked from
  **both** directions: assigning `admin` to a user must be rejected if
  they currently represent any organization, and (once the `organization`
  domain is built) inviting/creating a representative row for a user
  who currently holds `admin` must also be rejected.
- **Holds after operations**: role assignment (Fitur 5, `account`
  domain — enforceable and testable today), and representative invite
  (`kencleng-phase1-detail.md` §"Manage Representative", `organization`
  domain — not yet built).
- **Verification**: Today (account-only): test that assigning `admin`
  to a user with an existing `organization_representatives` row is
  rejected with 403/409. **Deferred**: the reverse direction (inviting
  a representative who is already Admin) cannot be tested until
  `organization` endpoints exist — see the reference note below.
- **Cross-domain note**: see "Reference for `organization` domain"
  section below — this invariant is declared here in full, per the
  cross-domain ownership rule (`kencleng-agentic-workflow.md` §5.1),
  because its primary mutation trigger point (role assignment) lives
  in `account`. `organization/invariants.md` must **reference** this
  entry, not redefine it, when that domain is worked on.

### INV-account-11: `user_logs` is append-only

- **Statement**: No row in `user_logs`, once inserted, may ever be
  updated or deleted — by any actor, including Admin, and including
  the application's own DB role.
- **Holds after operations**: every sensitive action listed in
  `kencleng-phase0-detail.md` Fitur 9 that targets a `user_id` (role
  assign/revoke, MFA enable/disable, account linking, reveal of a
  user's own PII by Admin/Kurator).
- **Verification**: DB-level — `REVOKE UPDATE, DELETE ON user_logs
  FROM kencleng_app` (already in ERD). Test: attempt an `UPDATE`/
  `DELETE` against `user_logs` using the app's DB role and assert it
  fails at the privilege level, not just "the app code doesn't do it."

### INV-account-12: Unlinking a non-`email_password` identity requires the remaining identity to be verified

- **Statement**: For an unlink operation removing `AuthIdentity` A
  from `user_id` X, there must exist at least one **other**
  `AuthIdentity` B for the same `user_id` where `B.verified_at IS NOT
  NULL`, at the time of the unlink. This is **stricter** than
  INV-account-02 (which only requires `COUNT(*) >= 1`, regardless of
  verification status) — added specifically as an unlink precondition,
  2026-08-05, per `docs/spec/account/features/05-account-linking.md`.
- **Rationale**: an `email_password` identity with `verified_at IS
  NULL` is still usable for login (per
  `docs/spec/account/features/03-login-session-management.md`), so
  INV-account-02 alone would let a user unlink Google right after
  typing an email they don't actually control — leaving them
  functionally locked out of the one channel (that email inbox) that
  would ever let them recover the account or receive
  security-relevant notifications, even though they could still log
  in with the password they just set.
- **Holds after operations**: unlink (`POST
  /account/security/google/unlink`).
- **Verification**: test that unlink is rejected (with a distinct
  error from the "no other identity at all" case) when the remaining
  identity exists but `verified_at IS NULL`.

### INV-account-13: At least one Admin always exists in the system

- **Statement**: `COUNT(*) FROM user_roles WHERE role = 'admin'` must
  never drop to 0 as a result of a role-revocation request.
- **Rationale**: Fitur 5's only Admin-bootstrap mechanism is a
  one-time manual seed script run at setup/deploy — there is no
  runtime/API path to create the first Admin. If the last Admin's role
  were revoked (self-revoke or by another Admin, accidentally or via
  a compromised session), the system would have zero Admins with no
  recovery path short of manual DB/migration intervention. Same
  protective spirit as INV-account-02 (never let an action produce an
  unrecoverable lockout state), applied at the system level instead of
  per-user. Resolved 2026-08-05, per
  `docs/spec/account/features/08-role-assignment.md` Assumption B —
  accepted as worth the small added cost (one `COUNT(*)` check, only
  on `admin`-role revocation) given the alternative is an
  unrecoverable-via-API state.
- **Holds after operations**: role revoke (`DELETE
  /admin/users/{userId}/roles?role=admin`).
- **Verification**: test that revoking `admin` from the sole remaining
  Admin is rejected (`409`), and a concurrency test — two simultaneous
  revoke requests when exactly 2 Admins exist, both targeting
  different Admins, must not be allowed to both succeed and leave 0
  (the guard must re-check the count atomically per request, not
  against a stale pre-fetched count).

### INV-account-14: Revoking the Kurator role requires no pending curation assignments

- **Statement**: For a role-revoke operation removing `role='kurator'`
  from `user_id` X, there must be no row in
  `organization_curation_assignments` where `kurator_id = X AND
  decision = 'pending'`, at the time of the revoke.
- **Rationale**: `kencleng-phase0-detail.md` Fitur 5 itself flagged
  this as an open question ("Kurator yang di-demote — assignment
  kurasi yang sedang `pending` miliknya diapakan?"), never resolved.
  Resolved 2026-08-05, per
  `docs/spec/account/features/08-role-assignment.md` Assumption C —
  **block the demote** rather than leaving a `pending` assignment
  orphaned to a user who's no longer a Kurator (which the
  `organization` domain would otherwise have to defensively handle
  everywhere it renders or processes the curation queue). Chosen over
  auto-reassignment specifically to avoid `account` needing to know
  anything about `organization`'s assignment logic — this guard only
  needs a **read** against a table `organization` owns, not a write
  into it, preserving domain boundaries
  (`kencleng-agentic-workflow.md` already rejected designs that
  couple domain write-logic together).
- **Holds after operations**: role revoke (`DELETE
  /admin/users/{userId}/roles?role=kurator`).
- **Cross-domain note**: this is the **reverse** of INV-account-10's
  situation — here, `account` reads a table owned by `organization`
  (`organization_curation_assignments`), rather than `organization`
  needing to reference an `account`-owned rule. The table's structure
  already exists in `docs/project/kencleng-erd.md` (the ERD is
  holistic across all domains even though each domain's spec is
  written incrementally), so this read dependency is implementable
  today even though `organization`'s own spec docs don't exist yet.
  When `docs/spec/organization/` is written, its curation-related
  feature specs should reference this invariant rather than
  re-deciding the demoted-Kurator question independently.
- **Verification**: test that revoking `kurator` from a user with a
  `pending` curation assignment is rejected (`409`); test that it
  succeeds once that assignment's `decision` moves to
  `approved`/`rejected`.

## State machines

### `auth_identities.verified_at`

```
null -> verified (timestamp)
```

One-way. `email_password` identities start at `null`, move to
verified via the email-verification token flow (or by admin action —
none defined in v1). `google` identities are created **already
verified** (`verified_at = now()` at insert), never passing through
`null`.

### `refresh_tokens` — two independent flags, not a single linear state

`refresh_tokens` rows are best modeled as **two independent boolean
transitions** rather than one enum, because either can happen without
the other, and revocation can retroactively apply to already-rotated
tokens:

- `replaced_by_id`: `NULL -> <child token id>` (at most once — see
  INV-account-03)
- `revoked_at`: `NULL -> timestamp` (at most once), settable either
  directly (logout) or as a side effect of reuse detection cascading
  across the whole `family_id` (see INV-account-04) — meaning a token
  that was already rotated (`replaced_by_id` set) can still later have
  `revoked_at` set retroactively.

Typical happy-path sequence for one token:
```
issued (active) -> used for refresh -> replaced_by_id set (rotated)
issued (active) -> logout -> revoked_at set directly
any state -> reuse of a rotated token detected -> revoked_at set on
             every row in the family_id, regardless of prior state
```

A token is considered usable for authentication only while
`revoked_at IS NULL AND replaced_by_id IS NULL` (this is exactly the
`ix_refresh_tokens_active` partial index condition in the ERD).

### `mfa_totp_secrets.enabled_at`

```
null -> enabled (timestamp) -> null (disable) -> enabled (timestamp) -> ...
```

Toggles freely, but every `null -> enabled` transition requires a full
enrollment cycle (new secret generated, verified, 10 new backup codes
issued — see INV-account-07). The row itself is never deleted
(upserted in place); `secret_encrypted` is overwritten on each new
enrollment.

### `mfa_backup_codes`

Not a per-row state machine beyond `used_at` (`NULL -> timestamp`,
single-use — INV-account-06), but **functional validity** additionally
depends on the owning user's current `mfa_totp_secrets.enabled_at`,
which is external to this table (see INV-account-06's implicit-
invalidation decision).

## Reference for `organization` domain

When `docs/spec/organization/invariants.md` is written, it must include
a reference to **INV-account-10** (declared above, not redefined) for
the representative-invite / representative-promote endpoints
(`kencleng-phase1-detail.md` §"Manage Representative"). Specifically,
the `organization` domain is responsible for:

- Enforcing and testing the direction not yet testable from `account`
  alone: rejecting an invite/promote-to-`owner` action that targets a
  user who currently holds `role = 'admin'`.
- Not re-declaring the invariant statement itself — link back to
  `docs/spec/account/invariants.md#inv-account-10`.

This is the reverse situation from the `donation`/`campaign` precedent
in `kencleng-agentic-workflow.md` §5.1 (there, the referenced domain
was already built by the time the reference was written). Here,
`organization` is built **after** `account` in the domain order (§3.3),
so this forward reference is unavoidable and is noted explicitly
rather than silently assumed.

**INV-account-14** is the mirror-image case: `account` reads a table
`organization` owns (`organization_curation_assignments`). When
`docs/spec/organization/`'s curation feature specs are written, they
should reference INV-account-14 for the demoted-Kurator question
rather than re-deciding it — and should confirm the `pending`→
`approved`/`rejected` transition is the only thing that unblocks a
previously-blocked Kurator role revoke (i.e. `organization`'s curation
decision flow is what indirectly resolves the `account`-side block,
not anything `account` does on its own).

## References

- Related ERD: `docs/project/kencleng-erd.md` §1 "Identity & Auth"
  (`users`, `user_roles`, `auth_identities`, `refresh_tokens`,
  `auth_tokens`, `mfa_totp_secrets`, `mfa_backup_codes`,
  `login_attempts`, `user_logs`)
- Related business process: `docs/project/kencleng-phase0-detail.md`
  Fitur 1 (Registrasi & Verifikasi Email), Fitur 1B (Login/Register
  Google), Fitur 2 (Login & Session Management), Fitur 2B (Forgot &
  Reset Password), Fitur 3 (TOTP MFA), Fitur 4 (Account Linking),
  Fitur 5 (Role Bootstrapping & Assignment), Fitur 9 (Audit Log)
- Related actors/roles: `docs/project/kencleng-actors-entities.md`