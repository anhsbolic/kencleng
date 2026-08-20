# Feature Spec — Account Linking

> File: `docs/spec/account/features/05-account-linking.md`
> Status: draft — all open items resolved, ready for human review
> Risk tier: 1
> Domain: account

## Endpoint

- `POST /account/security/google/unlink`
- `POST /account/security/set-password`

**Structural note**: linking Google *to* an existing `email_password`
account happens through the shared OAuth endpoints (`intent=link`),
already fully specced in
`docs/spec/account/features/02-google-oauth-login-register.md` — not
duplicated here. This file covers the two remaining linking-related
actions: removing a linked Google identity, and adding an
`email_password` identity to a Google-only account (the reverse
direction — see the redesign below, resolved 2026-08-05).

## Design note — `set-password` branches into two behaviors depending
on the caller's current identities

**Branch 1 — caller has no `email_password` identity yet
(Google-only)**: adds a **new** identity. The endpoint originally
assumed the new identity always uses `User.primary_email` (the email
Google already verified), letting it skip email verification entirely
(`verified_at=now` immediately). That assumption breaks once the
account might unlink Google in favor of a **different** email (e.g.
moving from a personal `@gmail.com` to a work `@company-mail.com`) —
a real, expected use case, not an edge case. Resolved: the request
takes an explicit `email` field (not implicitly `primary_email`), and
the new identity **is not instantly verified** — it goes through the
same email verification flow as registration, reusing `POST
/auth/verify-email` unchanged. This also means unlink needs a
stricter precondition than before — see INV-account-12, added
specifically for this.

**Branch 2 — caller already has an `email_password` identity — resolved
2026-08-05.** This is a genuine change-password action, not "add
another identity." Conflating it with Branch 1 would have let a user
end up with multiple `email_password` identities under different
emails (never intended), and would have forced an unnecessary
re-verification step on an email that's already verified. Resolved:
this branch **updates `credential_secret` on the existing identity in
place** — the email itself is not submitted and does not change — and
takes effect **immediately** (no verification needed, since the email
was already verified). Requires **`current_password`** confirmation in
the request body, on the same reasoning as the unlink and MFA-disable
re-auth requirements elsewhere in this domain: an attacker holding a
stolen-but-still-valid access token (15 min window) must not be able
to silently take over the account by changing its password. After a
successful change, **all existing refresh tokens for that user are
revoked** (same INV-account-05 pattern as `04-forgot-reset-password.md`,
reused here rather than re-derived).

**Resulting 3-step flow for Branch 1** (per Anhar's decision — this is
the flow that motivated the redesign; Branch 2 is a single-step
action, no flow needed):
1. `POST /account/security/set-password` — submit new `email` +
   `password`, creates an unverified `email_password` `AuthIdentity`.
2. `POST /auth/verify-email` — same endpoint as registration
   (`01-register-email-verification.md`), using the token from the
   verification email sent in step 1.
3. `POST /account/security/google/unlink` — now succeeds, since the
   remaining identity is verified (INV-account-12).

## Acceptance criteria

### `POST /account/security/set-password`

**Branch 1 — caller has no `email_password` identity yet:**

- Given the authenticated user submits a new `email` not currently
  claimed by any `email_password` identity, and a password that
  passes the length policy (≥8 chars) and isn't in the breach-list
  (or the breach-check API is unreachable — fail-open, same as
  registration), When submitted, Then a new `AuthIdentity` is created
  (`provider_type=email_password`, `identifier=<submitted email>`,
  `verified_at=null`), a single-use verification token is generated
  and a verification email sent (identical mechanics to Fitur 1), and
  the API responds `202` generic (see the enumeration note below) —
  **no re-authentication required**, since adding a new identity
  strengthens the account rather than weakening it (unchanged from
  the original resolved rule).
- Given the submitted `email` is already claimed by an
  `email_password` identity belonging to **any** user (including a
  different user), When submitted, Then the API responds identically
  (`202` generic) — **this branch follows the same anti-enumeration
  pattern as `/auth/register`** (`01-register-email-verification.md`),
  since it's creating a new `email_password` identity subject to the
  exact same INV-account-01 uniqueness guard, and the same
  enumeration risk applies even though the caller is authenticated. A
  distinct nudge email is sent instead of a verification email in the
  conflict case — no new identity or token created.
- Given the new password fails the length policy or is found in the
  breach-list, When submitted, Then `422` — this check happens before
  any enumeration-sensitive branching, same ordering as registration.

**Branch 2 — caller already has an `email_password` identity
(change-password) — resolved 2026-08-05:**

