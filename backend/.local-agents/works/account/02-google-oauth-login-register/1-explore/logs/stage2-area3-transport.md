# Stage 2 — Area 3: Transport/HTTP layer

> Feature: 02-google-oauth-login-register
> Date: 2026-08-22

## Current state

Files in `backend/internal/transport/http/`:
- `auth_register.go` — RegisterHandler pattern: decode JSON body → validate fields → call service → map error or write 202
- `auth_verify_email.go` — VerifyEmailHandler, ResendVerificationHandler. Same pattern.
- `errors.go` — WriteProblem, WriteValidationError, MapServiceError, write400InvalidJSON. All write JSON application/problem+json.
- `middleware.go` — RateLimit(rps, burst) middleware. Per-IP token bucket.
- `swagger.go` — dev-only Swagger UI handler.

All existing handlers are JSON request/response. No redirect handlers exist. No auth middleware exists.

Route wiring in cmd/server/main.go (line 118-122):
```go
authMux := http.NewServeMux()
authMux.HandleFunc("POST /auth/register", ...)
authMux.HandleFunc("POST /auth/verify-email", ...)
authMux.HandleFunc("POST /auth/verify-email/resend", ...)
mux.Handle("/auth/", transporthttp.RateLimit(rps, burst)(authMux))
```

## Requirement

1. GET /auth/google/redirect — reads intent query param, optionally validates session (for link/reauth), generates state+nonce, sets HttpOnly cookie, responds 302 to Google consent screen.
2. GET /auth/google/callback — reads code+state query params, reads cookie, validates state match, calls service to exchange code + verify id_token + branch on intent, sets tokens (or error), responds 302.

Both are GET endpoints with redirect responses, not JSON.

## Gap

1. **No redirect response pattern.** All existing handlers write JSON. Google OAuth handlers need 302 redirects with Location headers.
2. **No cookie management.** No utility for setting HttpOnly/Secure/SameSite cookies.
3. **No auth middleware for reading session.** link/reauth intents require existing authenticated session (access token). No middleware extracts user_id from JWT.
4. **No query param parsing pattern.** Existing handlers read from JSON body. OAuth handlers read from query params (intent, code, state).
5. **Conditional auth check.** GET /auth/google/redirect with intent=login is public, but intent=link/reauth requires auth. Handler needs to conditionally check auth based on intent — unusual pattern.

## Sniffing

- **Risk:** Conditional auth check is highest-risk pattern. If handler forgets to check auth for link/reauth, attacker could link their Google account to any user. Check must happen before redirect to Google.
- **Edge case:** Access token expired but cookie with state/nonce still valid. User returns from Google after minutes — session may have expired. Callback needs user_id from cookie (set at redirect time), not from fresh session check.
- **Miscontext:** Feature spec says callback "branches on intent" but callback is public endpoint. user_id for link/reauth comes from cookie, not session check at callback time.
- **Misleading signal:** Existing authMux pattern applies rate limiting to all /auth/ routes. New GET endpoints inherit this — correct behavior but worth noting.
- **Inconsistency:** Feature spec says state/nonce stored in "short-TTL HttpOnly cookie" but also says user_id is "encoded into the same cookie state." Cookie needs structured data (JSON), not simple string. Format not specified.
