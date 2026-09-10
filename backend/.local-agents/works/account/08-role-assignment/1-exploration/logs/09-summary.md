# Stage 2 — Summary

## Areas Explored (7)

| # | Area | Key Gap | Severity |
|---|------|---------|----------|
| 1 | OpenAPI Contract | No gap — contract is complete | — |
| 2 | Entities & Repository Interface | 7 missing repo methods, no UserRole entity | High |
| 3 | Repository DB Adapter | 7 new methods needed, pagination pattern new | High |
| 4 | Service Layer | 3 new methods, 5 new sentinels, 2 new action_types | High |
| 5 | Transport/HTTP Handlers | 3 new handlers, new adminMux, admin authz check | High |
| 6 | Cross-Domain Dependencies | `organization_representatives` and `organization_curation_assignments` tables don't exist yet | **Blocker** |
| 7 | Existing Patterns | New notification type, concurrency guard approach | Medium |

## Critical Finding: Cross-Domain Table Gap

The INV-account-10 and INV-account-14 guards require tables that don't exist yet:
- `organization_representatives` — needed for "admin ⊥ representative" check
- `organization_curation_assignments` — needed for "pending curation blocks kurator revoke"

**Decision needed:** Create these tables as part of this feature's migration, or defer the guards.

## Files to Create/Modify

**New files:**
- `internal/domain/account/role.go` — service methods (AssignRole, RevokeRole, ListUsers)
- `internal/domain/account/role_test.go` — unit tests
- `internal/transport/http/admin_roles.go` — handlers
- `internal/transport/http/admin_roles_test.go` — handler tests
- `migrations/000011_create_org_reps_and_curation.up.sql` — if tables created here
- `migrations/000011_create_org_reps_and_curation.down.sql`

**Files to modify:**
- `internal/domain/account/repository.go` — add 7 new methods to interface
- `internal/domain/account/repository_db.go` — implement 7 new methods
- `internal/domain/account/service.go` — add sentinels, action_type constants
- `internal/transport/http/errors.go` — add new sentinels to MapServiceError
- `cmd/server/main.go` — register admin routes
- `internal/platform/notification/sender.go` — add NudgeRoleChanged constant

## Concurrency Concerns

Two concurrency-sensitive areas identified:
1. **INV-account-09 exclusivity**: concurrent `admin` + `kurator` assignment to same user — needs `FOR UPDATE` lock on `user_roles` within tx
2. **INV-account-13 last-admin**: concurrent revoke of two admins when only 2 exist — needs atomic count+delete within tx
