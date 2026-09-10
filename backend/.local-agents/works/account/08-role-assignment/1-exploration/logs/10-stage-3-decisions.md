# Stage 3 — Decision Log

## D1: Cross-Domain Tables (INV-account-10 / INV-account-14)

**Context:** INV-account-10 (admin ⊥ representative) and INV-account-14 (pending curation blocks kurator revoke) require querying tables owned by the `organization` domain: `organization_representatives` and `organization_curation_assignments`. Neither table exists yet — no organization migrations are in the repo.

**Options:**

| Option | Description | Pros | Cons |
|--------|-------------|------|------|
| A. Create tables in this feature's migration | New migration `000011` creates both tables with the schema from the ERD | Guards are fully implementable and testable now; no deferred work; ERD schema is finalized | `account` domain creates tables it doesn't own; `organization` domain's future migrations must not re-create them |
| B. Defer INV-account-10/14 enforcement | Implement feature without the cross-domain guards; add TODO notes | No cross-domain schema ownership question; simpler migration | Violates spec (both invariants marked RESOLVED); ships incomplete security surface; requires follow-up task |
| C. Create tables + skip enforcement tests | Create empty tables so repo methods run (return false), but don't test the "user IS a representative" path | Tables exist for future use; code is structurally ready | Untested code paths; false sense of security; violates "every claim must have a named test" rule |

**Decision: Option A.**

Rationale:
- The ERD schema is finalized — the table structures won't change.
- The `account` domain's repository only does `SELECT EXISTS` reads against these tables — it doesn't write to them. This is the same cross-domain read pattern the spec explicitly endorses (invariants.md INV-account-14 line 263-266).
- The `organization` domain's future migrations must use `IF NOT EXISTS` or check the migration number to avoid re-creating these tables. This is a coordination concern, not a technical blocker.
- The alternative (deferring) ships an elevation-of-privilege surface with known guard gaps — unacceptable for a Tier 1 feature.

**Migration shape:**
```sql
-- 000011_create_org_reps_and_curation.up.sql
-- Schema pre-settle for organization domain tables needed by
-- account task #8 (role assignment) INV-account-10/14 guards.
-- organization domain migrations must NOT re-create these tables.

CREATE TABLE IF NOT EXISTS organization_representatives (
    id               UUID PRIMARY KEY,
    user_id          UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    organization_id  UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    level            TEXT NOT NULL CHECK (level IN ('owner', 'staff')),
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (user_id, organization_id)
);
-- indexes...

CREATE TABLE IF NOT EXISTS organization_curation_assignments (
    id               UUID PRIMARY KEY,
    organization_id  UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    kurator_id       UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    assigned_by      UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    assigned_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    decision         TEXT NOT NULL DEFAULT 'pending'
                     CHECK (decision IN ('pending', 'approved', 'rejected')),
    decision_note    TEXT,
    decided_at       TIMESTAMPTZ
);
-- indexes...
```

**Resolved:** The `organizations` table will be created in the same migration using `IF NOT EXISTS`. Migrations can be rearranged later when all domain migrations are available.

---

## D2: Admin Authorization Check Pattern

**Context:** The spec requires "Admin-only authz is checked explicitly at the handler level, not only via a query filter" (spec line 25). The handler needs to verify the caller has `role=admin`. `RequireSession` only injects `userID` — it doesn't check roles.

**Options:**

| Option | Description | Pros | Cons |
|--------|-------------|------|------|
| A. Inline check via `GetLoginUserView` | Handler calls `svc.GetProfile(callerID)`, checks `roles` contains "admin" | Uses existing method; no new code | Decrypts caller's email unnecessarily; `GetProfile` returns full profile just for authz |
| B. New `HasRole(userID, role) bool` repo method | Lightweight query: `SELECT EXISTS(SELECT 1 FROM user_roles WHERE user_id = ? AND role = ?)` | No decrypt overhead; purpose-built for authz | New repo method; more code |
| C. New `RequireAdmin` middleware | Middleware checks roles before handler; rejects with 403 | Reusable; clean separation | Needs role lookup in middleware (crosses into domain layer); less explicit per spec |

**Decision: Option B.**

Rationale:
- The spec says "checked explicitly at the handler level" — Option C (middleware) is less explicit. Option A (GetLoginUserView) works but decrypts email for no reason.
- `HasRole` is a single `SELECT EXISTS` — trivial to implement, trivial to test, no decrypt overhead.
- The handler calls `HasRole(callerID, "admin")` → if false, write 403 and return. This is explicit, visible, and testable.
- The same method can be reused for any future admin-check need.

