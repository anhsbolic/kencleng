# Stage 3 — Solutioning

## Decision 1: Service Method

**Options:**
- A. Thin pass-through: `GetProfile(ctx, userID) (*LoginUserView, error)` that just calls `s.repo.GetLoginUserView`
- B. No service method: handler calls repo directly (breaks pattern)
- C. No service method: handler takes repo as dependency (new pattern)

**Decision: A.** The codebase pattern is clear — handlers depend on a narrow service interface, never the repository directly. Every existing handler follows this. A thin pass-through is consistent and allows future business logic to be added without changing the handler. The method is trivial but the pattern matters.

```go
// GetProfile returns the authenticated user's own profile view.
// Email is decrypted — this is the resource owner's own data.
func (s *Service) GetProfile(ctx context.Context, userID uuid.UUID) (*LoginUserView, error) {
    return s.repo.GetLoginUserView(ctx, userID)
}
```

## Decision 2: Response Struct and JSON Tags

**Options:**
- A. Add JSON tags to `LoginUserView` in `entity.go` — fixes both login and profile endpoints
- B. Define a `userResponse` struct in transport layer with proper tags, map from `LoginUserView`
- C. Use `LoginUserView` as-is (camelCase keys) — violates OpenAPI contract

**Decision: B.** Reasons:
1. `LoginUserView` is a domain entity — JSON serialization is a transport concern, not a domain concern
2. Adding JSON tags to the domain entity couples it to the API contract
3. The login endpoint's camelCase issue is a pre-existing bug — fixing it here would be scope creep for this task
4. A transport-layer response struct keeps the fix scoped and explicit

```go
// userResponse mirrors openapi User schema. Defined in transport layer
// because JSON tag naming (snake_case) is an API concern, not a domain concern.
type userResponse struct {
    ID            string   `json:"id"`
    Name          string   `json:"name"`
    Email         string   `json:"email"`
    EmailVerified bool     `json:"email_verified"`
    Roles         []string `json:"roles"`
    AuthProviders []string `json:"auth_providers"`
    MFAEnabled    bool     `json:"mfa_enabled"`
    CreatedAt     string   `json:"created_at"`
}
```

Mapping function:
```go
func toUserResponse(v *account.LoginUserView) userResponse {
    return userResponse{
        ID:            v.ID.String(),
        Name:          v.Name,
        Email:         v.Email,
        EmailVerified: v.EmailVerified,
        Roles:         v.Roles,
        AuthProviders: v.AuthProviders,
        MFAEnabled:    v.MFAEnabled,
        CreatedAt:     v.CreatedAt.Format(time.RFC3339),
    }
}
```

Note: `uuid.UUID` marshals to string automatically in Go's JSON encoder, but using `.String()` explicitly in the response struct makes the contract clear. `time.Time` marshals to RFC 3339 by default, but explicit formatting in the mapping function makes the contract explicit.

## Decision 3: Handler Pattern

**Decision: Follow existing pattern exactly.**

```go
// profileService is the subset of *account.Service the profile handler
// depends on. *account.Service satisfies it; tests inject a stub.
type profileService interface {
    GetProfile(ctx context.Context, userID uuid.UUID) (*account.LoginUserView, error)
}

// AccountMeHandler handles GET /account/me — returns the authenticated
// user's own profile. No ID parameter; resource is keyed entirely by
// the session (no IDOR surface per threat-model component 5).
func AccountMeHandler(svc profileService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID, ok := UserIDFromContext(r.Context())
        if !ok {
            WriteProblem(w, http.StatusUnauthorized,
                "https://kencleng.dev/errors/unauthorized",
                "Unauthorized", "Authentication required.")
            return
        }

        view, err := svc.GetProfile(r.Context(), userID)
        if err != nil {
            MapServiceError(w, err)
            return
        }
        if view == nil {
            // User deleted after session was issued.
            WriteProblem(w, http.StatusUnauthorized,
                "https://kencleng.dev/errors/unauthorized",
                "Unauthorized", "Authentication required.")
            return
        }

        writeJSON(w, http.StatusOK, toUserResponse(view))
    }
}
```

Key points:
- `UserIDFromContext` check handles missing session (401)
- `view == nil` check handles deleted user (401 — same as missing session)
- `MapServiceError` handles unexpected errors (500)
- `writeJSON` writes the response with correct Content-Type

## Decision 4: Route Wiring

**Decision: Add directly to main mux with its own middleware chain.**

```go
// Account profile (task #07). Same middleware chain as account security.
mux.Handle("GET /account/me",
    transporthttp.RateLimit(rps, burst)(
        transporthttp.RequireSession(googleVerifyToken)(
            transporthttp.AccountMeHandler(accountSvc))))
```

Reasons:
- `GET /account/me` has a different path prefix than `/account/security/*`
- Creating a new sub-mux for `/account/` would require restructuring existing routes
- Direct route on main mux is simple and explicit
- Same middleware chain (RateLimit + RequireSession) as security routes

## Decision 5: Test Strategy

**Decision: Follow existing handler test pattern.**

Test file: `internal/transport/http/account_profile_test.go`

Tests:
1. **`TestAccountMe_Success`** — valid session, user exists → 200 + correct JSON shape with all fields
2. **`TestAccountMe_RequiresAuth`** — no session → 401 Problem Details
3. **`TestAccountMe_UserNotFound`** — valid session but user deleted → 401 Problem Details

Stub service:
```go
type stubProfileService struct {
    view *account.LoginUserView
    err  error
}

func (s *stubProfileService) GetProfile(_ context.Context, _ uuid.UUID) (*account.LoginUserView, error) {
    return s.view, s.err
}
```

Session injection: Use the same pattern as `account_security_test.go` — create a valid ES256 JWT, set it as cookie or Authorization header.

## Decision 6: Service Test

**Decision: Minimal test for the pass-through method.**

File: `internal/domain/account/service_test.go`

Test: `TestGetProfile` — verifies the service method calls `GetLoginUserView` and returns the result. Uses the existing `fakeRepo` which already implements `GetLoginUserView`.

## Summary of Changes

| File | Action | Description |
|------|--------|-------------|
| `internal/domain/account/service.go` | Modify | Add `GetProfile` method |
| `internal/domain/account/service_test.go` | Modify | Add `TestGetProfile` |
| `internal/transport/http/account_profile.go` | Create | Handler + response struct + mapping |
| `internal/transport/http/account_profile_test.go` | Create | Handler tests |
| `cmd/server/main.go` | Modify | Add route registration |

## Risk Note

- **Assumptions made**: `GetLoginUserView` returns correct data for all user states (verified, unverified, MFA-enabled, etc.). This is already tested by `TestGetLoginUserView_AssemblesFields` in `repository_db_integration_test.go`.
- **Edge cases intentionally NOT handled**: None beyond what's specified. The spec says "no invariant is exercised" — this is a pure read.
- **Concurrency assumptions**: None. Read-only endpoint, no state mutation.
- **What is not tested, and why**: The JSON tag mismatch in the login endpoint is not fixed here — it's a pre-existing issue outside this task's scope. The `userResponse` struct ensures `GET /account/me` returns correct snake_case keys.
