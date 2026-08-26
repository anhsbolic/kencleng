# Stage 3 — Solutioning

> Decisions D1–D6, amended with the three resolved questions (Q1–Q3,
> answered by Anhar during the explore session). Stage 2 findings live in
> `01`–`07` area logs; this doc only holds decisions, options, rationale.

## Q1 — VerifyEmail cross-purpose hole: FIX IN THIS SLICE (resolved)

`VerifyEmail` discards `RedeemToken`'s returned purpose (service.go:413),
so a `password_reset` token would work at `/auth/verify-email`. Decision:
fix in-slice — one-line purpose check in `VerifyEmail`
(`purpose != "email_verification"` → `ErrTokenNotFound`, deferred Rollback
un-redeems) + unit tests proving each flow rejects the other's token.
Rejected: deferring (reset tokens become live ammunition the moment this
feature ships). Scope addition is cross-referenced in the risk note.

## D1 — Repository additions

Two new methods on `Repository`, following existing precedents:

1. `UpdateIdentityCredentialSecret(ctx, tx pgx.Tx, userID uuid.UUID,
   providerType, passwordHash string) error` — keyed on
   `(user_id, provider_type)`, mirroring `SetUserVerified`: tokens carry
   `user_id`, not identity_id; INV-account-01 guarantees at most one
   identity per (user, provider).
   - Rejected: keying on identity ID (would widen `RedeemToken`'s
     RETURNING contract); raw SQL (golden rule).
2. `RevokeAllRefreshTokensForUser(ctx, tx pgx.Tx, userID uuid.UUID)
   error` — one goqu UPDATE `SET revoked_at = now() WHERE user_id = ?
   AND revoked_at IS NULL`, inside caller's tx; guard makes repeats
   idempotent (matches `RevokeRefreshTokenByHash` convention).
   - Rejected: looping `RevokeRefreshTokenFamily` per family (SELECT-first
     inside the critical tx + a family created between SELECT and UPDATE
     escapes revocation — genuine INV-05 hole under concurrency); revoke
     outside the tx (violates the invariant verbatim).

No migration needed — schema ready (Area 3).

## D2 — Token consumption & purpose check in ResetPassword

