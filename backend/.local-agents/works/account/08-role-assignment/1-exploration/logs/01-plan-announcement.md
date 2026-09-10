# Stage 1 — Plan Announcement

## Task Understanding

Build three Admin-only endpoints for role assignment: `GET /admin/users` (paginated user list, hard-capped at 20), `POST /admin/users/{userId}/roles` (assign admin/kurator with exclusivity guards), and `DELETE /admin/users/{userId}/roles` (revoke with last-admin and pending-curation guards). This is the account domain's only elevation-of-privilege surface — Tier 1 due to concurrency-safe exclusivity checks (INV-account-09/10), mandatory audit logging (INV-account-11), and cross-domain reads (INV-account-14).

## Areas to Explore

1. **Domain entities & repository interface** (`internal/domain/account/entity.go`, `repository.go`)
2. **Repository DB adapter** (`internal/domain/account/repository_db.go`)
3. **Service layer** (`internal/domain/account/service.go`)
4. **Transport/HTTP handlers & routing** (`internal/transport/http/`, `cmd/server/main.go`)
5. **OpenAPI contract** (`api/openapi.yaml` — schemas, endpoints, error responses)
6. **Cross-domain dependencies** (`organization_representatives`, `organization_curation_assignments` tables)
7. **Existing patterns** (how other features handle authz, audit logging, error mapping)

## Order & Rationale

1. **OpenAPI contract first** — defines the exact request/response shapes, error types, and status codes. This is the source of truth for what the handlers must produce.
2. **Domain entities & repository interface** — what structs and persistence methods exist today vs what needs to be added.
3. **Repository DB adapter** — how queries are built (goqu patterns), what's already implemented for user_roles.
4. **Service layer** — where business logic (exclusivity checks, audit logging) lives; what sentinel errors exist.
5. **Transport/HTTP handlers & routing** — how other handlers are structured, how authz is enforced, how routes are registered.
6. **Cross-domain dependencies** — what tables from `organization` domain are needed for INV-account-10/14 guards.
7. **Existing patterns** — how similar features (account-linking, MFA) handle the same concerns (audit log writes, error mapping, concurrency).
