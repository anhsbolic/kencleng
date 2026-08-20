# Task 01 — Dependencies + Migrations

> Ticket    : 01-register-email-verification
> Sub-task  : 1 of 5
> Axis      : Dependency/sequence chain (primary) + component boundary
> Status    : Ready
> Back-ref  : `../2-plan/techplan.md` (originating contract techplan — cross-check high-level decisions there whenever needed)

---

## 1. Scope

This task establishes the foundation: the two missing Go dependencies
(`goqu`, `golang.org/x/time/rate`) and the three database migrations
that create the `users`, `auth_identities`, and `auth_tokens` tables
along with the shared `set_updated_at()` trigger function.

No business logic, no Go code beyond `go.mod`/`go.sum`. Purely
additive (greenfield — no existing tables altered, no data backfill).

**In scope:**
- `go get github.com/doug-martin/goqu/v9 golang.org/x/time/rate`
- `golang.org/x/crypto` transitions indirect → direct dep (bcrypt
  import in Task 02's `platform/secrets/`)
- 3 up + 3 down migration files, numbered sequentially per
  `golang-migrate` convention
- Unique indexes, partial index, FK constraints, CHECK constraints
  per ERD §1

**Out of scope:**
- Any Go package code (subsequent tasks)
- Migration CLI installation/config — migrations are run manually via
  `golang-migrate` CLI per tech stack doc, not embedded in app startup

## 2. Dependencies

- **Hard deps:** none (this is the chain root)
- **Soft deps:** none
- **Blocks:** Task 03 (needs schema to map entities against),
  Task 05 (needs `golang.org/x/time/rate` for middleware)

## 3. Files

| File | Change Type |
|---|---|
| `backend/go.mod` | Edit |
| `backend/go.sum` | Edit (auto) |
| `backend/migrations/000001_create_users.up.sql` | New |
| `backend/migrations/000001_create_users.down.sql` | New |
| `backend/migrations/000002_create_auth_identities.up.sql` | New |
| `backend/migrations/000002_create_auth_identities.down.sql` | New |
| `backend/migrations/000003_create_auth_tokens.up.sql` | New |
| `backend/migrations/000003_create_auth_tokens.down.sql` | New |

## 4. Implementation detail

### `backend/go.mod`

Run:
```bash
go get github.com/doug-martin/goqu/v9 golang.org/x/time/rate
```
`golang.org/x/crypto` will transition from indirect to direct when
Task 02 imports `golang.org/x/crypto/bcrypt` — ensure `go mod tidy`
resolves it cleanly.

### Migration 000001 — `set_updated_at()` + `users`

`migrations/000001_create_users.up.sql`:
```sql
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE users (
    id                  UUID PRIMARY KEY,
    name                TEXT NOT NULL,
    primary_email       BYTEA NOT NULL,
    primary_email_hash  TEXT NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX ux_users_primary_email_hash ON users (primary_email_hash);

CREATE TRIGGER trg_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
```

`migrations/000001_create_users.down.sql`:
```sql
DROP TRIGGER IF EXISTS trg_users_updated_at ON users;
DROP TABLE IF EXISTS users;
-- NOTE: set_updated_at() dropped only in the final down migration,
-- after all triggers using it are gone. See 000003 down.
```

### Migration 000002 — `auth_identities`

`migrations/000002_create_auth_identities.up.sql`:
```sql
CREATE TABLE auth_identities (
    id                 UUID PRIMARY KEY,
    user_id            UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    provider_type      TEXT NOT NULL CHECK (provider_type IN ('email_password', 'google', 'phone_otp')),
    identifier         BYTEA NOT NULL,
    identifier_hash    TEXT NOT NULL,
    credential_secret  TEXT,
    verified_at        TIMESTAMPTZ,
    created_at         TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at         TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX ux_auth_identities_provider_identifier
    ON auth_identities (provider_type, identifier_hash);

CREATE INDEX ix_auth_identities_user_id ON auth_identities (user_id);

CREATE TRIGGER trg_auth_identities_updated_at
BEFORE UPDATE ON auth_identities
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
```

`migrations/000002_create_auth_identities.down.sql`:
```sql
DROP TRIGGER IF EXISTS trg_auth_identities_updated_at ON auth_identities;
DROP TABLE IF EXISTS auth_identities;
```

### Migration 000003 — `auth_tokens`

`migrations/000003_create_auth_tokens.up.sql`:
```sql
CREATE TABLE auth_tokens (
    id          UUID PRIMARY KEY,
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    purpose     TEXT NOT NULL CHECK (purpose IN ('email_verification', 'password_reset')),
    token_hash  TEXT NOT NULL,
    expires_at  TIMESTAMPTZ NOT NULL,
    used_at     TIMESTAMPTZ,
    revoked_at  TIMESTAMPTZ,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX ux_auth_tokens_token_hash ON auth_tokens (token_hash);
CREATE INDEX ix_auth_tokens_user_purpose ON auth_tokens (user_id, purpose);

CREATE INDEX ix_auth_tokens_valid ON auth_tokens (user_id, purpose)
    WHERE used_at IS NULL AND revoked_at IS NULL;
```

`migrations/000003_create_auth_tokens.down.sql`:
```sql
DROP INDEX IF EXISTS ix_auth_tokens_valid;
DROP INDEX IF EXISTS ix_auth_tokens_user_purpose;
DROP INDEX IF EXISTS ux_auth_tokens_token_hash;
DROP TABLE IF EXISTS auth_tokens;
-- 000003 is the last table using set_updated_at(); drop the function now.
DROP FUNCTION IF EXISTS set_updated_at();
```

## 5. Notes & constraints

- `set_updated_at()` is created in 000001 up but dropped only in
  000003 down — the function is shared by `users` and `auth_identities`
  triggers, so it must outlive the tables that reference it on the up
  path, and be dropped only after the last referencing trigger is gone
  on the down path.
- `primary_email` and `identifier` are `BYTEA` (AES-GCM ciphertext);
  `*_hash` columns are `TEXT` (HMAC-SHA256 hex) — schema matches the
  PII encryption pattern in `AGENTS.md` golden rules. Encryption/HMAC
  **functions** themselves are Tier 0 fenced (`platform/crypto/`) and
  are NOT part of this task — see manifest prerequisite note.
- Partial index `ix_auth_tokens_valid` encodes the
  `used_at IS NULL AND revoked_at IS NULL` portion of the INV-account-08
  redemption guard at the index level; the `expires_at > now()` clause
  is non-indexable (time-relative) and stays in the `UPDATE ... WHERE`
  predicate in Task 04.

## 6. Verification

- `go mod tidy` clean; `go build ./...` succeeds with no unresolved imports.
- `make migrate-up` applies all three migrations in order without error.
- `make migrate-down` rolls back all three without error (down order:
  000003 → 000002 → 000001; `set_updated_at()` dropped by 000003).
- Schema inspection: `\d users`, `\d auth_identities`, `\d auth_tokens`
  show the expected columns, indexes, constraints, and triggers.

## 7. Risk note

- Assumptions made: `golang-migrate` version compatibility assumed
  unchanged from scaffold; `goqu` v9 is the current major (per techplan
  §10). UUIDs stored as native `UUID` type (pgx handles encoding).
- Edge cases intentionally NOT handled: none — migrations are
  declarative; edge cases belong to the code that uses these tables.
- Concurrency assumptions: n/a (no concurrent code in this task).
- What is not tested, and why: no Go unit tests for this task —
  migration correctness is exercised by integration tests in
  Task 03/04 (`//go:build integration`).