`ResetPassword`: `BeginTx` → `RedeemToken(hash)` → if
`ok && purpose != "password_reset"` → return `ErrTokenNotFound`
(rollback un-redeems; zero state change per spec's 404 row). `RedeemToken`
signature stays untouched.

- Rejected: adding a `purpose` parameter to `RedeemToken` (4-clause
  guard) — cleaner long-term but rewrites a merged slice's shared
  redemption path when an equally safe local check exists. Scope
  discipline.

## D3 — ForgotPassword service method

Mirror `Register`'s three-branch shape one-to-one:

```
identifierHash := HMAC(email)
identity := FindAuthIdentityByIdentifierHash(email_password, hash)
switch {
case identity != nil:            // registered
    issueResetToken(userID)      // BeginTx → InsertAuthToken ONLY (no RevokeTokens — Assumption A) → Commit
    SendPasswordResetEmail(email, plainToken)   // post-commit
case googleIdentity != nil:      // Google-only
    dummyWrite(ctx)              // timing shaping, same device as Register R3/R4
    SendNudgeEmail(email, NudgeGoogleOnly)
default:                         // no account
    dummyWrite(ctx)
}
return nil                        // all branches → nil → identical 202
```

- **No proactive revocation** (Assumption A): `issueNewVerificationToken`
  deliberately NOT reused — it contains `RevokeTokens`; a new smaller
  helper avoids the trap-shaped surface (Area 3 sniffing #4).
- New const `resetTokenTTL = 1 * time.Hour` next to `tokenTTL`.
- Timing: dummyWrite on no-op branches, same anti-side-channel reasoning
  as Register R7.
- Email: new `Sender.SendPasswordResetEmail(ctx, to, token)` on interface
  + FakeSender + DevSender (third outbox line type; AGENTS.md §5 doc
  sentence drift accepted as trivial). Google-only notice reuses
  `NudgeGoogleOnly` unchanged.

## D4 — ResetPassword service method (Assumption B is structural)

Fixed order of operations:

1. `validatePassword(newPassword)` — before ANY DB work; weak password
   burns no round-trip and cannot interact with the token. HIBP network
   I/O stays outside the tx.
2. `secrets.HashPassword(newPassword)` — bcrypt before `BeginTx`
   (~100ms CPU never holds a tx open; Register hashes before branching).
3. `BeginTx` → `RedeemToken` → purpose check (D2) →
   `UpdateIdentityCredentialSecret(user, "email_password", hash)` →
   `RevokeAllRefreshTokensForUser(user)` → `Commit`. INV-05's
   same-transaction requirement satisfied structurally: any failure
   before commit rolls the redeem back — token survives a failed reset,
   which IS Assumption B's property, obtained from the tx pattern rather
   than ordering discipline alone.
4. On `!ok`: `FindAuthTokenByHash` after rollback → expired →
   `ErrTokenExpired` (410), else `ErrTokenNotFound` (404) — byte-for-byte
   the VerifyEmail disambiguation. Concurrent double-submit loser sees
   `used_at` set → 404 (spec-conformant).

Not handled (risk note): access tokens untouched by reset; they die on
their own ≤15-min expiry. Spec scopes INV-05 to refresh tokens.

## D5 — Transport handlers & routing

- `ForgotPasswordHandler` = resend-handler semantics: decode JSON (400
  invalid-json), `looksLikeEmail` guard → **422 fieldError** (Q3:
  keep, consistent with resend precedent), call svc, swallow internal
  errors into the identical 202 with sanitized server log, always write
  GenericAcceptedMessage-shaped body.
- `ResetPasswordHandler`: decode; empty token → early `ErrTokenNotFound`
  (verify-email precedent, no DB hit, no timing distinction);
  `MapServiceError` covers 422/410/404/500 with zero new mapping code;
  success → 200 `{message}` matching contract's Indonesian example.
- Wiring: two `authMux.HandleFunc("POST /auth/forgot-password|
  reset-password", ...)` lines; rate limit inherited from the mount-time
  `RateLimit(...)` wrapper. `TestForgotPassword_RateLimited` asserts it.
- Problem-type URIs: reuse MapServiceError's existing `problems/*`
  verbatim. Unifying `errors/*` vs `problems/*` across spec examples and
  code is repo-wide cleanup, out of scope; recorded as assumption note.

## Q2 — Contract completion: ADD 429 TO RESET-PASSWORD (resolved)

Add `"429": $ref TooManyRequests` to `/auth/reset-password` in
`api/openapi/account.yaml`, regenerate bundle (`npm run bundle`), commit
both together per api/README workflow. This documents existing middleware
behavior — a spec completion, not a loosening. No other spec edits.

## D6 — Test plan (Tier 1 gate mapping)

| Required proof | Test | Level |
|---|---|---|
| Generic 202 identical across 3 branches | `TestForgotPassword_GenericResponse_AllBranches` | unit |
| No proactive revocation (Assumption A) | prior unexpired reset tokens still redeemable after a new forgot call | unit |
| Policy fail → token NOT consumed (Assumption B) | seed token, weak pw reset, assert `used_at IS NULL` | unit |
| Wrong-purpose token (both directions, Q1) | reset w/ email_verification token → 404 unconsumed; verify-email w/ password_reset token → 404 unconsumed | unit |
| Expired/not-found/used → 410/404/404 | table-driven, mirrors VerifyEmail cases | unit |
| Breach fail-open | checker errors → proceed; returns true → 422 | unit |
| INV-08 single-use race ≥100 goroutines `-race` | `TestResetPassword_TokenSingleUse_Concurrent` (+stress modeled on `TestRefresh_Stress_MixedValidAndReplayed`) | integration |
| INV-05 atomicity property: credential-updated ⟺ all refresh rows revoked, 2+ tokens across ≥2 families; injected-failure variant asserts rollback leaves token usable + sessions alive | `TestResetPassword_AllSessionsRevoked_Atomic` | integration |
| Timing side-channel | `TestForgotPassword_Timing_Branches_RealPostgres` | integration |
| Rate limit on both endpoints | handler/middleware test asserting 429 path | unit |

Coverage KPI (≥80% new lines) and security layers ride `make verify`.

## Implementation order

1. Repository methods (+ interface docs)
2. Service: consts (`resetTokenTTL`, purpose constant), ForgotPassword,
   ResetPassword, VerifyEmail purpose check (Q1)
3. Notification: `SendPasswordResetEmail` on Sender + FakeSender +
   DevSender
4. Transport: two handlers + routing lines
5. openapi edit (Q2) + `npm run bundle`
6. Tests: unit → integration
7. `make verify`

## Assumptions / open notes carried into build

- Problem-type URI prefix unification deferred (repo-wide).
- DevSender outbox gains a third line type; AGENTS.md §5 wording drift
  accepted.
- Access-token non-revocation on reset accepted (INV-05 scopes refresh
  tokens only).
