# Stage 2 — Area 4: Service Layer

## Current State

**Service struct** (`service.go` line 88-127):
- `repo Repository` — persistence port
- `tx TxRunner` — transaction factory (`poolRunner` wraps `*pgxpool.Pool`)
- `breachCheck`, `email`, `keys`, `googleOAuth`, `authKeys`, `frontendURL` — dependencies
- `mfa MfaVerifier`, `now func() time.Time`, `compare`, `hashPassword`, `mintAccess`, `mintMFAPending`, `verifyPending` — seams

**Sentinel errors** (`service.go` line 30-41):
- `ErrValidation` (422), `ErrTokenExpired` (410), `ErrTokenNotFound` (404)
- `security.go` line 27-36: `ErrOnlyIdentity` (409), `ErrRemainingUnverified` (409)

**Transaction pattern** (e.g. `security.go` line 122-167):
```go
tx, err := s.tx.BeginTx(ctx)
committed := false
defer func() { if !committed { _ = tx.Rollback(ctx) } }()
// ... multiple repo calls with tx ...
if err := tx.Commit(ctx); err != nil { ... }
committed = true
```

**Audit logging pattern** (e.g. `security.go` — UnlinkGoogle):
- `InsertUserLog` called within the same tx as the mutation
- `UserLog{ID: uuid.New(), UserID: userID, ActionType: "account_unlinked", CreatedAt: s.now()}`

**Unique violation handling** (`service.go` line 784-789):
- `isUniqueViolation(err) bool` — checks `*pgconn.PgError` code 23505
- Used in `setPasswordBranch1` to map concurrent duplicate to clean no-op

**No admin/role-related service methods exist.** The service has no methods for:
- Assigning roles
- Revoking roles
- Listing users
- Checking admin status

## Requirement

From `08-role-assignment.md`:
- **Assign role** (line 29-31): create `user_roles` row, `granted_by = caller's user_id`, return updated `User`
- **Revoke role** (line 53): delete `user_roles` row, return 204
- **INV-account-09** (line 34-38): assigning `admin` to existing `kurator` → 409
- **INV-account-10** (line 35-36): assigning `admin` to organization representative → 409
- **INV-account-13** (line 61-66): revoking last admin → 409
- **INV-account-14** (line 67-73): revoking kurator with pending curation → 409
- **Duplicate assignment** (line 40-43): already has role → 409
- **Audit log** (line 162-168): write `user_logs` for both assign and revoke
- **Notification** (line 166-168): trigger user-facing notification to target user

## Gap

**3 new service methods needed:**

1. **`AssignRole(ctx, callerID, targetUserID, role string) (*LoginUserView, error)`**
   - Validate target user exists (→ 404)
   - Check INV-account-09: if role == "admin", check target doesn't have "kurator" role
   - Check INV-account-10: if role == "admin", check target isn't organization representative
   - Check duplicate: if target already has role → 409
   - Insert `user_roles` row within tx
   - Insert `user_logs` entry within tx
   - Return updated `LoginUserView` (needs `GetLoginUserView` after commit, or build from tx state)
   - Trigger notification after commit (same pattern as `sendVerification`)

2. **`RevokeRole(ctx, callerID, targetUserID, role string) error`**
   - Validate target user exists (→ 404)
   - Check target has the role (→ 404 if not)
   - Check INV-account-13: if role == "admin", count admin roles → 409 if last
   - Check INV-account-14: if role == "kurator", check pending curation assignments → 409
   - Delete `user_roles` row within tx
   - Insert `user_logs` entry within tx
   - Trigger notification after commit

3. **`ListUsers(ctx, cursor *uuid.UUID, limit int) ([]LoginUserView, *uuid.UUID, bool, error)`**
   - Delegate to `repo.ListUsersWithRoles`
   - Hard-cap limit at 20 (defense-in-depth, even though handler also caps)

**New sentinel errors needed:**
- `ErrRoleConflict` (409) — for INV-account-09/10 exclusivity violations
- `ErrDuplicateRole` (409) — for duplicate assignment
- `ErrLastAdmin` (409) — for INV-account-13
- `ErrPendingCuration` (409) — for INV-account-14
- `ErrUserNotFound` (404) — for target user not found

**New action_type constants for audit log:**
- `"role_assigned"` — for successful assignment
- `"role_revoked"` — for successful revocation

## Sniffing

1. **Risk**: `AssignRole` needs to check INV-account-09 (admin ⊥ kurator) and INV-account-10 (admin ⊥ representative) *before* inserting. But the insert itself has a `UNIQUE (user_id, role)` constraint. If the service checks "does user have kurator?" and then concurrently two requests arrive (one assigning admin, one assigning kurator), the check-then-insert could race. The `UNIQUE` constraint catches the duplicate-role case, but the exclusivity check (admin ⊥ kurator) is a *behavioral* invariant, not a DB constraint. The service needs to either: (a) use `SELECT ... FOR UPDATE` on `user_roles` within the tx, or (b) do the exclusivity check atomically with the insert. This is the concurrency concern from INV-account-09's verification requirement.

2. **Edge cases**: `RevokeRole` for `admin` role — the INV-account-13 count check must be atomic with the delete. If two admins exist and both are revoked concurrently, the count check could pass for both (reading count=2), then both deletes succeed, leaving 0 admins. The count check needs to be within the same tx as the delete, and ideally use `SELECT ... FOR UPDATE` or a similar guard.

3. **Miscontext**: The spec says `AssignRole` returns the updated `User` (201 response). But `LoginUserView` includes decrypted email — the same PII concern as `ListUsers`. The assign endpoint decrypts email for a single user (cheap), not a bulk list.

4. **Misleading signals**: The existing `GetLoginUserView` already reads roles — looks like the "return updated User" is trivial. But `GetLoginUserView` uses `r.db` (pool, not tx). If called within the assign tx, it would need a tx-aware variant. Alternatively, call it after commit (outside tx).

5. **Inconsistency**: The spec says "response 201 with the updated User" (line 31). The OpenAPI contract returns `User` schema (not `LoginUserView`). The `User` schema includes `email`, `roles`, `auth_providers`, `mfa_enabled` — same shape as `LoginUserView`. The service can return `*LoginUserView` and the handler maps it to the `User` JSON shape.
