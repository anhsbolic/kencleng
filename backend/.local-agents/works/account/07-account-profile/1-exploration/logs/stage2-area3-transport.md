# Stage 2 — Area 3: Transport Layer

## Current State

No handler exists for `GET /account/me`. The transport layer has:

- **`account_security.go`**: Contains `RequireSession` middleware (line 39), `UserIDFromContext` helper (line 23), and handlers for security endpoints (SetPassword, UnlinkGoogle, MFA). Defines `securityService` interface (line 66) — a narrow subset of `*account.Service`.

- **`auth_login.go`**: Contains `writeJSON` helper (line 58), `loginOKResponse` struct (line 37) that embeds `*account.LoginUserView` as `User` field. Shows the pattern for returning user data in JSON.

- **`errors.go`**: Contains `WriteProblem`, `WriteValidationError`, `MapServiceError` — shared error response helpers.

**Handler pattern observed:**
1. Handler function takes a narrow service interface (e.g., `securityService`)
2. Returns `http.HandlerFunc`
3. Extracts `userID` from context via `UserIDFromContext`
4. Calls service method
5. Writes JSON response via `writeJSON` or error via `MapServiceError`

**Test pattern observed:**
1. Stub service struct implementing the narrow interface
2. `httptest.NewRequest` + `httptest.NewRecorder`
3. Inject session user via context or JWT cookie
4. Assert status code, response body shape

## Requirement

Spec says: `GET /account/me` returns `200` with `User` schema on success, `401` on missing/expired/invalid token. No request body, no parameters.

## Gap

Need to create:
1. **New handler file** (e.g., `account_profile.go`) with:
   - Narrow service interface (`profileService`) with one method: `GetProfile(ctx, userID) (*LoginUserView, error)`
   - `AccountMeHandler(svc profileService) http.HandlerFunc`
   - Response struct (or reuse `loginOKResponse.User` pattern — but the spec says just the `User` fields, not wrapped in a login response)

2. **New test file** (e.g., `account_profile_test.go`) with:
   - Stub service
   - Tests for: success (200 + correct JSON shape), missing session (401), user not found (401)

**Response shape question:** The OpenAPI spec says `GET /account/me` returns `User` directly (not wrapped in a `LoginResponse`). The handler should write the `LoginUserView` fields directly as the JSON root, not nested under a `"user"` key.

## Sniffing

1. **Risk**: None. This is a simple read handler with no business logic. The only risk is getting the JSON shape wrong (wrapping vs not wrapping).

2. **Edge cases**: 
   - User deleted after session issued → `GetLoginUserView` returns `(nil, nil)` → handler should return 401 (session invalid)
   - Empty roles/auth_providers → `LoginUserView` initializes these as empty slices, so JSON will be `[]` not `null` — correct behavior

3. **Miscontext**: The `loginOKResponse` struct wraps `User` under a `"user"` key. The `GET /account/me` response should NOT use this wrapper — the spec says the response IS the `User` schema, not a container holding it.

4. **Misleading signals**: The `loginOKResponse.User` field type is `*account.LoginUserView` — looks like we could reuse it, but the wrapper shape is wrong for this endpoint.

5. **Inconsistency**: None identified. The pattern is consistent with existing handlers.
