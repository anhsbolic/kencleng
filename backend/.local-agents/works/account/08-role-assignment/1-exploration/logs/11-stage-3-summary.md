# Stage 3 — Summary

## 12 Decisions Made

| # | Decision | Choice | Key Rationale |
|---|----------|--------|---------------|
| D1 | Cross-domain tables | Create in this feature's migration | ERD is finalized; guards must be testable; `account` only reads, doesn't write |
| D2 | Admin authz check | New `HasRole` repo method | Lightweight; no decrypt overhead; explicit at handler level per spec |
| D3 | INV-account-09 concurrency | `FOR UPDATE` on `user_roles` | Proven pattern (UnlinkGoogle); serializes concurrent assignments |
| D4 | INV-account-13 concurrency | Atomic `DELETE` with embedded `COUNT` subquery | No race window; single statement; goqu-parameterized |
| D5 | Pagination strategy | Keyset on `users.id` (UUIDv7) | Cursor-based per spec; single-value cursor; `LIMIT N+1` for hasMore |
| D6 | Notification type | New `NudgeRoleChanged` constant | Follows established pattern; FakeSender logs in v1 |
| D7 | Service method signatures | 3 methods + `HasRole` | AssignRole, RevokeRole, ListUsers, HasRole |
| D8 | Repository method signatures | 8 new methods | Covers all query needs: CRUD, guards, cross-domain reads, pagination |
| D9 | AssignRole flow | 11-step tx flow | Guards first, insert, audit log, commit, notify after |
| D10 | RevokeRole flow | Atomic delete-with-count for admin; check-then-delete for kurator | Different guard strategies per role |
| D11 | MapServiceError additions | 6 new sentinel mappings | Covers all 409/404 cases with distinct Problem types |
| D12 | Route registration | New `adminMux` with RequireSession | Same pattern as account security; admin check inline |

## Implementation Order (Recommended)

1. **Migration** — create `organization_representatives` and `organization_curation_assignments` tables (D1)
2. **Repository interface** — add 8 methods to `Repository` (D8)
3. **Repository DB adapter** — implement 8 methods in `repository_db.go`
4. **Service layer** — add sentinels, constants, 4 methods (D7, D9, D10)
5. **Error mapping** — add 6 cases to `MapServiceError` (D11)
6. **Handlers** — create `admin_roles.go` with 3 handlers (D2)
7. **Route registration** — add `adminMux` to `main.go` (D12)
8. **Notification** — add `NudgeRoleChanged` constant (D6)
9. **Tests** — unit tests for service + handlers; integration tests for concurrency (D3, D4)

## Open Items — Resolved

- **`organizations` table FK** — **Resolved:** Create full `organizations` table in the same migration (`000011`) using `IF NOT EXISTS`. The `organization` domain's future migrations will skip re-creation. Migrations can be rearranged later when all domain migrations are available.
- **Notification email content** — **Resolved:** Add `NudgeRoleChanged = "role_changed"` constant now. `FakeSender` logs suffice for v1. Email template deferred to `notification` domain.
