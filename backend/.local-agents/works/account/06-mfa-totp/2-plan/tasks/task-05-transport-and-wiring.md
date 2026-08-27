# Task 05: Transport + wiring — three handlers, marker consumption, route registration

> Back-reference : `.local-agents/works/account/06-mfa-totp/2-plan/techplan.md` (Status: Approved by Anhar) — sections 8 (Transport delta + API changes), 10 (account_security.go, auth_google.go, errors.go, main.go), 5 (D6, D7, D8), 9 (steps 6–7)
> Depends on    : task-03 (service methods exist), task-04-scaffold (`NewMfaVerifier` constructor name for main.go; final merge ordering still waits on pairing sign-off)
> Model         : DeepSeek V4 Pro (contract-shape precision: exact status codes, problem-type URIs, response schemas)

## Objective

Expose the three endpoints behind the existing `RequireSession` + rate-limited `/account/security/*` group; add marker *consumption* beside the existing checker; map the new sentinels to Problem Details.

## Files

| File | Change |
|---|---|
| `backend/internal/transport/http/account_security.go` | `securityService` += 3 signatures; `MfaEnrollHandler`, `MfaEnrollConfirmHandler`, `MfaDisableHandler` |
| `backend/internal/transport/http/auth_google.go` | +`ConsumeReauthMarker(userID uuid.UUID) bool` — `LoadAndDelete` + expiry recheck (D7); doc comment cites feature-06 consume-on-use clause; existing functions untouched |
| `backend/internal/transport/http/errors.go` | `MapServiceError`: `ErrMfaAlreadyEnabled` → 409 type `https://kencleng.dev/errors/mfa-already-enabled`; both 422-sentinels → shared validation writer |
| `backend/internal/transport/http/account_security_test.go` | modified — handler-contract + marker tests (R12/R13 transport halves, R14 matrix, R16) |
| `backend/cmd/server/main.go` | `mfaVerifier := account.NewMfaVerifier(account.NewRepositoryDB(pool, keys), keys)` replacing bare `nil`; 3 `accountMux.HandleFunc("POST /account/security/mfa/…")` registrations |

## Handler contracts

- **Enroll**: session userID → service → 200 `{otpauth_uri}` / 409 via Map. URI never logged.
- **Confirm**: decode `{totp_code}` required-field validate → service → 200 `{backup_codes}` (echoed once) / 422 identical-shape for both failure reasons (R7 discipline maintained at handler too).
- **Disable** (techplan §8 Transport delta verbatim):
  ```
  body.password empty && caller google-only -> ConsumeReauthMarker(userID)? proceed : 401
  (marker state never crosses into the domain service)
  ```
  Empty body parses as zero-value struct (tolerant decode). Missing password for email_password caller → 422 field-required. Provider determination stays server-side through the service's identity check — R14 keeps branch selection out of client hands; misleading `password` from a Google-only caller does NOT bypass the marker requirement.
- All three: follow the decode → boundary-validate → service → `MapServiceError` shape used by SetPasswordHandler (never bare-default 401 misclassifying DB failures).

## Implementation constraints

- Route registration rides the EXISTING wrapped accountMux — no new middleware, no limiter change (R16 inherited; prove via `TestMfaRoutes_WiredBehindRequireSession`)
- `ConsumeReauthMarker` does not alter `CheckReauthMarker` or the sweeper (D7 constraint)
- Wire only after suites from tasks 02–04 are green

## Rules enabled (proven here at transport level)

R3 (409 problem-type), R7 (identical bodies through the wire), R12–R14 (status-code matrix), R16.

## Verification

- `go test ./internal/transport/http/... -race` green
- Existing RequireSession suite re-run untouched and green against enlarged mux
- `go build ./...` + hand smoke: server boots with real verifier injected