**Handler pattern:**
```go
func AdminUsersHandler(svc adminService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        callerID, ok := UserIDFromContext(r.Context())
        if !ok { /* 401 */ }

        isAdmin, err := svc.HasRole(r.Context(), callerID, "admin")
        if err != nil { MapServiceError(w, err); return }
        if !isAdmin {
            WriteProblem(w, http.StatusForbidden,
                "https://kencleng.dev/errors/forbidden",
                "Forbidden", "Anda tidak memiliki izin untuk melakukan aksi ini.")
            return
        }
        // ... proceed
    }
}
```

---

## D3: Concurrency Guard for INV-account-09 (Admin ⊥ Kurator)

**Context:** The `user_roles` table structurally allows a user to have both `admin` and `kurator` rows (ERD line 498-504). The exclusivity invariant is behavioral, not a DB constraint. Concurrent assignment of `admin` and `kurator` to the same user must not both succeed.

**Options:**

| Option | Description | Pros | Cons |
|--------|-------------|------|------|
| A. `SELECT ... FOR UPDATE` on `user_roles` | Lock all `user_roles` rows for the target user within the tx; then check exclusivity | Serializable; proven pattern (UnlinkGoogle uses this); simple | Locks all roles for the user (not just the conflicting one); slight over-locking |
| B. Check-then-insert without locking | Read roles, check exclusivity, insert — no lock | Simple; no lock overhead | Race window: two concurrent requests can both pass the check before either inserts |
| C. DB-level trigger | Trigger rejects insert if conflicting role exists | No app-level race possible | ERD explicitly rejects this approach (line 498-504); adds DB complexity; harder to test |

**Decision: Option A.**

