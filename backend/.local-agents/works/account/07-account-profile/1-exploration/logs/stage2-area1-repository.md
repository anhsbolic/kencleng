# Stage 2 — Area 1: Repository Layer

## Current State

`GetLoginUserView` at `repository_db.go:679` assembles a `LoginUserView` struct by:

1. **Profile row** (lines 687-708): Queries `users` for `name`, `primary_email` (ciphertext), `created_at`. Decrypts email via `crypto.Decrypt`. Returns `(nil, nil)` if user not found.

2. **Identity aggregation** (lines 711-741): Queries `auth_identities` for `provider_type` + `verified_at`. Builds distinct `AuthProviders` list. Sets `EmailVerified = true` if any `email_password` identity has `verified_at` set.

3. **Roles** (lines 744-767): Queries `user_roles` ordered by `role ASC`. Empty until task #8 ships.

4. **MFA flag** (lines 772-784): Counts `mfa_totp_secrets` rows where `enabled_at IS NOT NULL`. Sets `MFAEnabled = count > 0`.

The `LoginUserView` struct (`entity.go:118-127`) has exactly the fields the spec requires:
- `ID uuid.UUID` → maps to OpenAPI `id` (uuid)
- `Name string` → maps to `name`
- `Email string` → maps to `email` (decrypted plaintext, never logged)
- `EmailVerified bool` → maps to `email_verified`
- `Roles []string` → maps to `roles` (array of Role)
- `AuthProviders []string` → maps to `auth_providers` (array of ProviderType)
- `MFAEnabled bool` → maps to `mfa_enabled`
- `CreatedAt time.Time` → maps to `created_at`

The method is on `*RepositoryDB` (concrete), not behind the `Repository` interface. It's already used by the login flow (`login.go:483` — `issueSessionTokens` calls it).

## Requirement

Spec says: response is the caller's own `User` resource with `id`, `name`, `email` (decrypted), `email_verified`, `roles`, `auth_providers`, `mfa_enabled`, `created_at`. `200`.

## Gap

**None.** `GetLoginUserView` already returns exactly the fields needed, with email decrypted. The `LoginUserView` struct maps 1:1 to the OpenAPI `User` schema. No new SQL or repository method is required.

## Sniffing

1. **Risk**: None identified. This is a pure read with no state mutation. The query is keyed by `userID` from the session — no user-supplied identifier to tamper with.

2. **Edge cases**: User not found returns `(nil, nil)` — handler must check for nil and return 401 (session references a deleted user). This is an unlikely but valid edge case.

3. **Miscontext**: None. The spec explicitly says "decrypted, since this is always the resource owner viewing their own data" — matches `GetLoginUserView`'s decrypt-on-read behavior exactly.

4. **Misleading signals**: None. The method is already used in production login flow, not dead code.

5. **Inconsistency**: None. The `LoginUserView` doc comment (`entity.go:110-117`) explicitly describes this as the read model for `User` schema, and notes it's "the one place the repository adapter decrypts primary_email on read."
