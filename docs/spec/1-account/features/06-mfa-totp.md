# Feature Spec — MFA TOTP

> File: `docs/spec/account/features/06-mfa-totp.md`
> Status: draft
> Risk tier: 1 (with a Tier 0 fenced sub-area)
> Domain: account

## Endpoint

- `POST /account/security/mfa/enroll`
- `POST /account/security/mfa/enroll/confirm`
- `POST /account/security/mfa/disable`

Login-time TOTP/backup-code verification (`POST /auth/login/mfa`) is
specced in `03-login-session-management.md`, including that
endpoint's own MFA-stage lockout — not duplicated here.

## Acceptance criteria

### `POST /account/security/mfa/enroll`

- Given the authenticated user does **not** currently have MFA
  enabled (`mfa_totp_secrets.enabled_at IS NULL`), When called, Then a
  new TOTP secret is generated, encrypted at rest (same AES-GCM
  encryption-at-rest mechanism used for other sensitive fields —
  `secret_encrypted`, no new key-management scheme introduced), stored
  with `enabled_at` still `NULL`, and the response includes an
  `otpauth://` URI for the client to render as a QR code.
- Given the authenticated user calls this endpoint again **before**
  confirming (e.g. abandoned and restarted), When called, Then the
  pending (unconfirmed) secret is simply overwritten with a fresh one
  — safe, since `enabled_at` stays `NULL` throughout (INV-account-07
  guarantees no half-enabled state regardless of how many times this
  is called).
- Given the authenticated user **already has MFA enabled**
  (`enabled_at IS NOT NULL`), When called, Then `409` — re-enrollment
  while already active is rejected; this endpoint is not how backup
  codes get regenerated (that's disable → enable, an already-resolved
  rule) or how the secret gets rotated while staying enabled. Without
  this guard, a stray re-enroll call could silently replace
  `secret_encrypted` out from under an already-`enabled_at`-set
  account, breaking the user's existing authenticator app while the
  system still believes MFA is active.

### `POST /account/security/mfa/enroll/confirm`

- Given a pending (unconfirmed) secret exists and the submitted TOTP
  code validates against it, When called, Then `mfa_totp_secrets.
  enabled_at` is set to now (INV-account-07), exactly 10 backup codes
  are generated (hashed at rest, shown in the response **once** and
  never retrievable again), response `200`.
- Given the submitted code doesn't validate, When called, Then `422`,
  no state change — `enabled_at` stays `NULL`, and the pending secret
  is **not** discarded (the user can retry without re-scanning the
  QR code).
- Given no pending secret exists (enroll was never called, or the
  session/user mismatch), When called, Then `422` (treated the same
  as an invalid code — no distinguishing response needed, this isn't
  an enumeration-sensitive endpoint since it's authenticated and
  self-targeting only).

### `POST /account/security/mfa/disable`

- Given the caller is an `email_password` user and submits the
  correct current `password`, When called, Then `mfa_totp_secrets.
  enabled_at` is set to `NULL`, response `200` — existing
  `mfa_backup_codes` rows are **not** deleted (INV-account-06's
  implicit-invalidation decision — they become permanently unusable
  via the `enabled_at IS NOT NULL` check at verification time, not via
  a cleanup step here).
- Given the caller is an `email_password` user and submits an
  incorrect `password`, When called, Then `401`, no state change.
- Given the caller is Google-only, When called with no body, Then the
  short-lived server-side re-auth marker (set by `GET
  /auth/google/redirect?intent=reauth` →
  `02-google-oauth-login-register.md` Assumption A, proposed 5 min
  TTL) must be present and unexpired for this session; if valid, same
  outcome as above; the marker is consumed (invalidated) on use so it
  can't be replayed for a second disable call. If missing/expired,
  `401`.

## Error cases

| Condition | Expected response |
|---|---|
| Enroll called while MFA already enabled | `409` |
| Confirm: wrong TOTP code, or no pending secret | `422` |
| Disable: wrong password (`email_password` users) | `401` |
| Disable: missing/expired re-auth marker (Google-only users) | `401` |

## Applicable invariants

