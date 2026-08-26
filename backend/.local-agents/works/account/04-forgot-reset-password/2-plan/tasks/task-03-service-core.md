# Task 03 — Service core: ForgotPassword, ResetPassword, VerifyEmail purpose check

> Back-reference (contract): `../techplan.md` — sections 1–8 are the source of truth. Techplan wins over this file on any apparent conflict.
> Splitting axis: dependency/sequence chain (see `manifest.md`). Depends on Task 01 (repository methods) and Task 02 (`Sender` method) — both must be merged first.

## Scope

**In scope:**
- `internal/domain/account/service.go`: new consts, `issueResetToken` helper, `ForgotPassword`, `ResetPassword`, and the one-line `VerifyEmail` purpose check
- The full fake-based unit suite for all three flows

**Out of scope (this task):**
- Handlers/routing (Task 04), openapi edit (Task 05), integration/real-Postgres suite (Task 06)

## Dependencies

- Task 01: `UpdateIdentityCredentialSecret`, `RevokeAllRefreshTokensForUser` exist and are merged
- Task 02: `notification.Sender.SendPasswordResetEmail` exists and all fakes compile

## Requirements (techplan §3 — rows governing this task)

| Condition | Requirement | Source/Note |
|---|---|---|
| Forgot for email with `email_password` identity | New single-use token (`purpose=password_reset`, 1h), reset email sent, generic 202 handled by handler | feature AC; INV-account-08 |
| Forgot for google-only email | No token; `NudgeGoogleOnly` notice; identical API outcome | threat model §3 |
| Forgot for unknown email | Nothing sent; identical outcome | anti-enumeration |
| Repeated forgot requests | Each issues its own token; prior unexpired tokens NOT revoked | Assumption A (resolved) |
| Reset, valid token + passing password | ONE tx: credential updated, `used_at` set, EVERY refresh token revoked | INV-account-05 |
| Password policy | ≥8 chars + breach-list fail-open; validated BEFORE the token-consuming tx | Assumption B — ordering is load-bearing |
| Policy failure | `ErrValidation`; token stays unused | feature AC error table |

## Rules this task must satisfy (verbatim from techplan §4)

- **R1**: token row `{purpose: password_reset, expires_at ≈ now+1h, used_at NULL}` + reset email carries plain token.
- **R2**: google-only → no row, nudge sent, response-equivalent to R1.
- **R3**: no match → no row, no email, equivalent to R1.
- **R4**: prior unexpired reset token remains redeemable after a new forgot call (no `revoked_at` set).
- **R7**: success path updates credential + sets used_at + revokes every refresh row in ONE tx.
- **R8**: policy failure → `ErrValidation` AND token's `used_at` remains NULL.
- **R10**: token not found / already used / wrong purpose → `ErrTokenNotFound`; no state change.
- **R12**: `password_reset` token at `VerifyEmail` → `ErrTokenNotFound` and NOT consumed (rollback). *(Q1 fix)*
- **R13**: `email_verification` token at `ResetPassword` → `ErrTokenNotFound` and NOT consumed. *(D2/Q1 mirror)*
- **R14**: post-commit send failure → operation still succeeds; log carries sanitized category only — no recipient, no token.

(R5 timing-shaping, R6 rate-limit, R9/R11/R15–R18 belong to Tasks 04/06 but the service logic this task writes is what they prove.)

## Binding decisions (techplan §5)

| Decision | Resolution |
|---|---|
| D2 purpose enforcement | Check `RedeemToken`'s RETURNING purpose in-service; mismatch → `ErrTokenNotFound`; deferred Rollback un-redeems. Do NOT add a `purpose` param to `RedeemToken` (rejected: rewrites merged task #1/#3 shared path) |
| Q1 VerifyEmail hole | Fix IN THIS SLICE (Anhar, resolved): `userID, purpose, ok := ...` + reject non-`email_verification` purposes |
| D3 issuance helper | NEW insert-only `issueResetToken` — do NOT reuse `issueNewVerificationToken` (contains `RevokeTokens` → would violate Assumption A; trap-shaped surface) |
| D4 ordering | `validatePassword` (HIBP network I/O) then bcrypt hash BOTH before `BeginTx` — no external calls or ~100ms CPU inside the tx |
| D3b timing | No-op branches call existing `dummyWrite(ctx)` so DB-time matches the real branch (Register R7 precedent) |

## Implementation details (redistributed from techplan §10)

