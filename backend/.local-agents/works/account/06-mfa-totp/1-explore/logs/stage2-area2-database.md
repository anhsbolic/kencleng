# Stage 2 — Area 2: Database Layer — `mfa_totp_secrets` & `mfa_backup_codes`

> Files: `migrations/000007_mfa_totp_secrets.up.sql`, `migrations/000008_mfa_backup_codes.up.sql`, `internal/domain/account/repository.go`, `internal/domain/account/repository_db.go`

## Current State

Two migrations already exist (pre-created by task #3):

**`000007_mfa_totp_secrets.up.sql`:**
- `user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE` — one row per user, upserted on re-enrollment
- `secret_encrypted BYTEA NOT NULL` — AES-GCM ciphertext
- `enabled_at TIMESTAMPTZ` — nullable; NULL = pending/unconfirmed, non-NULL = active
- `created_at`, `updated_at` with `set_updated_at()` trigger

**`000008_mfa_backup_codes.up.sql`:**
- `id UUID PRIMARY KEY`
- `user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE`
- `code_hash TEXT NOT NULL` — SHA-256 or bcrypt hash of the plain code
- `used_at TIMESTAMPTZ` — nullable; NULL = unused, non-NULL = redeemed
- `ix_mfa_backup_codes_user_id` (full index)
- `ix_mfa_backup_codes_unused` (partial: `WHERE used_at IS NULL`)

**No repository methods exist yet** for these tables. `repository.go` has no MFA-related methods. `repository_db.go` has no MFA-related implementations.

`GetLoginUserView` (repository.go:217) already reads `MFAEnabled` from `mfa_totp_secrets.enabled_at IS NOT NULL` — so at least one read path exists, but the write paths are entirely missing.

## Requirement

Repository methods needed:
1. **UpsertMFATOTPSecret** — INSERT or UPDATE `mfa_totp_secrets` (upsert on `user_id`)
2. **EnableMFATOTP** — UPDATE `SET enabled_at = now() WHERE enabled_at IS NULL`
3. **DisableMFATOTP** — UPDATE `SET enabled_at = NULL`
4. **GetMFATOTPSecret** — SELECT `secret_encrypted, enabled_at` by `user_id`
5. **InsertMFABackupCodes** — Batch INSERT into `mfa_backup_codes`
6. **RedeemMFABackupCode** — UPDATE `SET used_at = now() WHERE used_at IS NULL`, returns bool

## Gap

All write-path repository methods are missing. Only `GetLoginUserView` reads from `mfa_totp_secrets` (and only the `enabled_at` boolean, not the encrypted secret). The verifier implementation needs methods 4 and 6; the service layer needs methods 1-3 and 5.

## Sniffing

- **Risk:** `RedeemMFABackupCode` is the most concurrency-sensitive method. Two concurrent login attempts using the same backup code must not both succeed. The `used_at IS NULL` guard in the UPDATE handles this under READ COMMITTED — the first tx wins, the second sees `used_at` already set and matches 0 rows. Same pattern as `RedeemToken` (INV-account-08). The partial index `ix_mfa_backup_codes_unused` supports this query efficiently.
- **Edge cases:** On re-enrollment (disable → enable cycle), old backup codes from the previous cycle remain in the table with `used_at` possibly NULL. The new enrollment inserts a fresh set of 10. The old codes are implicitly invalid because `enabled_at` was set to NULL during disable and back to non-NULL on re-enable — but the `enabled_at` check happens at the verifier level (INV-account-06), not at the DB query level. The `RedeemMFABackupCode` query must include the `enabled_at IS NOT NULL` check (join or subquery against `mfa_totp_secrets`), or the verifier must do it in two steps.
- **Miscontext:** The migration comment says "Owned by account task #6 (generation/enrollment); created here as schema-pre-settle per task #3's approved techplan D1-C." Task #6 must NOT create new migrations — only add repository/service/transport code.
- **Misleading signal:** `GetLoginUserView` already reads `MFAEnabled` — someone might think "MFA reading is done." But the verifier needs the actual `secret_encrypted` bytes, which `GetLoginUserView` does not return. A separate `GetMFATOTPSecret` method is needed.
- **Inconsistency:** None found. Schema matches the invariants and feature spec.
