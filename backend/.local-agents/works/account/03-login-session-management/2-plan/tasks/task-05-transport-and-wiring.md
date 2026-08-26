# Task 05: HTTP handlers, cookies, error mapping, wiring

> Back-reference : `.local-agents/works/account/03-login-session-management/2-plan/techplan.md` (Status: Approved) — sections 8 (API changes table), 10 (transport entries), 13
> Depends on    : task-04 (service methods exist)
> Model         : DeepSeek V4 Pro (byte-equal contract precision; no diagram judgment needed)
> Rules touched : R1–R4, R7–R8, R10–R11, R13, R16 (transport halves), R19
> Tier 0        : none in this task

## Objective

Expose the four endpoints on the existing rate-limited `authMux`, shape responses byte-exactly per the amended contract (`api/openapi/account.yaml` — `LockedOutGenericCredentials`), and wire `MFA_PENDING_TOKEN_SECRET` end-to-end.

## Files

| File | Change |
|---|---|
| `backend/internal/transport/http/auth_login.go` | New |
| `backend/internal/transport/http/cookie.go` | Edit |
| `backend/internal/transport/http/errors.go` | Edit |
| `backend/cmd/server/main.go` | Edit |
| `backend/internal/transport/http/auth_login_test.go` | New |

## Handlers (`auth_login.go`) — pattern per `auth_register.go`

Constructor-injected service seam (interface like `googleOAuthService`, so tests stub it) + `cookieSecure bool`.

- **`LoginHandler`**: decode `{email,password}` → boundary validation (reuse `looksLikeEmail`; password non-empty check only) → `svc.Login` →
  - `Status:"ok"`: JSON `200 LoginResponse{status:"ok", access_token, access_token_expires_at, user}` (**access token in body — NOT a cookie**) + `writeRefreshCookie`.
  - `Status:"mfa_required"`: JSON `200 {status:"mfa_required", mfa_pending_token}` — **no Set-Cookie of any kind**.
  - Sentinels via `MapServiceError`.
- **`LoginMfaHandler`**: decode `{mfa_pending_token, totp_code?, backup_code?}` → exactly-one-code boundary check → `svc.LoginMfa` → 200 `LoginResponse` + refresh cookie; sentinels mapped.
- **`RefreshHandler`**: read cookie via `readRefreshCookie` (absent/empty ⇒ 401 Problem directly) → `svc.Refresh` → 200 `RefreshResponse{access_token, access_token_expires_at}` + replacement `writeRefreshCookie`. Never read body/bearer here.
- **`LogoutHandler`**: best-effort read cookie → `svc.Logout(plain)` (nil-safe on absence) → `clearRefreshCookie` → **always 204**, no body.

## Cookie helpers (`cookie.go`)

```go
// writeRefreshCookie sets kencleng_refresh: HttpOnly, Secure per env,
// SameSite=Strict (contract-mandated), Path="/", MaxAge = refreshTokenCookieTTL.
func writeRefreshCookie(w http.ResponseWriter, cookieSecure bool, value string)
// clearRefreshCookie: MaxAge<0 deletion variant.
func clearRefreshCookie(w http.ResponseWriter, cookieSecure bool)
```

Do NOT touch `writeAuthCookies` (OAuth 302 contract) or `sessionToken()`/`GoogleTokenVerifier`. Do NOT reuse `writeAuthCookies` at `/auth/login` — it would over-deliver an undefined access cookie.

## Error mapping (`errors.go`) — bodies are contract, byte-for-byte

```go
case errors.Is(err, account.ErrInvalidCredentials):
    // 401, type "https://kencleng.dev/errors/invalid-credentials",
    // title "Invalid Credentials", detail "Email atau password salah."
case errors.Is(err, account.ErrLockedOut):
    // 429, SAME title+detail as above,
    // type "https://kencleng.dev/errors/too-many-requests"
    //   ← matches openapi LockedOutGenericCredentials (amended 2026-08-26)
case errors.Is(err, account.ErrMfaPendingInvalid):
    // 401, generic invalid-credentials-family body (spec: one 401 for
    // expired/malformed/wrong-code family — see contract table)
```

Anti-enumeration is load-bearing: wrong-email, wrong-password, and lockout differ ONLY by status code. Consider extracting shared title/detail constants asserted byte-equal in tests.

## Wiring (`cmd/server/main.go`)

1. Add `"MFA_PENDING_TOKEN_SECRET"` to `requireEnv(...)`.
2. `pendingSecret := auth.ValidateMFAPendingSecret(os.Getenv(...))` — fail fast at startup.
3. Build mint/verify closures over task-03 funcs (injecting clock).
4. Construct stub `mfa_verifier` (task #6 replaces later).
5. Extend `NewService(...)` call with new seams.
6. Register routes on `authMux` (inherits the `/auth/` RateLimit wrapper automatically — no new middleware):
   ```go
   authMux.HandleFunc("POST /auth/login", transporthttp.LoginHandler(accountSvc, cookieSecure))
   authMux.HandleFunc("POST /auth/login/mfa", transporthttp.LoginMfaHandler(accountSvc, cookieSecure))
   authMux.HandleFunc("POST /auth/refresh", transporthttp.RefreshHandler(accountSvc, cookieSecure))
   authMux.HandleFunc("POST /auth/logout", transporthttp.LogoutHandler(accountSvc, cookieSecure))
   ```

## Handler tests (`auth_login_test.go`, stubbed service seam)

- Cookie attributes on login-ok: name `kencleng_refresh`, HttpOnly, Strict, MaxAge=30d, Secure toggled by `cookieSecure`; **no access-token cookie**.
- mfa-required: zero Set-Cookie headers.
- Byte-equal bodies: wrong-email response == wrong-password response == lockout response modulo status (R3/R4 transport half).
- Refresh: replacement cookie present; missing cookie ⇒ 401 Problem.
- Logout: 204 with cleared cookie (MaxAge<0); no-cookie ⇒ still 204.
- R19 marker leak-sweep across all four handlers (tokens/passwords/emails absent from output AND log capture).
- Validation errors → 422 field-error shape per house pattern.

## Out of scope

Service internals (task-04); DB behavior (tasks 02/06); touching `middleware.go` or any OAuth handler.
