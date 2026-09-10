# Stage 2 — Area 3: Repository DB Adapter

## Current State

**Structure** (`repository_db.go`):
- `RepositoryDB` struct (line 30): `{db *pgxpool.Pool, keys *crypto.Keys}`
- Uses `goqu.Dialect("postgres")` for all SQL (line 23)
- All insert methods take `pgx.Tx` for transactional writes
- All read methods use `r.db` (pool) for standalone queries
- PII encryption pattern: `crypto.Encrypt` at insert, `crypto.Decrypt` at read, `crypto.HMAC` for hash columns

**Existing patterns for error handling:**
- `InsertUser` (line 51): wraps `*pgconn.PgError` with code 23505 (unique violation) — returns it up the chain without swallowing
- `InsertAuthIdentity` (line 83): same pattern — unique violation bubbles up
- `isUniqueViolation` helper in `service.go` (line 784): `errors.As(err, &pgErr) && pgErr.Code == "23505"`

**Existing user_roles read** (`GetLoginUserView`, line 743-767):
- Queries `user_roles` with `SELECT role WHERE user_id = ? ORDER BY role ASC`
- Returns roles as `[]string` in `LoginUserView.Roles`
- Uses `r.db` (pool, not tx) — read-only, no transaction needed

**Pagination pattern:**
- No existing cursor-based pagination in the account domain. Other domains (if any) would need to be checked. The `GET /admin/users` endpoint is the first paginated list endpoint in this domain.

## Requirement

From `08-role-assignment.md` and invariants:
- **Insert user_roles row** (assign): `INSERT INTO user_roles (id, user_id, role, granted_by, granted_at)` — spec line 31
- **Delete user_roles row** (revoke): `DELETE FROM user_roles WHERE user_id = ? AND role = ?` — spec line 53
- **Count admin roles** (INV-account-13): `COUNT(*) FROM user_roles WHERE role = 'admin'` — invariants.md line 222
- **Check organization_representatives** (INV-account-10): `SELECT EXISTS(...)` — invariants.md line 159
- **Check organization_curation_assignments** (INV-account-14): `SELECT EXISTS(...)` — invariants.md line 249
- **Paginated user list** (GET /admin/users): cursor-based, hard-capped at 20, decrypt email per row — spec line 18-22

## Gap

**7 new repository methods needed** (matching Area 2's gap):

1. **`InsertUserRole`**: New method. Pattern matches `InsertUserLog` — goqu insert into `user_roles`, takes `pgx.Tx`. Must let unique violation (23505) bubble up so the service can map it to 409.

2. **`DeleteUserRole`**: New method. `DELETE FROM user_roles WHERE user_id = ? AND role = ?` — returns `(ok bool, err error)` where `ok = false` when zero rows affected (user doesn't have role → 404). Pattern: `Tag.RowsAffected()` check after `tx.Exec`.

3. **`FindUserRolesByUser`**: New method. `SELECT role FROM user_roles WHERE user_id = ?` — returns `[]string`. Similar to the existing query in `GetLoginUserView` lines 744-767 but as a standalone method.

4. **`CountAdminRoles`**: New method. `SELECT COUNT(*) FROM user_roles WHERE role = 'admin'` — returns `int`. Uses `r.db` (pool, read-only).

5. **`ExistsOrganizationRepresentative`**: New method. `SELECT EXISTS(SELECT 1 FROM organization_representatives WHERE user_id = ?)` — returns `bool`. Uses `r.db`. Cross-domain read against a table `organization` owns.

6. **`ExistsPendingCurationAssignment`**: New method. `SELECT EXISTS(SELECT 1 FROM organization_curation_assignments WHERE kurator_id = ? AND decision = 'pending')` — returns `bool`. Uses `r.db`. Cross-domain read.

7. **`ListUsersWithRoles`**: New method. Most complex — needs:
   - Cursor-based pagination on `users.id` (or `users.created_at`)
   - Join with `user_roles` to get roles per user
   - Decrypt `primary_email` per row (same pattern as `GetLoginUserView` line 704)
   - Hard cap at `limit = 20` (enforced at query level, not just handler)
   - Return `([]LoginUserView, nextCursor *uuid.UUID, hasMore bool, err error)`

## Sniffing

1. **Risk**: `ListUsersWithRoles` decrypts email for every row — this is the exact concern in threat-model component 6. The hard cap of 20 bounds the per-request decrypt cost. But the query itself should use `LIMIT 21` (or `LIMIT $1 + 1`) to detect `hasMore` without a separate COUNT query — the existing pagination pattern in other domains should be checked.

2. **Edge cases**: `DeleteUserRole` with `Tag.RowsAffected()` — if the user exists but doesn't have the role, `RowsAffected() == 0`. This maps to 404 per spec line 55-59. But what if the user doesn't exist at all? The spec says both cases collapse to 404 (line 55-59), so no separate user-existence check is needed.

3. **Miscontext**: The `organization_representatives` and `organization_curation_assignments` tables exist in the DB (migrations exist or will exist by the time this is built). The repository adapter just queries them — it doesn't own them. This is the cross-domain read pattern from INV-account-10/14.

4. **Misleading signals**: `GetLoginUserView` already queries `user_roles` — looks like the read path is done. But it only returns roles for a single user. The `ListUsersWithRoles` method needs a different query shape (paginated list with joins), not a reuse of `GetLoginUserView`.

5. **Inconsistency**: The existing `InsertUser` and `InsertAuthIdentity` methods let unique violations bubble up as raw `*pgconn.PgError`. The service uses `isUniqueViolation` to detect them. For `InsertUserRole`, the same pattern applies — but the service needs to distinguish "user already has this role" (409 duplicate) from other unique violations. Since `user_roles` only has one unique constraint `(user_id, role)`, any 23505 from this table is a duplicate role assignment.
