# Stage 2 — Area 5: Transport/HTTP Handlers & Routing

## Current State

**Handler pattern** (e.g. `SetPasswordHandler`, `account_security.go` line 90-151):
- Returns `http.HandlerFunc` (closure over service interface)
- Extracts `userID` from context via `UserIDFromContext(r.Context())`
- Decodes JSON body with `json.NewDecoder(r.Body).Decode(&req)`
- Boundary validation (defense-in-depth) before calling service
- Maps service errors via `MapServiceError(w, err)` or inline `WriteProblem`
- Returns JSON response with appropriate status code

**Error mapping** (`errors.go` line 76-137):
- `MapServiceError` switches on sentinel errors from `account` package
- Each sentinel maps to a specific HTTP status + Problem Details type URI
- Unknown errors → 500 with generic detail (never leaks internals)
- Pattern: `case errors.Is(err, account.ErrXxx): WriteProblem(w, status, type, title, detail)`

**Middleware chain** (`main.go` line 193-194):
```
mux.Handle("/account/security/", RateLimit(rps, burst)(RequireSession(googleVerifyToken)(accountMux)))
```
- `RequireSession` verifies ES256 JWT, injects `userID` into context
- `RateLimit` applies per-IP token bucket

**Route registration** (`main.go` line 186-202):
- `accountMux` is a sub-mux for `/account/security/` prefix
- Each handler registered with `accountMux.HandleFunc("METHOD /path", Handler(svc))`
- Profile route (`GET /account/me`) registered directly on main mux with inline middleware chain

**No admin routes exist.** No `/admin/` prefix, no admin mux, no admin handlers.

**Service interface pattern** (`account_security.go` line 66-73):
```go
type securityService interface {
    SetPassword(ctx context.Context, userID uuid.UUID, email, currentPassword, newPassword string) (bool, error)
    UnlinkGoogle(ctx context.Context, userID uuid.UUID, password string) error
    MfaEnroll(ctx context.Context, userID uuid.UUID) (string, error)
    // ...
}
```
- Narrow interface — only the methods the handler needs
- Tests inject a stub; production uses `*account.Service`

## Requirement

From `08-role-assignment.md`:
- **GET /admin/users**: Admin-only authz checked explicitly at handler level (spec line 25), returns `UserListResponse`
- **POST /admin/users/{userId}/roles**: Admin-only, returns 201 with `User` or 403/404/409
- **DELETE /admin/users/{userId}/roles?role=...**: Admin-only, returns 204 or 403/404/409
- **Authz**: "checked explicitly at the handler level, not only via a query filter" (spec line 25, AGENTS.md golden rule)

## Gap

**New files needed:**

1. **`admin_roles.go`** — handler file for the three endpoints
2. **`admin_roles_test.go`** — unit tests

**New handler functions:**

1. **`AdminUsersHandler(svc adminService) http.HandlerFunc`** — `GET /admin/users`
   - Extract userID from context (RequireSession)
   - Check caller has `role=admin` — needs a way to get caller's roles. Options: (a) service call `GetLoginUserView(callerID)` and check roles, or (b) new service method `HasRole(userID, role)`. The spec says "checked explicitly at the handler level" — but the handler needs the service to check roles.
   - Parse `cursor` and `limit` query params
   - Hard-cap limit at 20 (defense-in-depth)
   - Call `svc.ListUsers(ctx, cursor, limit)`
   - Return `UserListResponse` JSON

2. **`AssignRoleHandler(svc adminService) http.HandlerFunc`** — `POST /admin/users/{userId}/roles`
   - Extract caller userID from context
   - Check caller has `role=admin`
   - Parse `userId` from path (Go 1.22 `r.PathValue("userId")`)
   - Parse JSON body → `AssignRoleRequest`
   - Validate `role` is one of `["admin", "kurator"]`
   - Call `svc.AssignRole(ctx, callerID, targetUserID, role)`
   - Map errors: `ErrUserNotFound` → 404, `ErrRoleConflict` → 409, `ErrDuplicateRole` → 409
   - Return 201 with `User` JSON

3. **`RevokeRoleHandler(svc adminService) http.HandlerFunc`** — `DELETE /admin/users/{userId}/roles?role=...`
   - Extract caller userID from context
   - Check caller has `role=admin`
   - Parse `userId` from path
   - Parse `role` from query string
   - Validate `role` enum
   - Call `svc.RevokeRole(ctx, callerID, targetUserID, role)`
   - Map errors: `ErrUserNotFound` → 404, `ErrLastAdmin` → 409, `ErrPendingCuration` → 409
   - Return 204

**New service interface:**
```go
type adminService interface {
    ListUsers(ctx context.Context, cursor *uuid.UUID, limit int) ([]account.LoginUserView, *uuid.UUID, bool, error)
    AssignRole(ctx context.Context, callerID, targetUserID uuid.UUID, role string) (*account.LoginUserView, error)
    RevokeRole(ctx context.Context, callerID, targetUserID uuid.UUID, role string) error
    GetProfile(ctx context.Context, userID uuid.UUID) (*account.LoginUserView, error) // for admin check
}
```

**Admin authz check pattern:**
- The handler needs to verify the caller has `role=admin`. The existing `RequireSession` only injects `userID` — it doesn't check roles.
- Options: (a) new `RequireAdmin` middleware that checks roles, or (b) inline check in each handler. The spec says "checked explicitly at the handler level" — inline check is more explicit.
- The inline check: call `svc.GetProfile(callerID)`, check `roles` contains "admin". If not → 403.

**Route registration in `main.go`:**
```go
adminMux := http.NewServeMux()
adminMux.HandleFunc("GET /admin/users", transporthttp.AdminUsersHandler(accountSvc))
adminMux.HandleFunc("POST /admin/users/{userId}/roles", transporthttp.AssignRoleHandler(accountSvc))
adminMux.HandleFunc("DELETE /admin/users/{userId}/roles", transporthttp.RevokeRoleHandler(accountSvc))
mux.Handle("/admin/", transporthttp.RateLimit(rps, burst)(
    transporthttp.RequireSession(googleVerifyToken)(adminMux)))
```

## Sniffing

1. **Risk**: The admin authz check (caller has `role=admin`) requires a `GetLoginUserView` call per request. This decrypts the caller's email — unnecessary overhead for an authz check. A lighter alternative: a dedicated `HasRole(userID, role) bool` repository method that just checks `user_roles` without decrypting anything. But this adds another repository method.

2. **Edge cases**: `DELETE /admin/users/{userId}/roles?role=...` — if the `role` query param is missing or invalid, the handler should return 422 (validation error), not 404. The OpenAPI contract defines `role` as required with `Role` enum schema — invalid values need explicit handling.

3. **Miscontext**: The spec says "Admin-only authz is checked explicitly at the handler level, not only via a query filter." This means the handler must check roles, not just rely on middleware. But the handler still needs the service to *provide* the role information. The check is "at the handler level" in the sense that the handler makes the 403 decision, not that it does the DB query itself.

4. **Misleading signals**: `RequireSession` already exists and injects `userID` — looks like adding admin routes is just "add routes behind RequireSession." But RequireSession doesn't check roles. The admin check is an additional layer on top of session verification.

5. **Inconsistency**: The existing `MapServiceError` doesn't have cases for the new role-related sentinels (`ErrRoleConflict`, `ErrDuplicateRole`, `ErrLastAdmin`, `ErrPendingCuration`, `ErrUserNotFound`). These need to be added. The handler could also map them inline (like `UnlinkGoogleHandler` does for its 409 cases), but `MapServiceError` is the preferred pattern.
