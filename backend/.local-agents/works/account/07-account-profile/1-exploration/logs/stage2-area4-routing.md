# Stage 2 — Area 4: Route Wiring

## Current State

`cmd/server/main.go` registers routes as follows:

- **Auth routes** (lines 154-177): Mounted on `authMux`, wrapped with `RateLimit`, served under `/auth/`.
- **Account security routes** (lines 186-194): Mounted on `accountMux`, wrapped with `RateLimit` + `RequireSession`, served under `/account/security/`.

The middleware chain for account security is:
```
mux.Handle("/account/security/", RateLimit(rps, burst)(RequireSession(googleVerifyToken)(accountMux)))
```

The `accountMux` is a `http.NewServeMux()` that handles `POST /account/security/set-password`, `POST /account/security/google/unlink`, and the three MFA endpoints.

## Requirement

Add `GET /account/me` behind `RequireSession` (auth required). Rate limiting is appropriate (public-facing endpoint).

## Gap

`GET /account/me` has a different path prefix (`/account/me`) than the existing security routes (`/account/security/*`). Two options:

**Option A: New route on main mux** — Add directly to `mux` with its own middleware chain:
```go
mux.Handle("GET /account/me", RateLimit(rps, burst)(RequireSession(googleVerifyToken)(AccountMeHandler(accountSvc))))
```
Simple, explicit, no restructuring needed.

**Option B: New `/account/` sub-mux** — Create a broader `accountMux` that handles both `/account/me` and `/account/security/*`. More restructuring, but groups all account routes together.

Option A is simpler and follows the existing pattern (each route group has its own middleware chain declaration).

## Sniffing

1. **Risk**: None. Route wiring is mechanical. The middleware chain is already proven.

2. **Edge cases**: None. Go 1.22+ pattern routing handles `GET /account/me` correctly.

3. **Miscontext**: None. The spec says `GET /account/me` — no prefix ambiguity.

4. **Misleading signals**: The existing `accountMux` handles `/account/security/*` — looks like it could be extended, but the path prefix mismatch means it can't serve `/account/me`.

5. **Inconsistency**: None. The pattern is consistent with how other route groups are wired.
