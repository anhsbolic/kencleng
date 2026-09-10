# Stage 2 — Area 2: Domain Entities & Repository Interface

## Current State

**Entities** (`entity.go`):
- `User` (line 29): `{ID, Name, PrimaryEmail, CreatedAt, UpdatedAt}` — no Roles field
- `LoginUserView` (line 118): `{ID, Name, Email, EmailVerified, Roles []string, AuthProviders []string, MFAEnabled, CreatedAt}` — Roles field exists, comment says "empty until account task #8 ships assignment"
- `UserLog` (line 133): `{ID, UserID, ActionType, CreatedAt}` — exists, ready for audit writes
- No `UserRole` entity exists — the `user_roles` table is only referenced in SQL comments

**Repository interface** (`repository.go`):
- `GetLoginUserView` (line 219): already reads from `user_roles` (see `repository_db.go` line 744) — returns roles as `[]string`
- `InsertUserLog` (line 225): exists, takes `*UserLog` within a `pgx.Tx`
- No methods exist for:
  - Inserting a `user_roles` row (assign role)
  - Deleting a `user_roles` row (revoke role)
  - Querying `user_roles` for a specific user's roles (needed for exclusivity checks)
  - Counting admin roles (needed for INV-account-13 last-admin guard)
  - Checking `organization_representatives` (needed for INV-account-10)
  - Checking `organization_curation_assignments` (needed for INV-account-14)
  - Listing users with pagination (needed for `GET /admin/users`)

## Requirement

From `08-role-assignment.md`:
- **Assign** (line 29-31): create a `user_roles` row with `granted_by = caller's user_id`
- **Revoke** (line 53): delete the `user_roles` row
- **INV-account-09** (invariants.md line 146-155): Admin ⊥ Kurator — check both directions
- **INV-account-10** (invariants.md line 157-180): Admin ⊥ organization_representatives — check on assign
- **INV-account-13** (invariants.md line 220-244): at least one Admin always exists — count check on revoke
- **INV-account-14** (invariants.md line 246-283): revoke kurator requires no pending curation assignments
- **Audit log** (spec line 162-168): write `user_logs` entry for both assign and revoke

## Gap

**Missing repository methods (7 new methods needed):**

1. **`InsertUserRole(ctx, tx, userID, role, grantedBy)`** — insert into `user_roles`. Needs to handle `UNIQUE (user_id, role)` constraint violation cleanly (return a distinguishable error for duplicate detection).
2. **`DeleteUserRole(ctx, tx, userID, role)`** — delete from `user_roles` where `(user_id, role)` matches. Returns `ok=false` if no row deleted (user doesn't have role → 404).
3. **`FindUserRolesByUser(ctx, userID)`** — list roles for a user. Needed by assign to check INV-account-09 (does user have kurator? → can't assign admin).
4. **`CountAdminRoles(ctx)`** — `COUNT(*) FROM user_roles WHERE role = 'admin'`. Needed by revoke for INV-account-13 last-admin guard.
5. **`ExistsOrganizationRepresentative(ctx, userID)`** — `SELECT EXISTS(...)` against `organization_representatives`. Needed by assign for INV-account-10.
6. **`ExistsPendingCurationAssignment(ctx, userID)`** — `SELECT EXISTS(...)` against `organization_curation_assignments WHERE kurator_id = ? AND decision = 'pending'`. Needed by revoke for INV-account-14.
7. **`ListUsersWithRoles(ctx, cursor, limit)`** — paginated user list with roles, for `GET /admin/users`. Needs to decrypt `primary_email` (same as `GetLoginUserView`), join with `user_roles`, and apply cursor-based pagination.

**Missing entity:**
- No `UserRole` struct. The `user_roles` table has `{id, user_id, role, granted_by, granted_at}` — may need a struct if the service layer needs to reason about role rows (e.g. to check `granted_by` for audit). However, for the current feature, the service only needs to know *whether* a role exists, not the full row. A struct may not be needed — the repository methods can return booleans/counts.

## Sniffing

1. **Risk**: `InsertUserRole` must handle the `UNIQUE (user_id, role)` constraint violation. If the service does a check-then-insert (does user already have role?), a concurrent duplicate request could still hit the constraint. The service needs to either: (a) catch the unique violation and map it to 409, or (b) use `INSERT ... ON CONFLICT DO NOTHING` and check the result. The existing codebase uses `pgconn.PgError` to catch specific PG error codes (see `service.go` line 16 for the import).
2. **Edge cases**: `ListUsersWithRoles` with cursor pagination — if a user's roles change between pages (e.g. a role is assigned/revoke while the admin is paginating), the cursor may skip or duplicate users. This is an inherent cursor-pagination trade-off, acceptable for an admin-only tool.
3. **Miscontext**: The spec says `GetLoginUserView` already reads `user_roles` (line 744 of `repository_db.go`). The comment in `entity.go` line 115 says "empty until account task #8" — this is misleading: the *query* already runs, but the table was empty because no assignment API existed. Once this feature ships, `GetLoginUserView` will return real roles without any code change.
4. **Misleading signals**: `LoginUserView.Roles` exists and is populated by `GetLoginUserView` — looks like roles are "already working." But the roles are only populated for the login path; there's no way to *write* roles yet. The read path is ready; the write path is entirely missing.
5. **Inconsistency**: The `UserLog.ActionType` comment (line 131) says "the full action_type vocabulary are owned by the user-logs task (#08)." This is the *same* task number as role assignment. The spec says role assignment is one of the mandatory write-sites for `user_logs`. The action_type values for role assign/revoke need to be defined (e.g. `"role_assigned"`, `"role_revoked"`).
