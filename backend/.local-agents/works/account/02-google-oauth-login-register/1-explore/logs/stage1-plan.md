# Stage 1 — Plan Announcement

> Feature: 02-google-oauth-login-register
> Date: 2026-08-22

## Areas to explore

### Area 1: Domain service layer (`service.go`, `service_test.go`)
**Why first:** Core OAuth business logic lives here. Need to understand existing patterns (TxRunner, breachChecker, notification.Sender), what already exists for Google (providerGoogle constant, FindAuthIdentityByIdentifierHash), and what new methods are needed. Test patterns (fakeRepo, fakeTx) also need understanding.

### Area 2: Domain repository layer (`repository.go`, `repository_db.go`)
**Why second:** Repository interface defines available DB operations. Need to check if new queries are needed (FindUserByID for link intent, InsertAuthIdentity for attaching identity to existing user).

### Area 3: Transport/HTTP layer (`auth_register.go`, `middleware.go`, `errors.go`)
**Why third:** Handler patterns (service wrapping, error mapping, response writing) and middleware (rate limiting, auth for link/reauth). Redirect handlers are non-JSON (302 + cookie) — different pattern.

### Area 4: Platform dependencies (`auth/`, `crypto/`)
**Why fourth:** OAuth flow needs new platform capabilities: Google OAuth client, JWKS verification, cookie management. Need to check what exists in platform/auth/ (currently only keys.go).

### Area 5: Entry point wiring (`cmd/server/main.go`)
**Why fifth:** How account domain is wired, what new config needed (env vars), route registration pattern.

### Area 6: API contract (`api/openapi/account.yaml`)
**Why last:** Verification pass — confirm spec expectations match code findings.

## Order rationale

Dependency-ordered: service depends on repository, transport depends on service, platform is shared infra. Entry point wiring depends on all prior areas. API contract is verification.
