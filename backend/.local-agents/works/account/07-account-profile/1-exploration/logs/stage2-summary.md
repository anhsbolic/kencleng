# Stage 2 — Consolidated Gap Summary

## Areas Explored

1. **Repository layer** — No gap. `GetLoginUserView` already returns all required fields with decrypted email.

2. **Service layer** — Gap: need a thin `GetProfile(ctx, userID) (*LoginUserView, error)` method on `*Service` (pass-through to `GetLoginUserView`).

3. **Transport layer** — Gap: need new handler file with:
   - `profileService` interface (one method: `GetProfile`)
   - `AccountMeHandler` function
   - `userResponse` struct with proper JSON tags (snake_case)
   - Tests (stub service, success/401 cases)

4. **Route wiring** — Gap: need to add `GET /account/me` route on main mux with `RateLimit` + `RequireSession` middleware chain.

5. **OpenAPI contract** — Gap: `LoginUserView` has no JSON tags, produces camelCase keys. Need transport-layer response struct with snake_case tags. Pre-existing issue in login endpoint (separate concern).

## Key Findings

- **No new SQL needed** — `GetLoginUserView` covers everything
- **No new repository method needed** — already on the `Repository` interface
- **JSON tag mismatch** — `LoginUserView` produces `ID`, `MFAEnabled`, `EmailVerified` instead of `id`, `mfa_enabled`, `email_verified`. Fix is a transport-layer response struct.
- **User-not-found edge case** — `GetLoginUserView` returns `(nil, nil)` for deleted users. Handler must check and return 401.
- **Response shape** — `GET /account/me` returns `User` directly (not wrapped in a container)

## Files to Create/Modify

**Create:**
- `internal/transport/http/account_profile.go` — handler + response struct
- `internal/transport/http/account_profile_test.go` — tests

**Modify:**
- `internal/domain/account/service.go` — add `GetProfile` method
- `internal/domain/account/service_test.go` — add test for `GetProfile`
- `cmd/server/main.go` — add route registration
