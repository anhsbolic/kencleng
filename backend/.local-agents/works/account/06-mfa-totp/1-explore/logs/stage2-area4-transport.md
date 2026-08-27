# Stage 2 — Area 4: Transport Layer — HTTP Handlers & Routing

> Files: `internal/transport/http/account_security.go`, `cmd/server/main.go`

## Current State

`account_security.go` has:
- `securityService` interface with `SetPassword` and `UnlinkGoogle` — no MFA methods
- `SetPasswordHandler` and `UnlinkGoogleHandler` — both use `RequireSession` middleware
- `UserIDFromContext` helper for extracting authenticated user ID

`main.go` wiring (lines 184-188):
```go
accountMux.HandleFunc("POST /account/security/set-password", ...)
accountMux.HandleFunc("POST /account/security/google/unlink", ...)
mux.Handle("/account/security/", RateLimit(rps, burst)(RequireSession(googleVerifyToken)(accountMux)))
```

The `/account/security/` route group already has `RequireSession` middleware. New MFA endpoints under this path will inherit it automatically.

## Requirement

Three new handlers:
1. **`MfaEnrollHandler`** — `POST /account/security/mfa/enroll`
   - Extract userID, call `svc.MfaEnroll`, return 200 with `MfaEnrollResponse` (`otpauth_uri`)
   - Error mapping: 409 (already enabled) → Problem Details

2. **`MfaEnrollConfirmHandler`** — `POST /account/security/mfa/enroll/confirm`
   - Extract userID, decode `MfaEnrollConfirmRequest` (`totp_code`)
   - Call `svc.MfaEnrollConfirm`, return 200 with `MfaEnrollConfirmResponse` (`backup_codes[]`)
   - Error mapping: 422 (invalid code / no pending secret) → Problem Details

3. **`MfaDisableHandler`** — `POST /account/security/mfa/disable`
   - Extract userID, decode optional `MfaDisableRequest` (`password`)
   - For email_password users: password required
   - For Google-only users: no body, check reauth marker
   - Call `svc.MfaDisable`, return 200 with generic message
   - Error mapping: 401 (wrong password / missing reauth marker) → Problem Details

## Gap

All three handlers are missing. The `securityService` interface needs to be extended with MFA methods. Routes need to be registered in `main.go`.

## Sniffing

- **Risk:** The `MfaDisableHandler` has a branching body requirement (password for email_password, no body for Google-only). The handler must determine which branch to take — either by checking the user's auth identities (extra DB call) or by having the service handle both paths. The service method signature needs to accommodate this.
- **Edge cases:** What if a Google-only user sends a body with `password: ""`? The handler should treat empty password the same as no body (check reauth marker). What if an email_password user sends no body? The handler should return 422 (password required).
- **Miscontext:** The OpenAPI spec says `MfaDisableRequest.password` is optional (not in `required`). This means the handler must handle both cases. The service needs to know which auth provider the user has to decide the re-auth method.
- **Misleading signal:** The `/account/security/` route group already has `RequireSession` — someone might think "auth is handled." It is, but the disable endpoint needs *re*-authentication (password or reauth marker) on top of the existing session. The session check is necessary but not sufficient.
- **Inconsistency:** None found. The OpenAPI contract is clear and matches the feature spec.
