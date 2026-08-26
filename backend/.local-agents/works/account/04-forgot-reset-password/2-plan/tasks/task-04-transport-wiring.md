# Task 04 — Transport: handlers, routing, rate-limit inheritance

> Back-reference (contract): `../techplan.md` — sections 1–8 are the source of truth. Techplan wins over this file on any apparent conflict.
> Splitting axis: dependency/sequence chain (see `manifest.md`). Depends on Task 03 (service methods exist).

## Scope

**In scope:**
- New file `internal/transport/http/auth_password_reset.go`: two handlers
- Two route registrations in `cmd/server/main.go`
- Handler-level tests (decode/validation/status mapping/rate-limit path)

**Out of scope (this task):**
- Service internals (Task 03), contract file edits (Task 05), integration suite (Task 06)

## Dependencies

Task 03 merged (`Service.ForgotPassword`, `Service.ResetPassword` callable).

## Requirements (techplan §3 rows governing this task)

| Condition | Requirement | Source |
|---|---|---|
| Rate limiting | Both endpoints inherit the stricter `/auth/*` bucket automatically | mount-time middleware, main.go:172 |
| Malformed email on forgot | 422 fieldError (`email`) — resend-handler precedent kept | Q3 resolved |
| Internal failure on forgot | Swallowed into IDENTICAL 202 with sanitized server log | enumeration-leak rule |
| Empty token at boundary | 404 without DB round-trip | verify-email precedent |

## Rules this task must satisfy (verbatim from techplan §4)

- **R5**: all three forgot branches produce byte-identical 202 bodies from this layer.
- **R6**: burst-exceeding requests to either endpoint → 429 Problem Details from the shared limiter.
- **R15**: empty `token` field → 404 without a DB round-trip.
- **R16**: malformed JSON body → 400 invalid-json Problem.
- **R17**: malformed email on forgot-password → 422 fieldError (`email`).

## Binding decisions (techplan §5)

| Decision | Resolution |
|---|---|
| Q3 malformed-email behavior | KEEP 422 like the resend handler; do not skip validation; do not edit openapi to add 422 on forgot (behavior addition beyond precedent needs human) |
| Problem-type URIs | Reuse `MapServiceError`'s existing `problems/*` URIs verbatim — NO new URI vocabulary (prefix unification is deferred repo-wide, techplan §14 Active #1) |
| Success body reset-password | 200 `{"message": "Password berhasil diubah. Silakan login ulang."}` exactly as contract example (account.yaml:284) |
| Generic 202 body | `GenericAcceptedMessage` shape — Indonesian generic text per contract example (account.yaml:882) |

## Implementation details (redistributed from techplan §10)

**File**: `internal/transport/http/auth_password_reset.go` (new; naming follows `auth_*.go` siblings)
- `ForgotPasswordHandler(svc *account.Service) http.HandlerFunc` — clone of `ResendVerificationHandler` semantics (auth_verify_email.go:52–86):
  1. decode JSON → `write400InvalidJSON` on error
  2. `looksLikeEmail(req.Email)` false → `WriteValidationError(w, []fieldError{{Field: "email", Message: "must be a valid email"}})`
  3. call `svc.ForgotPassword`; on err: log `transport: forgot password failed (recipient redacted): %v` — wait: sanitize like resend does (resend logs the error chain because its leaf is a DB driver error without PII values; keep that exact pattern and comment)
  4. ALWAYS write 202 + generic message JSON regardless of internal outcome
- `ResetPasswordHandler(svc *account.Service, ...) http.HandlerFunc`:
  1. decode JSON → 400 invalid-json
  2. empty token → `MapServiceError(w, account.ErrTokenNotFound)` and return (pre-DB, no timing distinction — mirror auth_verify_email.go:32–35)
  3. call `svc.ResetPassword`; err → `MapServiceError` (already maps ErrValidation→422, ErrTokenExpired→410, ErrTokenNotFound→404, unknown→generic 500 with server-side log — errors.go:76–109; ZERO new mapping code)
  4. success → 200 Content-Type json + the message object above
- Request DTOs: `forgotPasswordRequest{Email string}`, `resetPasswordRequest{Token, NewPassword string}` — names/shapes match openapi `ForgotPasswordRequest`/`ResetPasswordRequest`.

**File**: `cmd/server/main.go`
- Two lines beside the existing authMux registrations (main.go:152–163):
  `"POST /auth/forgot-password"` and `"POST /auth/reset-password"` → the two handlers. They inherit the mount-time `transporthttp.RateLimit(rps, burst)` wrapper (main.go:172) automatically — DO NOT add per-endpoint limiters.
- TBD — verify current `rps`/`burst` values at build time; read them, do not retune them in this slice (techplan §14 Active #4).

## Testing checklist (this task's items from techplan §12)

Handler tests in `internal/transport/http/auth_password_reset_test.go` (pattern: `auth_verify_email_test.go`):

- [ ] R5 `TestForgotPassword_Handler_Generic202_AllBranches` — stub svc returning nil vs error paths; assert byte-identical 202 bodies
- [ ] R6 `TestForgotPassword_RateLimited` + `TestResetPassword_RateLimited` — burst-exhausting requests against the wired mux → 429 problem shape
- [ ] R15 `TestResetPassword_EmptyToken_404_NoDBHit` — inject a service spy; assert zero service calls on empty token
- [ ] R16 malformed JSON → 400 invalid-json on BOTH endpoints
- [ ] R17 `TestForgotPassword_MalformedEmail_422`
- [ ] Status-mapping passthrough test: svc returning each sentinel (`ErrValidation`, `ErrTokenExpired`, `ErrTokenNotFound`) yields 422/410/404 with the established problem shapes
- [ ] Reset success body matches the contract example string exactly

Gate: `go build ./...`, `go test -race ./...` green; server boots locally with the two routes reachable (`go run ./cmd/server`, manual smoke optional).

## Common mistakes that apply here (techplan §13)

| Mistake | Fix |
|---|---|
| Adding endpoint-specific rate limiting or retuning global values | Inherit mount wrapper; values untouched |
| Distinguishing forgot branches at HTTP layer (different bodies/status) | Branch distinction lives ONLY in which email gets sent; HTTP surface identical |
| Echoing the token back in any response/log | Token appears nowhere except the email argument |
