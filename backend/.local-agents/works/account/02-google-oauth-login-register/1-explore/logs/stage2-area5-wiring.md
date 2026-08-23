# Stage 2 — Area 5: Entry point wiring

> Feature: 02-google-oauth-login-register
> Date: 2026-08-22

## Current state

cmd/server/main.go wiring sequence:
1. Load .env, validate required env vars (line 59-73)
2. Initialize crypto.Keys (encryption + HMAC)
3. Open DB pool
4. Initialize MinIO
5. Load JWT key pair (ES256)
6. Wire account domain: breachClient, emailSender, accountSvc
7. Configure rate limiting
8. Register routes on authMux
9. Graceful shutdown

Route registration pattern:
```go
authMux.HandleFunc("POST /auth/register", ...)
authMux.HandleFunc("POST /auth/verify-email", ...)
authMux.HandleFunc("POST /auth/verify-email/resend", ...)
mux.Handle("/auth/", transporthttp.RateLimit(rps, burst)(authMux))
```

## Requirement

1. New env vars: GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, GOOGLE_REDIRECT_URI, FRONTEND_URL
2. New platform dependency: Google OAuth client (passed to service)
3. New routes: GET /auth/google/redirect, GET /auth/google/callback
4. Auth middleware or handler-level JWT verification for link/reauth intents

## Gap

1. **No Google OAuth env vars.** Need to add to requireEnv and .env.example.
2. **No Google OAuth client construction.** Wiring needs to create client and pass to account.NewService (or new constructor).
3. **No auth middleware wiring.** link/reauth intents need session extraction. JWT public key needs to be available to handler/middleware.
4. **Route registration.** New GET endpoints need registration on authMux. Existing pattern works for GET too.

## Sniffing

- **Risk:** Service constructor (NewService) currently takes 5 parameters. Adding Google OAuth client makes 6. Still manageable.
- **Edge case:** If GOOGLE_CLIENT_ID or GOOGLE_CLIENT_SECRET is empty/misconfigured, server should fail fast at startup (same pattern as other required env vars).
- **Miscontext:** Feature spec mentions GOOGLE_REDIRECT_URI as "fixed env var, exact-match registered in Google Console." Wiring just needs to pass it through.
- **Misleading signal:** Existing auth.Load for JWT keys might suggest Google credentials should use similar Load pattern. But env vars are simpler for OAuth credentials.
- **Inconsistency:** None found.
