# Stage 2 — Area 3: Service Layer — Enroll / Confirm / Disable

> File: `internal/domain/account/service.go`, `internal/domain/account/login.go`

## Current State

`service.go` has no MFA enrollment/disable methods. The only MFA-related code in the service is:
- The `mfa MfaVerifier` field (line 103) — injected dependency
- `LoginMfa` in `login.go` — consumes the verifier for login-time TOTP/backup-code verification
- `NewService` accepts `mfa MfaVerifier` parameter (line 138)

Existing patterns available to follow:
- Transactional multi-row writes (`registerNewUser`, `ResetPassword`, `UnlinkGoogle`)
- Audit log insertion (`InsertUserLog` in `UnlinkGoogle`, `VerifyEmail`)
- Password comparison (`s.compare` seam)
- Error vocabulary (`ErrValidation`, `ErrInvalidCredentials`, etc.)

## Requirement

Three new service methods:

1. **`MfaEnroll(ctx, userID) (otpauthURI string, error)`**
   - Check if MFA already enabled → 409
   - Generate TOTP secret, encrypt at rest
   - Upsert into `mfa_totp_secrets` with `enabled_at = NULL`
   - Return `otpauth://` URI for QR code

2. **`MfaEnrollConfirm(ctx, userID, totpCode) (backupCodes []string, error)`**
   - Load pending secret (`enabled_at IS NULL`)
   - Verify TOTP code against the pending secret
   - If invalid → 422, secret stays (user can retry)
   - If valid → set `enabled_at = now()`, generate 10 backup codes, hash+insert
   - Return plain backup codes (shown once)
   - Write `user_logs` audit entry (MFA enabled)

3. **`MfaDisable(ctx, userID, password) error`**
   - Re-auth: for email_password users, verify password
   - For Google-only users, empty password means "caller already re-authed"
   - Set `enabled_at = NULL`
   - Do NOT delete backup codes (INV-account-06 implicit invalidation)
   - Write `user_logs` audit entry (MFA disabled)

## Gap

All three methods are entirely missing. The service has no MFA enrollment/disable logic. The patterns exist but the specific methods don't.

## Sniffing

- **Risk:** `MfaEnroll` must reject re-enrollment when `enabled_at IS NOT NULL` (409). If this check is missing or races, a stray re-enroll call could silently replace `secret_encrypted` on an already-active account, breaking the user's authenticator app while the system still thinks MFA is active.
- **Edge cases:**
  - Enroll called multiple times before confirm: spec says "simply overwrite with a fresh one" — safe because `enabled_at` stays NULL. The upsert handles this.
  - Confirm called with no pending secret: spec says 422 (same as wrong code). The service should not distinguish — load the secret, if `enabled_at` is already non-NULL or no row exists, return the same error as wrong code.
  - Disable for Google-only user: needs the reauth marker from `auth_google.go`. The marker is in-memory (`sync.Map`), 5-min TTL. If the server restarted, the marker is lost — the user must re-authenticate. Accepted (same as `05-account-linking.md` precedent).
- **Miscontext:** The feature spec says disable for Google-only users uses "the short-lived server-side re-auth marker." This marker already exists (`CheckReauthMarker` in `auth_google.go:109`). The service method needs to call it — but it's a transport-layer function. The service needs a seam for this (injected dependency or the handler checks before calling service).
- **Misleading signal:** `LoginMfa` already handles TOTP/backup-code verification for login. Someone might think "MFA verification is already implemented." But `LoginMfa` consumes the `MfaVerifier` interface — which is still the stub. The verification logic is the stub, not real code. `LoginMfa` is the consumer; task #6 is the provider.
- **Inconsistency:** The feature spec says "exactly 10 backup codes" — the service must enforce this count. Backup codes should be generated with sufficient entropy (e.g., 8-10 alphanumeric characters from `crypto/rand`).
