# Stage 2 — Gap Analysis, Area 2: DB schema

> Files: `migrations/000002_create_auth_identities.*`,
> `000003_create_auth_tokens.*`, `000004_create_refresh_tokens.*`,
> `000005_create_user_logs.*`

## Current state (concrete)

- **`auth_identities`**: `verified_at TIMESTAMPTZ` nullable ✓
  (INV-account-12 checkable), `UNIQUE INDEX
  ux_auth_identities_provider_identifier ON (provider_type,
  identifier_hash)` ✓ (the race-backing guard for Branch 1's
  duplicate-email case exists at the schema level), non-unique
  `ix_auth_identities_user_id` ✓ (per-user identity queries will be
  index-supported). `provider_type` CHECK constrains to
  `('email_password','google','phone_otp')`. **No soft-delete column** —
  spec's "hard-deleted" claim matches reality. `updated_at` trigger
  exists (irrelevant for DELETE).
- **`auth_tokens.purpose`**: CHECK `('email_verification',
  'password_reset')` — set-password's token reuses `email_verification`,
  fits with zero migration. Relevant later: adding a third purpose value
  would require a migration (see Stage 3 D7).
- **`refresh_tokens`**: `ix_refresh_tokens_user_id` and partial
  `ix_refresh_tokens_active ON (user_id) WHERE revoked_at IS NULL AND
  replaced_by_id IS NULL` exist — a revoke-all-by-user UPDATE would be
  index-supported.
- **`user_logs.action_type`**: plain TEXT, no CHECK constraint, no REVOKE
  yet — vocabulary/immutability deferred to task #08 per entity.go.
  `actionAccountLinking = "account_linking"` constant already defined in
  `google_oauth.go:44`.

## Requirement vs Gap

No migration appears needed for this feature's primary paths — every
schema primitive (nullable verified_at, unique provider+hash index,
user_id indexes, purpose values) already supports both endpoints. Gap is
purely code-level (Area 1's missing repo/service methods).

## Sniffing

- *Miscontext check*: spec assumes hard delete possible → confirmed;
  spec assumes same `/auth/verify-email` endpoint reusable → confirmed
  (`email_verification` purpose + existing token mechanics fit
  unchanged).
- *Edge case*: uniqueness is `(provider_type, identifier_hash)` globally
  (not per-user), so Branch 1's "email claimed by ANY user" conflict
  surfaces as a unique violation on insert — matching the registration
  pattern exactly.
- *Observation*: deleting a google identity row cascades nothing
  problematic (no FK points at auth_identities); refresh tokens are
  keyed to users, not identities — unlink does not itself kill sessions,
  which matches the spec (only Branch 2 password change revokes
  sessions).