**File**: `internal/domain/account/service.go`
- Consts next to the existing block (service.go:43–52): `resetTokenTTL = time.Hour` (do NOT reuse `tokenTTL`=24h) and `purposePasswordReset = "password_reset"`.
- Unexported `issueResetToken(ctx context.Context, userID uuid.UUID) (plain string, err)`: modeled line-for-line on `issueNewVerificationToken` (service.go:352–390) MINUS the `RevokeTokens` call; BeginTx → `InsertAuthToken{Purpose: purposePasswordReset, ExpiresAt: now+resetTokenTTL}` → Commit; plain token returned for post-commit send.
- `ForgotPassword(ctx, email) error`:
```
identifierHash := crypto.HMAC([]byte(email), s.keys)
identity := repo.FindAuthIdentityByIdentifierHash(email_password, hash)
match       -> issueResetToken(identity.UserID); s.sendPasswordReset post-commit; return nil
google-only -> dummyWrite(ctx); s.sendNudge(email, NudgeGoogleOnly); return nil
none        -> dummyWrite(ctx); return nil
```
  All branches return nil — the handler owns the identical 202 (same contract as `ResendVerification`).
- `ResetPassword(ctx, token, newPassword string) error`:
```
validatePassword(newPassword)            // R8 gate; BEFORE any DB work
hash := secrets.HashPassword(newPassword) // CPU before tx
tx { userID,purpose,ok := RedeemToken(sha256Hex(token))
     if ok && purpose != purposePasswordReset -> ErrTokenNotFound (deferred rollback un-redeems)
     UpdateIdentityCredentialSecret(userID, providerEmailPassword, hash)
     RevokeAllRefreshTokensForUser(userID)
     commit }
!ok -> FindAuthTokenByHash after rollback -> !ExpiresAt.After(now) ? ErrTokenExpired : ErrTokenNotFound
```
  Byte-for-byte the `VerifyEmail` disambiguation shape (service.go:434–444).
- `VerifyEmail` change (Q1): capture purpose from `RedeemToken` (currently discarded at service.go:413); `if ok && purpose != purposeEmailVerify { return ErrTokenNotFound }` — deferred Rollback handles un-redeem. This is a deliberate behavior tightening of a shipped endpoint (techplan §6).
- Error wrapping `%w` throughout; no PII/token in any log line.

## Testing checklist (this task's items from techplan §12)

Unit suite, fake-based (extend `service_test.go` or new `password_reset_test.go`; table-driven where >2 cases):

- [ ] R1 `TestForgotPassword_Match_IssuesTokenAndSendsEmail` — asserts purpose/expiry window/plain-token-in-sender-arg
- [ ] R2 `TestForgotPassword_GoogleOnly_NoticeNoToken`
- [ ] R3 `TestForgotPassword_NoMatch_NothingSent`
- [ ] R4 `TestForgotPassword_Repeat_DoesNotRevokePriorTokens` — Assumption A
- [ ] R5(partial) `TestForgotPassword_GenericResponse_AllBranches` — all branches return nil + expected sender calls (timing proof is Task 06's real-DB test)
- [ ] R7 `TestResetPassword_HappyPath_UpdatesAndReturnsNil` (fake asserts call order: redeem→update→revoke-all, same tx handle)
- [ ] R8 `TestResetPassword_PasswordPolicy_TokenNotConsumed` — length + breach-hit variants; assert redeem never invoked OR rollback left token unconsumed per fake semantics
- [ ] R9-side unit disambiguation `TestResetPassword_Expired_vs_NotFound_Mapping`
- [ ] R10 `TestResetPassword_NotFound_Used_404_Errors` + `TestResetPassword_WrongPurpose_404_Unconsumed`
- [ ] R12 `TestVerifyEmail_RejectsResetPurposeToken_Unconsumed` — Q1
- [ ] R13 covered by wrong-purpose test above (count-parity noted in techplan §12)
- [ ] Breach fail-open `TestResetPassword_BreachCheck_FailOpen` (checker errors → proceed; returns true → `ErrValidation`)
- [ ] R14 `TestPasswordReset_SendFails_LogsNoPIIOrToken` — mirror `TestRegister_SendVerificationFails_LogNoPII` (service_test.go:1140)

Gate hygiene: `go test -race ./internal/domain/account/...` green.

## Common mistakes that apply here (techplan §13)

| Mistake | Fix |
|---|---|
| Redeeming before password validation | D4 order; R8 asserts token survives 422 |
| Reusing `tokenTTL` (24h) | `resetTokenTTL`; R1 asserts window |
| Calling `RevokeTokens(userID,"password_reset")` anywhere in forgot flow | Violates Assumption A; insert-only helper; R4 |
| Discarding redeem purpose (the exact VerifyEmail bug shape) | Purpose checks both directions; R12/R13 |
| Sending email inside the tx | Post-commit only; R14 |
| Logging `err` verbatim on sender/DB errors | `notificationErrorCategory` style; R14 |
