# Stage 2 — Area 5: OpenAPI Contract

## Current State

**Endpoint definition** (`openapi.yaml:425-438`):
```yaml
/account/me:
  get:
    tags: [account]
    summary: Get the current authenticated user's profile
    responses:
      '200':
        description: OK
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/User'
      '401':
        $ref: '#/components/responses/Unauthorized'
```

**User schema** (`openapi.yaml:2448-2477`):
- `id`: string (uuid)
- `name`: string
- `email`: string (email format)
- `email_verified`: boolean
- `roles`: array of `Role` (enum: `admin`, `kurator`)
- `auth_providers`: array of `ProviderType` (enum: `email_password`, `google`)
- `mfa_enabled`: boolean
- `created_at`: string (date-time)

**LoginUserView struct** (`entity.go:118-127`):
- `ID uuid.UUID` → `id`
- `Name string` → `name`
- `Email string` → `email`
- `EmailVerified bool` → `email_verified`
- `Roles []string` → `roles`
- `AuthProviders []string` → `auth_providers`
- `MFAEnabled bool` → `mfa_enabled`
- `CreatedAt time.Time` → `created_at`

**JSON tag mapping needed:** The `LoginUserView` struct has no JSON tags. The handler's response struct must add `json:"id"`, `json:"name"`, etc. to match the OpenAPI snake_case convention.

## Requirement

Response must match the `User` schema exactly: snake_case JSON keys, correct types. The `LoginUserView` struct lacks JSON tags, so a transport-layer response struct with proper tags is needed.

## Gap

1. **JSON tags**: `LoginUserView` has NO JSON tags (`entity.go:118-127`). Go's default `encoding/json` marshals without tags uses the field name as-is: `ID`, `Name`, `Email`, `EmailVerified`, `Roles`, `AuthProviders`, `MFAEnabled`, `CreatedAt`. The OpenAPI spec expects snake_case: `id`, `name`, `email`, `email_verified`, `roles`, `auth_providers`, `mfa_enabled`, `created_at`.

   The login response (`auth_login.go:41`) embeds `*account.LoginUserView` directly — it has the same mismatch. This is a **pre-existing contract violation** in the login endpoint, not introduced by this task.

   For `GET /account/me`, the handler should define a `userResponse` struct in the transport layer with proper JSON tags, mapping from `LoginUserView`. This keeps the fix scoped to this endpoint without touching the domain entity or the login response.

2. **Response codes**: Spec says `200` (success) and `401` (auth failure). The `RequireSession` middleware already returns 401 for missing/invalid tokens. The handler needs to return 200 on success.

## Sniffing

1. **Risk**: JSON serialization mismatch is a real finding. The `LoginUserView` struct produces camelCase keys (`ID`, `MFAEnabled`) instead of snake_case (`id`, `mfa_enabled`). The login endpoint has the same issue. For this task, defining a transport-layer response struct with correct tags is the fix.

2. **Edge cases**: Empty `Roles` and `AuthProviders` slices serialize as `[]` (not `null`) because `GetLoginUserView` initializes them as empty slices. This is correct behavior.

3. **Miscontext**: The spec says `email` is "decrypted, since this is always the resource owner viewing their own data" — matches `GetLoginUserView`'s behavior exactly.

4. **Misleading signals**: The login response already returns `LoginUserView` — looks like it works, but the JSON keys are camelCase, not snake_case. The contract is already broken there.

5. **Inconsistency**: The OpenAPI spec expects snake_case keys. The current `LoginUserView` serialization produces camelCase. This is inconsistent. The fix for this task is a transport-layer response struct; the login endpoint's mismatch is a separate pre-existing issue.