- `docs/spec/account/invariants.md#inv-account-06` — backup codes
  single-use + implicit invalidation tied to MFA enabled state; this
  is the endpoint pair (confirm generates them, disable is what makes
  old ones functionally dead).
- `docs/spec/account/invariants.md#inv-account-07` — MFA can never be
  "enabled" without a verified TOTP confirmation; this is the endpoint
  that directly implements it.

## Threat breakdown

Derived from `docs/spec/account/threat-model.md` component 5:

| Threat | Mitigation at this endpoint's level | Test that proves it |
|---|---|---|
| MFA disabled by an attacker with a hijacked live session | Re-authentication required (password, or the Google reauth marker) — same reasoning as `05-account-linking.md`'s unlink re-auth requirement | `TestMfaDisable_RequiresReauth_EmailPassword`, `TestMfaDisable_RequiresReauth_GoogleOnly` |
| Enrollment "active" with an unverified secret | INV-account-07 — no code path sets `enabled_at` without a successful confirm | `TestMfaEnroll_NoHalfEnabledState` |
| Stray re-enroll silently breaking an already-active MFA setup | **Found while drafting** — `409` guard on enroll when already enabled (see acceptance criteria) | `TestMfaEnroll_RejectsWhenAlreadyEnabled` |
| Backup codes usable after MFA disabled | INV-account-06 implicit invalidation, verified at the login-time check (`03-login-session-management.md`), not here | `TestMfaDisable_OldBackupCodesUnusable` (lives in `03`'s test suite, referenced here) |

## Risk tier & rationale

**Tier 1**, with a **Tier 0 fenced sub-area**: TOTP secret generation
and verification (`pquerna/otp` usage, encryption of `secret_encrypted`
at rest). Matches the Tier 0 examples in `kencleng-agentic-workflow.md`
§13.2 verbatim ("JWT/TOTP") — same reasoning as the refresh-token core
in `03-login-session-management.md`. The rest (request handling,
backup code generation/hashing, re-auth checks) is ordinary Tier 1
agent-generated work.

## Assumptions / open questions

**A. Resolved while drafting — enroll-while-already-enabled guard.**
Not explicitly stated in `kencleng-phase0-detail.md`, but directly
implied by the already-resolved "regenerate only via disable→enable"
rule: this endpoint must reject re-enrollment attempts while MFA is
already active (`409`), rather than silently overwriting a live
secret. High-confidence derivation, not flagged for further
confirmation.

**B. Audit log scope extended to include a user notification,
consistent with the `05-account-linking.md` precedent.** Fitur 9
explicitly lists "MFA enable/disable" as in the audit-log action list
(unlike the account-linking case, this one didn't need inferring).
Extending it to also trigger a user-facing notification ("MFA berhasil
diaktifkan di akunmu" / "MFA berhasil dinonaktifkan") follows the same
reasoning Anhar gave for `05` — worth applying consistently rather
than re-deciding per feature. Same cross-domain dependency on the
not-yet-built `notification` domain applies here too (see `05`'s
Audit log section for the full reasoning). Flagged as a **consistent
extension of an established pattern**, not a fresh open question — but
easy to revert here specifically if that reasoning shouldn't have
generalized this far.

## Audit log entry?

**Yes** — explicitly in Fitur 9's scope ("MFA enable/disable"). Write
a `user_logs` entry on successful enroll/confirm (MFA enabled) and on
successful disable. Per Assumption B, also triggers a user
notification — cross-domain dependency on `notification`, same as
`05-account-linking.md`.

## References

- `docs/project/kencleng-phase0-detail.md` Fitur 3
- `docs/spec/account/features/02-google-oauth-login-register.md` —
  `reauth` intent mechanics for Google-only disable
- `docs/spec/account/features/03-login-session-management.md` —
  login-time TOTP/backup-code verification and MFA-stage lockout
- `docs/spec/account/invariants.md` INV-account-06, 07
- `docs/spec/account/threat-model.md` component 5
- `api/openapi.yaml` — `POST /account/security/mfa/enroll`,
  `POST /account/security/mfa/enroll/confirm`,
  `POST /account/security/mfa/disable`