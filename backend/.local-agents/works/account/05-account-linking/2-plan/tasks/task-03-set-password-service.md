# Task 03: SetPassword service (both branches) + VerifyEmail audit delta

> Back-reference : `.local-agents/works/account/05-account-linking/2-plan/techplan.md` (Status: Approved) — sections 3, 4 (R1–R8, R14, R16), 5 (D4–D7), 8 (SetPassword flow block — keep verbatim as the logic source)
> Depends on    : task-02 (repository methods); task-01 (migration makes `email_verification_link` insertable)
> Model         : GLM 5.2 (max) (multi-step branching + timing-parity reasoning; compensating control: mandatory Complex-tier dual-model code review per model-routing)

## Objective

Implement the set-password flow with server-side branch selection: Branch 1 (add unverified identity + verification email, anti-enumeration incl. conflict nudge and race-loser fallback) and Branch 2 (change password in place + user-wide session revocation in ONE tx). Plus the `VerifyEmail` delta that writes the Branch-1 audit entry when the redeemed token carries the new purpose.

## Files

| File | Change |
|---|---|
| `backend/internal/domain/account/security.go` | New — `SetPassword`, sentinels shared by later tasks, `purposeEmailVerifyLink = "email_verification_link"` const |
| `backend/internal/domain/account/security_test.go` | New — unit + race tests below |
| `backend/internal/domain/account/service.go` | Modified — `VerifyEmail` captures `RedeemToken`'s purpose (currently discarded `_`) and conditionally writes the audit row inside the same tx; no signature change |
| `backend/internal/domain/account/service_test.go` | Modified — R14 unit coverage |
| `backend/internal/platform/notification/sender.go` | + `NudgeSetPasswordConflict = "set_password_conflict"` |

## Flow (authoritative — techplan §8)

```
SetPassword(ctx, userID, req):
  validatePassword(req.Password)                    # 422 before anything (R4)
  hash := HashPassword(req.Password)                # ALWAYS — timing parity
  hasEP := FindIdentifierHashByUserAndProvider(userID, "email_password").found
  if !hasEP:                                        # ---- Branch 1
     if FindAuthIdentityByIdentifierHash("email_password", HMAC(req.Email)) != nil:
        dummyWrite(); sendNudge(set_password_conflict); return nil        # generic 202
     tx {
        InsertAuthIdentity({user_id, email_password, Identifier:req.Email, verified_at:nil})
        InsertAuthToken({purpose:"email_verification_link", ...})          # single-use, 24h TTL
        on unique-violation: rollback; sendNudge(conflict); return nil     # race loser
     } commit
     sendVerificationEmail(req.Email, plainToken)   # AFTER commit, never inside tx
     return 202-generic
  else:                                             # ---- Branch 2
     secret := current identity.credential_secret
     compare(req.CurrentPassword, secret) != nil -> ErrInvalidCredentials # 401, no change
     tx {
        UpdateCredentialSecret(userID, hash)        # identifier untouched
        RevokeAllRefreshTokensForUser(userID)       # INV-account-05, same tx
     } commit -> 200
```

Key decisions inherited: D4 (branch signal = identity existence, never client fields), D5 (conflict handling mirrors registration incl. `dummyWrite` DB-time parity and `isUniqueViolation` fallback), D6 (`compare` seam against own identity's secret; comparable CPU burned on failure paths), D7 (audit at redemption, not creation).

## Rules to prove (unit/race scope; integration truth-tests land in task-06)

- **R1** `TestSetPassword_Branch1_CreatesUnverifiedIdentity_SendsVerification` — unverified identity + token purpose committed together; email after commit
- **R2** `TestSetPassword_Branch1_ClaimedEmail_NudgeNoIdentity_Generic202` — byte-identical body vs R1's response
- **R3** `TestSetPassword_ConcurrentDuplicateEmail_Race` — `-race`, loser rolls back clean + nudge + generic nil
- **R4** `TestSetPassword_PasswordPolicy_PrecedesBranching` + `TestSetPassword_BreachCheck_FailOpen` — zero side-effect rows on 422/fail-open
- **R5** `TestSetPassword_GenericResponse_AllBranches` — created/claimed/race-loser status+body parity table
- **R6** `TestSetPassword_BranchSelection_ServerSide` — misleading-field matrix both directions
- **R7/R8** unit portions of `TestSetPassword_Branch2_*` (atomicity probe with forced mid-tx failure; wrong-password zero-row assertion)
- **R14** `TestVerifyEmail_LinkPurpose_WritesLinkAudit` + `TestVerifyEmail_RegistrationPurpose_NoLinkAudit`; existing `TestVerifyEmail_TokenSingleUse_Concurrent` must pass UNTOUCHED
- **R16** `TestSecurity_LogsFreeOfSecrets` — log-scan over exercised paths; no plaintext email/password/hash/token in any log line; sanitized categories only on send failure (copy `notificationErrorCategory`)

## Common mistakes (techplan §13 subset)

- Writing the Branch-1 audit at identity creation → forbidden (spec: "not at initial creation"); it rides the link-purpose redemption
- Comparing `current_password` before the policy check → policy first, same ordering as registration
- Sending emails inside the open tx → post-commit only; failures non-fatal + sanitized log
- Echoing the submitted email anywhere in responses → constants only

## Out of scope here

Handlers/status codes (task-05); integration/testcontainers proofs (task-06); unlink entirely (task-04).

## Verification

`go test -race ./internal/domain/account/...` green including new tests.
