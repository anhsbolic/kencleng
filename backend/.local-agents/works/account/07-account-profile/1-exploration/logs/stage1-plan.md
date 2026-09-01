# Stage 1 — Plan Announcement

## Task Understanding

Implement `GET /account/me` — a read-only endpoint that returns the authenticated user's own profile (id, name, email, email_verified, roles, auth_providers, mfa_enabled, created_at). The resource is entirely determined by the session token; there is no ID parameter, so no IDOR surface.

## Areas to Explore (in order)

1. **Repository layer** — `GetLoginUserView` in `repository_db.go` already assembles a `LoginUserView` with decrypted email, roles, auth providers, and MFA status. Need to confirm it covers all fields the spec requires and that it's suitable for direct reuse (vs needing a new query).

2. **Service layer** — No `GetProfile` or `GetMe` method exists on `*Service` today. Need to determine whether a thin service wrapper is needed or if the handler can call the repository directly (per existing patterns in this codebase).

3. **Transport layer** — Need a new handler file. The `RequireSession` middleware and `UserIDFromContext` helper already exist in `account_security.go`. Need to confirm the handler pattern (interface seam for testing, JSON response shape matching OpenAPI `User` schema).

4. **Route wiring** — `cmd/server/main.go` currently mounts `/account/security/` under `RequireSession`. Need to add `GET /account/me` under the same authenticated mux (or a sibling).

5. **OpenAPI contract** — Confirm the `User` schema fields and response codes match what we'll emit.

## Why This Order

Repository first because it determines whether we need new SQL or can reuse existing queries. Service layer next because it decides the handler's dependency shape. Transport and wiring are downstream of those decisions. OpenAPI is last as a cross-check.
