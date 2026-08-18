# Feature Spec — Forgot & Reset Password

> File: `docs/spec/account/features/04-forgot-reset-password.md`
> Status: draft
> Risk tier: 1
> Domain: account

## Endpoint

- `POST /auth/forgot-password`
- `POST /auth/reset-password`

## Acceptance criteria

### `POST /auth/forgot-password`

- Given an email matching an `email_password` `AuthIdentity`, When
  requested, Then a new single-use reset token is generated
  (`auth_tokens`, `purpose=password_reset`, expires in 1h), a reset
  email is sent, and the API responds `202` generic.
- Given an email matching only a `google` `AuthIdentity` (no
  `email_password` identity exists), When requested, Then **no token
  is created**, a distinct notice email is sent instead ("this account
  uses Google login"), and the API response is identical (`202`
  generic) to the success case.
- Given an email that doesn't match any account, When requested, Then
  no token/email, and the API response is identical (`202` generic) —
  anti-enumeration, per the resolved rule already in `openapi.yaml`.
- Given repeated `forgot-password` requests for the same email in
  quick succession, When requested, Then each request independently
  issues its own new token; previously-issued, still-unexpired tokens
  are **not** proactively revoked (see Assumption A, resolved).
- Rate limit: stricter `/auth/*` limit applies.

### `POST /auth/reset-password`

- Given a valid, unexpired, unused token and a new password that
  passes the length policy (≥8 chars) and isn't in the breach-list
  (or the breach-check API is unreachable — fail-open, per Fitur 1),
  When submitted, Then **in one transaction**:
  `AuthIdentity.credential_secret` is updated, the token's `used_at`
  is set, and **every** existing refresh token for that `user_id` is
  revoked (INV-account-05) — response `200`.
- Given the new password fails the length policy or is found in the
  breach-list, When submitted, Then `422`, and — critically — **the
  token is NOT consumed** (`used_at` stays `NULL`): password
  validation must happen **before** the atomic token-consuming update,
  so the user can retry with the same reset link rather than being
  forced to request a new one over a fixable input mistake (see
  Assumption B).
- Given an expired token, When submitted, Then `410`, no state change.
- Given a token that doesn't exist or was already used, When
  submitted, Then `404`, no state change.
- Given the same valid token submitted twice concurrently
  (double-submit), Then exactly one request succeeds (guarded `UPDATE
  ... WHERE used_at IS NULL AND expires_at > now()`, per
  INV-account-08); the other gets `404`.

## Error cases

| Condition | Expected response |
|---|---|
| New password fails length policy or found in breach-list | `422` — token remains unused, can retry |
| Reset token expired | `410` |
| Reset token not found / already used | `404` |
| Too many requests | `429` |

## Applicable invariants

- `docs/spec/account/invariants.md#inv-account-05` — successful
  password reset revokes all of the user's existing sessions, in the
  same transaction as the credential update.
- `docs/spec/account/invariants.md#inv-account-08` — `auth_tokens`
  single-use and time-bound; applies to the reset token.

## Threat breakdown

Derived from `docs/spec/account/threat-model.md`, component 3:

| Threat | Mitigation at this endpoint's level | Test that proves it |
|---|---|---|
| Attacker triggers a reset for a victim's email | Only the inbox owner can act on the link; API response never distinguishes success from no-match | `TestForgotPassword_GenericResponse_AllBranches` |
| Reset token guessing/tampering | Random, single-use, hashed, 1h expiry, guarded `UPDATE` | `TestResetPassword_TokenSingleUse_Concurrent` |
| Reset token replay after already used | Same guarded `UPDATE ... WHERE used_at IS NULL` | `TestResetPassword_TokenSingleUse_Concurrent` (same test covers both) |
| User enumeration via forgot-password response | Resolved — identical `202` generic response for all three branches (registered/unregistered/Google-only) | `TestForgotPassword_GenericResponse_AllBranches` |
| Reset-email flood via repeated requests | Stricter `/auth/*` rate limit | `TestForgotPassword_RateLimited` |
| Stale session hijack after a suspected-compromised reset | INV-account-05 — all sessions revoked atomically with the credential update, not as a best-effort follow-up step | `TestResetPassword_AllSessionsRevoked_Atomic` |
| Weak/breached password accepted on reset | Same length policy + HaveIBeenPwned check as registration, fail-open on API outage | `TestResetPassword_PasswordPolicy`, `TestResetPassword_BreachCheck_FailOpen` |

## Risk tier & rationale

**Tier 1** — INV-account-05 requires a property/invariant test proving
the credential update and the mass session-revoke happen atomically
(not as two separate steps where a crash between them could leave
old sessions alive after a "successful" reset), and INV-account-08
requires the same single-use race test pattern as email verification.
No Tier 0 sub-area (password hashing is standard library-backed, same
reasoning as `01-register-email-verification.md`).

## Assumptions / open questions

**A. Resolved — 2026-08-05.** `kencleng-phase0-detail.md` Fitur 2B
itself flagged this as "left to implementation, not financial-critical"
without picking a side. Formalized here: **repeated forgot-password
requests do not revoke previously-issued tokens.** Each token is
guarded independently by INV-account-08's single-use check, so having
more than one valid outstanding token for the same user is safe —
whichever is used first consumes itself; the others remain valid
until their own expiry or use, and using an older one afterward is
still safe (a real password-reset flow, not a state where multiple
uses could conflict — see the case where a stale token is used after
a newer one already succeeded: it would just reset the password again
to whatever the request specifies, correctly revoking sessions again
too). Chosen over proactively revoking older tokens because it's
simpler (no extra write on every forgot-password call) and the safety
property doesn't actually depend on it.

**B. Ordering requirement, not an ambiguity — stated explicitly
because getting it backwards would be a real bug.** Password-policy
validation (length + breach-list) must run **before** the
token-consuming atomic update, not after. If implemented in the wrong
order (mark `used_at` first, validate second), a user who mistypes a
weak password would burn their only reset token on a `422` and have
to request an entirely new one — a real correctness bug, not just a
UX papercut, since the acceptance criteria above explicitly promise
retry-with-the-same-link on a validation failure.

## Audit log entry?

No — per `kencleng-phase0-detail.md` Fitur 9 scope, password reset is
not in the audit-logged action list (role assign/revoke, MFA
enable/disable, account linking, self-PII reveal are the only account
actions in scope).

## References

- `docs/project/kencleng-phase0-detail.md` Fitur 2B
- `docs/project/kencleng-erd.md` §1 (`auth_tokens`, `refresh_tokens`)
- `docs/spec/account/threat-model.md` component 3
- `api/openapi.yaml` — `POST /auth/forgot-password`,
  `POST /auth/reset-password`