Rationale:
- This is the exact same pattern used by `UnlinkGoogle` (`FindAuthIdentitiesByUserForUpdate` — `FOR UPDATE` row locks within tx).
- Under READ COMMITTED, the `FOR UPDATE` lock serializes concurrent assignments: the loser blocks until the winner commits, then re-reads and sees the conflicting role.
- The lock scope is narrow (only the target user's `user_roles` rows, only during the assignment tx).
- Option B has a race window that violates INV-account-09's concurrency test requirement.

**Implementation:**
- New repo method: `FindUserRolesByUserForUpdate(ctx, tx, userID) ([]string, error)` — same pattern as `FindAuthIdentitiesByUserForUpdate`
- Service calls this within tx before inserting the new role
- If conflicting role found → return `ErrRoleConflict`

---

## D4: Concurrency Guard for INV-account-13 (Last Admin)

**Context:** When revoking `admin`, the system must ensure at least one admin remains. Two concurrent revoke requests when exactly 2 admins exist must not both succeed.

**Options:**

| Option | Description | Pros | Cons |
|--------|-------------|------|------|
| A. `SELECT COUNT(*) ... FOR UPDATE` | Lock the counted rows (or use `FOR UPDATE` on a subquery) | Serializable | `FOR UPDATE` on aggregate doesn't lock individual rows in a useful way |
| B. `DELETE ... WHERE role='admin' AND user_id=? AND (SELECT COUNT(*) FROM user_roles WHERE role='admin') > 1` | Atomic delete with count guard in a single statement | No race window; single statement | Complex SQL; harder to read |
| C. Check count + delete within tx, relying on `UNIQUE (user_id, role)` | Count admins in tx → if count <= 1, reject → else delete | Simple; readable | Race window between count and delete |

**Decision: Option B (atomic delete with embedded count).**

Rationale:
- The count check and the delete must be atomic. A separate count-then-delete (Option C) has a race window.
- Option A's `FOR UPDATE` on a COUNT doesn't reliably lock the rows needed.
- Option B uses a single `DELETE` statement with a subquery guard: `DELETE FROM user_roles WHERE user_id = ? AND role = 'admin' AND (SELECT COUNT(*) FROM user_roles WHERE role = 'admin') > 1`. If only 1 admin exists, the `> 1` condition fails, zero rows are deleted, and `RowsAffected() == 0` signals the rejection.
- This is parameterized via goqu (not raw SQL), so it satisfies the AGENTS.md golden rule.

**Implementation:**
```go
// In repository_db.go
func (r *RepositoryDB) DeleteUserRoleIfNotLastAdmin(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (deleted bool, err error) {
    // DELETE FROM user_roles
    // WHERE user_id = ? AND role = 'admin'
    //   AND (SELECT COUNT(*) FROM user_roles WHERE role = 'admin') > 1
    sql, args, err := pgDialect.Delete("user_roles").
        Where(goqu.Ex{"user_id": userID, "role": "admin"},
              goqu.L("(SELECT COUNT(*) FROM user_roles WHERE role = 'admin') > 1")).
        Prepared(true).ToSQL()
    // ...
    tag, err := tx.Exec(ctx, sql, args...)
    return tag.RowsAffected() == 1, nil
}
```

---

## D5: ListUsers Pagination Strategy

**Context:** `GET /admin/users` returns a paginated list of users with their roles. Cursor-based pagination, hard-capped at 20. This is the first paginated endpoint in the account domain.

**Options:**

| Option | Description | Pros | Cons |
|--------|-------------|------|------|
| A. Keyset pagination on `users.id` | `WHERE users.id > cursor ORDER BY users.id LIMIT N+1` | Stable under concurrent inserts/deletes; simple cursor | UUIDv7 ordering may not match creation order (depends on generation) |
| B. Keyset pagination on `users.created_at, users.id` | `WHERE (created_at, id) > (cursor_ts, cursor_id) ORDER BY created_at, id LIMIT N+1` | Matches intuitive ordering; composite cursor | More complex cursor encoding; two values in cursor |
| C. Offset pagination | `OFFSET ? LIMIT ?` | Simple | Unstable under concurrent changes; spec requires cursor-based |

**Decision: Option A.**

Rationale:
- The spec says "cursor-based" (spec line 19). Option C is out.
- `users.id` is UUIDv7 (per ERD/techstack) — naturally time-ordered. Keyset on `id` is equivalent to keyset on `created_at` for UUIDv7.
- Single-value cursor is simpler to encode/decode (just the UUID).
- The `LIMIT N+1` pattern (fetch 21 rows, return 20, use the 21st to determine `hasMore`) avoids a separate COUNT query.

**Implementation:**
```go
// In repository_db.go
func (r *RepositoryDB) ListUsersWithRoles(ctx context.Context, cursor *uuid.UUID, limit int) ([]LoginUserView, *uuid.UUID, bool, error) {
    // 1. SELECT users with cursor + LIMIT (limit+1)
    // 2. For each user, decrypt primary_email
    // 3. For each user, SELECT user_roles
    // 4. Return results, nextCursor, hasMore
}
```

**Optimization note:** The per-user role query (step 3) could be a single `SELECT user_id, role FROM user_roles WHERE user_id IN (...)` batch query instead of N individual queries. This reduces round-trips. But for limit=20, the N-query approach is acceptable and simpler.

---

## D6: Notification for Role Changes

**Context:** The spec says role assign/revoke "triggers a user-facing notification to the target user" (spec line 166-168). The `notification.Sender` interface has `SendNudgeEmail(ctx, to, nudgeType)`. No role-change nudge type exists.

**Decision: Add `NudgeRoleChanged = "role_changed"` constant to `notification/sender.go`.**

Rationale:
- Follows the established pattern (`NudgePasswordReset`, `NudgeGoogleOnly`, `NudgeSetPasswordConflict`).
- `FakeSender` logs the nudge type — no PII leak. Real SMTP implementation comes with the `notification` domain.
- The service calls `s.email.SendNudgeEmail(ctx, targetEmail, notification.NudgeRoleChanged)` after commit.
- The target user's email is obtained from `GetLoginUserView` (called before tx to check existence and get roles).
- Email template deferred to `notification` domain — `FakeSender` logs suffice for v1.

---

## D7: Service Method Signatures

**Decision: Three new service methods + one helper.**

```go
// AssignRole assigns a platform role to a user. Caller must hold role=admin.
// Returns the updated LoginUserView on success.
// Errors: ErrUserNotFound (404), ErrRoleConflict (409), ErrDuplicateRole (409)
func (s *Service) AssignRole(ctx context.Context, callerID, targetUserID uuid.UUID, role string) (*LoginUserView, error)

// RevokeRole revokes a platform role from a user. Caller must hold role=admin.
// Errors: ErrUserNotFound (404), ErrRoleNotFound (404), ErrLastAdmin (409), ErrPendingCuration (409)
func (s *Service) RevokeRole(ctx context.Context, callerID, targetUserID uuid.UUID, role string) error

// ListUsers returns a paginated list of users with their roles.
// Hard-caps limit at 20. Cursor is a user UUID (keyset pagination on users.id).
func (s *Service) ListUsers(ctx context.Context, cursor *uuid.UUID, limit int) ([]LoginUserView, *uuid.UUID, bool, error)

// HasRole reports whether userID currently holds the specified role.
// Lightweight check — no decrypt, no full profile assembly.
func (s *Service) HasRole(ctx context.Context, userID uuid.UUID, role string) (bool, error)
```

**New sentinel errors:**
```go
var (
    ErrRoleConflict      = errors.New("role conflict")       // 409 — INV-account-09/10
    ErrDuplicateRole     = errors.New("duplicate role")      // 409 — already has role
    ErrLastAdmin         = errors.New("last admin")          // 409 — INV-account-13
    ErrPendingCuration   = errors.New("pending curation")    // 409 — INV-account-14
    ErrUserNotFound      = errors.New("user not found")      // 404
    ErrRoleNotFound      = errors.New("role not found")      // 404 — user doesn't have role
)
```

**New action_type constants:**
```go
const (
    actionRoleAssigned = "role_assigned"
    actionRoleRevoked  = "role_revoked"
)
```

---

## D8: Repository Method Signatures

**Decision: 8 new methods on the Repository interface.**

```go
// InsertUserRole inserts a platform role assignment. The UNIQUE (user_id, role)
// constraint may fire — the caller handles 23505 via isUniqueViolation.
InsertUserRole(ctx context.Context, tx pgx.Tx, userID uuid.UUID, role string, grantedBy uuid.UUID) error

// DeleteUserRole removes a platform role. Returns ok=false if the user
// doesn't hold the specified role (zero rows affected).
DeleteUserRole(ctx context.Context, tx pgx.Tx, userID uuid.UUID, role string) (ok bool, err error)

// FindUserRolesByUserForUpdate returns the roles held by userID, taking
// FOR UPDATE row locks within the caller's tx. Serialization point for
// INV-account-09 exclusivity check.
FindUserRolesByUserForUpdate(ctx context.Context, tx pgx.Tx, userID uuid.UUID) ([]string, error)

// CountAdminRoles returns the total number of admin-role rows.
CountAdminRoles(ctx context.Context) (int, error)

// ExistsOrganizationRepresentative reports whether userID has any row
// in organization_representatives (any level, any org). Cross-domain
// read for INV-account-10.
ExistsOrganizationRepresentative(ctx context.Context, userID uuid.UUID) (bool, error)

// ExistsPendingCurationAssignment reports whether userID has a
// 'pending' row in organization_curation_assignments. Cross-domain
// read for INV-account-14.
ExistsPendingCurationAssignment(ctx context.Context, userID uuid.UUID) (bool, error)

// ListUsersWithRoles returns a paginated list of users with their
// roles, using keyset pagination on users.id. Fetches limit+1 rows
// to determine hasMore. Decrypts primary_email per row.
ListUsersWithRoles(ctx context.Context, cursor *uuid.UUID, limit int) ([]LoginUserView, *uuid.UUID, bool, error)

// HasRole reports whether userID holds the specified role.
HasRole(ctx context.Context, userID uuid.UUID, role string) (bool, error)
```

---

## D9: AssignRole Flow (Detailed)

**Pre-condition:** Caller is authenticated and holds `role=admin` (checked by handler).

**Flow:**
1. **Validate role** — must be "admin" or "kurator" (defense-in-depth; handler also validates)
2. **Begin tx**
3. **Check target exists** — `GetLoginUserView(targetUserID)` → nil means 404
4. **Check duplicate** — `FindUserRolesByUserForUpdate(targetUserID)` → if role already in list → `ErrDuplicateRole` (409)
5. **Check INV-account-09** — if role == "admin" AND "kurator" in roles → `ErrRoleConflict` (409)
6. **Check INV-account-10** — if role == "admin" → `ExistsOrganizationRepresentative(targetUserID)` → true means `ErrRoleConflict` (409)
7. **Insert role** — `InsertUserRole(targetUserID, role, callerID)` → if 23505 → `ErrDuplicateRole` (concurrent duplicate fallback)
8. **Insert audit log** — `InsertUserLog(UserLog{actionRoleAssigned, ...})`
9. **Commit tx**
10. **Send notification** — `s.email.SendNudgeEmail(targetEmail, NudgeRoleChanged)` (after commit)
11. **Return** — `GetLoginUserView(targetUserID)` (fresh read for response)

**Error mapping:**
- `ErrUserNotFound` → 404
- `ErrDuplicateRole` → 409 `https://kencleng.dev/errors/duplicate-role`
- `ErrRoleConflict` → 409 `https://kencleng.dev/errors/role-conflict` (distinct detail per INV-account-09 vs INV-account-10)

---

## D10: RevokeRole Flow (Detailed)

**Pre-condition:** Caller is authenticated and holds `role=admin`.

**Flow:**
1. **Validate role** — must be "admin" or "kurator"
2. **Begin tx**
3. **Check target exists** — `GetLoginUserView(targetUserID)` → nil means 404
4. **If role == "admin":**
   - **INV-account-13** — `DeleteUserRoleIfNotLastAdmin(targetUserID)` → if `deleted == false` → `ErrLastAdmin` (409)
5. **If role == "kurator":**
   - **INV-account-14** — `ExistsPendingCurationAssignment(targetUserID)` → true means `ErrPendingCuration` (409)
   - **Delete role** — `DeleteUserRole(targetUserID, "kurator")` → if `ok == false` → `ErrRoleNotFound` (404)
6. **If role == "admin" (reached here means not-last-admin):**
   - The delete already happened in step 4 via `DeleteUserRoleIfNotLastAdmin`
7. **Insert audit log** — `InsertUserLog(UserLog{actionRoleRevoked, ...})`
8. **Commit tx**
9. **Send notification** — `s.email.SendNudgeEmail(targetEmail, NudgeRoleChanged)`

**Note on admin revoke flow:** Steps 4 and 6 need careful ordering. The `DeleteUserRoleIfNotLastAdmin` method atomically checks count AND deletes — it returns `deleted=false` if the user is the last admin. If `deleted=true`, the role is already removed; no separate delete step needed.

**Error mapping:**
- `ErrUserNotFound` → 404
- `ErrRoleNotFound` → 404
- `ErrLastAdmin` → 409 `https://kencleng.dev/errors/last-admin`
- `ErrPendingCuration` → 409 `https://kencleng.dev/errors/pending-curation-assignment`

---

## D11: MapServiceError Additions

**New cases to add to `errors.go` MapServiceError:**

```go
case errors.Is(err, account.ErrUserNotFound):
    WriteProblem(w, http.StatusNotFound,
        "https://kencleng.dev/errors/not-found",
        "Not Found", "User tidak ditemukan.")
case errors.Is(err, account.ErrRoleNotFound):
    WriteProblem(w, http.StatusNotFound,
        "https://kencleng.dev/errors/not-found",
        "Not Found", "User tidak memiliki role ini.")
case errors.Is(err, account.ErrDuplicateRole):
    WriteProblem(w, http.StatusConflict,
        "https://kencleng.dev/errors/duplicate-role",
        "Role Already Assigned", "User ini sudah punya role ini.")
case errors.Is(err, account.ErrRoleConflict):
    WriteProblem(w, http.StatusConflict,
        "https://kencleng.dev/errors/role-conflict",
        "Role Conflict", "Role ini konflik dengan role atau status user saat ini.")
case errors.Is(err, account.ErrLastAdmin):
    WriteProblem(w, http.StatusConflict,
        "https://kencleng.dev/errors/last-admin",
        "Cannot Revoke Last Admin", "Tidak bisa melepas role Admin dari satu-satunya Admin yang tersisa di sistem.")
case errors.Is(err, account.ErrPendingCuration):
    WriteProblem(w, http.StatusConflict,
        "https://kencleng.dev/errors/pending-curation-assignment",
        "Kurator Has Pending Assignment", "User ini masih punya tugas kurasi yang pending. Selesaikan atau alihkan dulu sebelum melepas role Kurator.")
```

**Note:** `ErrRoleConflict` maps to a generic detail string. The handler can override with a more specific detail inline if needed (e.g. for INV-account-09 vs INV-account-10 distinction), following the `MfaEnrollConfirmHandler` pattern where the handler intercepts specific sentinels before falling through to `MapServiceError`.

---

## D12: Route Registration

**New routes in `cmd/server/main.go`:**

```go
// Admin role management (task #08). Behind RequireSession + per-IP rate
// limiter, same as account security. Admin-only authz is checked inline
// at the handler level (spec requirement), not via middleware.
adminMux := http.NewServeMux()
adminMux.HandleFunc("GET /admin/users", transporthttp.AdminUsersHandler(accountSvc))
adminMux.HandleFunc("POST /admin/users/{userId}/roles", transporthttp.AssignRoleHandler(accountSvc))
adminMux.HandleFunc("DELETE /admin/users/{userId}/roles", transporthttp.RevokeRoleHandler(accountSvc))
mux.Handle("/admin/", transporthttp.RateLimit(rps, burst)(
    transporthttp.RequireSession(googleVerifyToken)(adminMux)))
```

**Pattern:** Same as `/account/security/` — sub-mux, rate limit, RequireSession. The admin authz check is inside each handler, not in middleware.
