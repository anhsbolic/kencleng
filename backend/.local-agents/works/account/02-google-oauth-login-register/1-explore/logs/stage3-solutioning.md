# Stage 3 — Solutioning

> Feature: 02-google-oauth-login-register
> Date: 2026-08-22

## Decisions

### 1. Cross-task dependency — Token issuance

**Problem:** Task #2 callback needs to issue access + refresh tokens. Token issuance is task #3 scope.

**Decision: Option C** — Pull minimal token issuance into task #2.

Task #2 creates refresh_tokens table + IssueTokens method. Task #3 then adds login endpoint, MFA check, lockout, refresh rotation, logout — consuming the primitives task #2 created.

### 2. Google OAuth client — library choice

**Decision: Option B** — golang-jwt/jwt/v5 + manual JWKS fetch.

Already a transitive dependency. ~40 lines for JWKS keyfunc with cache-on-miss. Token exchange is single http.Post to Google endpoint. Avoids pulling in go-oidc full OIDC machinery.

### 3. Cookie management — where to put it

**Decision: Option A** — Handler-level utility in transport/http/cookie.go.

Two functions: setOAuthStateCookie(w, state, nonce, intent, userID) and readOAuthStateCookie(r) (state, nonce, intent, userID, error). JSON-encoded, base64-serialized, 10-min Max-Age.

### 4. Auth middleware for link/reauth intents

**Decision: Option A** — Handler-level JWT verification.

Handler checks intent first. If link/reauth, reads Authorization header, verifies JWT using keys.Public, extracts user_id, stores in cookie. If login, skips auth. Extract shared verifyAccessToken(tokenString) (userID, error) in platform/auth/.

### 5. Reauth marker — storage

**Decision: Option A** — In-memory map with TTL eviction, using Redis.

Redis-backed. 5-min TTL. MFA-disable endpoint checks this. Survives restarts, supports multi-instance.

### 6. Error redirect format

**Decision: Proposed format.**

- Success: 302 to {FRONTEND_URL}/dashboard (cookies set)
- Error: 302 to {FRONTEND_URL}/login?error={error_code}

Error codes:
- google_email_conflict — no auto-merge (login intent, email claimed by email_password)
- google_link_conflict — email claimed by different user (link intent)
- google_unavailable — Google API timeout/unreachable
- state_mismatch — state param doesn't match cookie
- nonce_mismatch — nonce in id_token doesn't match stored nonce

## Implementation scope

### New files

- backend/internal/domain/account/google_oauth.go — service methods: GoogleRedirect, GoogleCallback, IssueTokens
- backend/internal/domain/account/google_oauth_test.go — tests
- backend/internal/transport/http/auth_google.go — handlers: GoogleRedirectHandler, GoogleCallbackHandler
- backend/internal/transport/http/cookie.go — cookie utility functions
- backend/internal/platform/googleoauth/client.go — Google OAuth client (token exchange, JWKS verification)
- backend/migrations/000004_create_refresh_tokens.up.sql / .down.sql — refresh tokens table

### Modified files

- backend/internal/domain/account/service.go — add googleOAuth dependency, add IssueTokens method
- backend/internal/domain/account/repository.go — add InsertRefreshToken
- backend/internal/domain/account/repository_db.go — implement new repository methods
- backend/internal/domain/account/entity.go — add RefreshToken entity
- backend/internal/platform/auth/keys.go — add VerifyAccessToken helper
- backend/cmd/server/main.go — wire Google OAuth client, new env vars, new routes
- backend/.env.example — add GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, GOOGLE_REDIRECT_URI, FRONTEND_URL

### Not built (deferred to task #3)

- Login endpoint (POST /auth/login)
- MFA check at login
- Login attempt lockout
- Refresh rotation + reuse detection
- Logout endpoint