- Given the authenticated user submits the correct `current_password`
  and a `password` that passes the length policy and isn't in the
  breach-list (or breach-check unreachable — fail-open), When
  submitted, Then, **in one transaction**: the existing identity's
  `credential_secret` is updated (the `identifier`/email is
  untouched) and **every** existing refresh token for that `user_id`
  is revoked (same INV-account-05 pattern as
  `04-forgot-reset-password.md`) — response `200`. Takes effect
  immediately, no verification step (the email is already verified).
- Given an incorrect `current_password`, When submitted, Then `401`,
  no state change — this is the re-auth guard, closing the
  hijacked-session-changes-your-password gap (see threat breakdown).
- Given the new password fails the length policy or is found in the
  breach-list, When submitted, Then `422`, no state change — same
  ordering requirement as `04-forgot-reset-password.md` Assumption B
  (validate the new password *before* touching anything, so a fixable
  input mistake doesn't have side effects).

**Branch selection**: server-side, based on whether the authenticated
`user_id` currently has an `email_password` `AuthIdentity` — not a
client-supplied flag. Request schema is conditionally shaped like
`MfaDisableRequest` already is in `06-mfa-totp.md` (`email`+`password`
for Branch 1, `current_password`+`password` for Branch 2) — same
established pattern, not a new one.

### `POST /account/security/google/unlink`

- Given the authenticated user has **at least one other** `AuthIdentity`
  with `verified_at IS NOT NULL` (INV-account-12 — stricter than
  simply having another identity at all), When unlink is requested,
  Then the `google` `AuthIdentity` row is hard-deleted (no soft-delete
  column exists on this table), **re-authentication is required first**
  (see below), response `200` `UnlinkGoogleResponse`.
- Given the authenticated user has **no other** `AuthIdentity` at all
  (INV-account-02's original guarded case), When unlink is requested,
  Then `409`, message: "Google adalah satu-satunya metode login Anda.
  Atur email dan password dulu sebelum melepas tautan." (unchanged
  from the original rule).
- Given the authenticated user has another `AuthIdentity`, but it is
  **not yet verified** (mid-way through the 3-step flow above — new
  case, INV-account-12), When unlink is requested, Then `409`, a
  **distinct** message from the case above: "Kamu sudah atur email dan
  password, tapi belum diverifikasi. Verifikasi email kamu dulu
  sebelum bisa melepas tautan Google." — distinguishing the two `409`
  cases matters so the user knows whether they need to *start*
  set-password or just *finish* verifying an email they already
  started.
- **Re-authentication requirement** (resolved 2026-08-05, reversing
  the earlier draft's "no re-auth" default): unlink requires the
  caller to confirm their current `email_password` password in the
  request body (`password` field) — this is always possible by the
  time unlink can succeed at all, since INV-account-12 guarantees a
  verified `email_password` identity exists at that point. No
  Google-reauth branch is needed here (unlike MFA-disable), because by
  definition the identity being removed is the Google one. Given a
  wrong password, When submitted, Then `401`, no state change.
- Given the count-check (INV-account-02/12) and the delete/re-auth
  are all evaluated concurrently by two overlapping requests, Then
  the whole check-then-delete sequence must be atomic (not racy) — see
  the threat breakdown.

## Error cases

| Condition | Expected response |
|---|---|
| Set-password (Branch 1): email already claimed (any case) | `202` generic — no distinguishing status (anti-enumeration) |
| Set-password (Branch 1): password fails length/breach-list policy | `422` |
| Set-password (Branch 2): wrong `current_password` | `401` |
| Set-password (Branch 2): new password fails length/breach-list policy | `422` |
| Unlink: `google` is the only identity at all | `409`, "set up email+password first" message |
| Unlink: another identity exists but is unverified | `409`, distinct "verify your email first" message |
| Unlink: wrong password on re-auth | `401` |

## Applicable invariants

- `docs/spec/account/invariants.md#inv-account-01` — uniqueness per
  provider; `set-password` creates a new `email_password` identity
  subject to the same uniqueness guard and the same anti-enumeration
  response pattern as registration.
- `docs/spec/account/invariants.md#inv-account-02` — minimum one auth
  identity at all times (necessary but no longer sufficient on its
  own for unlink — see INV-account-12).
- `docs/spec/account/invariants.md#inv-account-08` — the verification
  token issued by `set-password` is single-use and time-bound, same
  as the registration token (reuses the same `auth_tokens` mechanism
  and the same `/auth/verify-email` endpoint).
- `docs/spec/account/invariants.md#inv-account-05` — successful
  password change (Branch 2) revokes all of the user's existing
  sessions, same pattern as reset-password, applied here too.
- `docs/spec/account/invariants.md#inv-account-12` — **new**, added
  specifically for this feature: unlink requires the remaining
  identity to be verified, not merely present.

## Threat breakdown

Derived from `docs/spec/account/threat-model.md` components 4 and 5,
updated for the redesign:

| Threat | Mitigation at this endpoint's level | Test that proves it |
|---|---|---|
| Unlink bypassing INV-account-12 via a race (verify-then-unlink racing against a concurrent unlink) | Atomic check-then-delete in one guarded operation, not a check followed by a separate delete | `TestUnlinkGoogle_ConcurrentRequests_GuardHolds` |
| Set-password identifier conflict / enumeration | Same unique-index-backed guard + generic response as registration (INV-account-01) | `TestSetPassword_ConcurrentDuplicateEmail_Race`, `TestSetPassword_GenericResponse_AllBranches` |
| Weak/breached password accepted via set-password | Same length policy + HaveIBeenPwned check as registration, fail-open on API outage | `TestSetPassword_PasswordPolicy`, `TestSetPassword_BreachCheck_FailOpen` |
| Verification token replay/double-submit (set-password's token) | Same guard as `01`'s `INV-account-08` test — reused mechanism, reused test pattern | `TestVerifyEmail_TokenSingleUse_Concurrent` (already covers this — same endpoint) |
| Hijacked-session unlink (attacker with a stolen access token removes the victim's Google identity) | **Resolved 2026-08-05** — re-authentication (password confirmation) now required, closing this gap | `TestUnlinkGoogle_RequiresReauth`, `TestUnlinkGoogle_WrongPassword_Rejected` |
| User locked into an email they don't control (typo, or someone else's inbox) after unlinking | **Resolved via INV-account-12** — unlink blocked until the new identity is verified, i.e. inbox ownership proven | `TestUnlinkGoogle_RejectsUnverifiedRemainingIdentity` |
| Hijacked-session password change (Branch 2) — attacker with a stolen access token silently changes the victim's password | **Resolved 2026-08-05** — `current_password` confirmation required, plus all sessions revoked after a successful change (so even a successful hijack loses its own access immediately) | `TestSetPassword_Branch2_RequiresCurrentPassword`, `TestSetPassword_Branch2_AllSessionsRevoked` |

## Risk tier & rationale

**Tier 1** — INV-account-01, INV-account-02, and the new
INV-account-12 all require concurrency-safe guards with race tests
(project goal #2). No Tier 0 sub-area (no JWT/TOTP core logic touched
by either endpoint).

## Assumptions / open questions

All open items are now resolved:

- **Change-password while authenticated** (originally Assumption A):
  resolved as **Branch 2** above — a real change-password action on
  the existing identity, with `current_password` confirmation and
  force-logout-all-sessions, not a new identity. Caught during
  drafting (`07-account-profile.md`'s review of the `User` schema
  surfaced that the original single-branch design would have let a
  user accumulate multiple `email_password` identities under
  different emails, which was never intended). **Naming note, not yet
  decided**: the endpoint is still called `set-password` in
  `openapi.yaml`, but it now does two meaningfully different things
  depending on caller state — worth reconsidering the name (or
  splitting into two endpoints) during the `openapi.yaml` follow-up
  pass. Not blocking.
- **Unlink re-authentication** (originally Assumption B): resolved —
  yes, required, via password confirmation (see acceptance criteria
  above).

## Audit log entry?

**Yes, for both**, resolved 2026-08-05 — write a `user_logs` entry for
both `set-password` (on successful verification, i.e. when the
identity actually becomes active — not at the initial unverified-
creation step) and `unlink` (on successful removal). **Additionally**,
per this decision, both actions should trigger a user-facing
notification (e.g. "Metode login baru berhasil ditambahkan ke akunmu,"
"Google berhasil dilepas dari akunmu") — this is a **cross-domain
dependency on the `notification` domain**, which is built directly
after `account` in the domain order
(`kencleng-agentic-workflow.md` §3.3: `account → notification → ...`).
This feature spec can only reference that dependency now, not fully
specify it — the concrete notification-sending mechanism/API belongs
to `docs/spec/notification/` once that domain's spec is written. This
mirrors how `docs/spec/account/invariants.md` INV-account-10 already
handles a forward reference to the not-yet-built `organization`
domain.

## References

- `docs/project/kencleng-phase0-detail.md` Fitur 4
- `docs/spec/account/features/01-register-email-verification.md` —
  reused verification mechanics
- `docs/spec/account/features/02-google-oauth-login-register.md` —
  OAuth mechanics for the `link` intent direction
- `docs/spec/account/features/04-forgot-reset-password.md` — the
  other password-change entry point, same force-logout-after pattern
- `docs/spec/account/invariants.md` INV-account-01, 02, 05, 08, 12
- `docs/spec/account/threat-model.md` components 4, 5
- `api/openapi.yaml` — `POST /account/security/google/unlink`,
  `POST /account/security/set-password` (both need a schema update —
  flagged as a follow-up, alongside the other pending `openapi.yaml`
  changes from earlier feature specs)