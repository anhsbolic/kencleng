# Task 01: Migrations 000006–000009 (schema-pre-settle)

> Back-reference : `.local-agents/works/account/03-login-session-management/2-plan/techplan.md` (Status: Approved) — sections 8 (DB Schema), 9 (Architecture), 5 (D1-C decision)
> Depends on    : nothing (first task in the chain)
> Model         : DeepSeek V4 Pro (precision DDL, rule-table-heavy work)
> Rules touched : R20 (additive, reversible); enables every later task

## Objective

Create the four additive golang-migrate pairs this slice owns. Three of the tables (`mfa_totp_secrets`, `mfa_backup_codes`, `user_roles`) are **schema-pre-settle only** (techplan Decision D1-C): long-term logic owners are account tasks #6/#8, who must not re-create these migrations (ownership note already applied to `docs/spec/1-account/tasks.md`, 2026-08-26).

## Files

| File | Change |
|---|---|
| `backend/migrations/000006_login_attempts.up.sql` | New |
| `backend/migrations/000006_login_attempts.down.sql` | New |
| `backend/migrations/000007_mfa_totp_secrets.up.sql` | New |
| `backend/migrations/000007_mfa_totp_secrets.down.sql` | New |
| `backend/migrations/000008_mfa_backup_codes.up.sql` | New |
| `backend/migrations/000008_mfa_backup_codes.down.sql` | New |
| `backend/migrations/000009_user_roles.up.sql` | New |
| `backend/migrations/000009_user_roles.down.sql` | New |

`TBD — verify`: confirm at build time that no migration numbered ≥ 000006 appeared on another branch since 2026-08-26 (golang-migrate requires unique increasing versions).

## DDL (verbatim from ERD / techplan §8 — do not redesign)

### 000006 — `login_attempts` (this slice owns)

```sql
CREATE TABLE login_attempts (
    id               UUID PRIMARY KEY,
    identifier_hash  TEXT NOT NULL,
    user_id          UUID REFERENCES users(id) ON DELETE SET NULL,
    stage            TEXT NOT NULL DEFAULT 'password'
                       CHECK (stage IN ('password', 'mfa')),
    success          BOOLEAN NOT NULL,
    attempted_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX ix_login_attempts_identifier_time
    ON login_attempts (identifier_hash, attempted_at DESC);
CREATE INDEX ix_login_attempts_user_stage_time
    ON login_attempts (user_id, stage, attempted_at DESC)
    WHERE user_id IS NOT NULL;
CREATE INDEX ix_login_attempts_attempted_at_brin
    ON login_attempts USING BRIN (attempted_at);
```

Index rationale (keep as SQL comments): first index = password-stage lockout hot query; second = MFA-stage lockout keyed by user_id (partial — most rows have NULL user_id); BRIN = cheap time-range audit queries on an append-only, naturally time-ordered table.

Down: `DROP TABLE IF EXISTS login_attempts;`

### 000007 — `mfa_totp_secrets` (schema-pre-settle; owner task #6)

```sql
CREATE TABLE mfa_totp_secrets (
    user_id           UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    secret_encrypted  BYTEA NOT NULL,
    enabled_at        TIMESTAMPTZ,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE TRIGGER trg_mfa_totp_secrets_updated_at
BEFORE UPDATE ON mfa_totp_secrets
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
```

Note: one row per user (`user_id` is the PK directly — upserted in place, never deleted). Down must drop trigger before table.

### 000008 — `mfa_backup_codes` (schema-pre-settle; owner task #6)

```sql
CREATE TABLE mfa_backup_codes (
    id          UUID PRIMARY KEY,
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    code_hash   TEXT NOT NULL,
    used_at     TIMESTAMPTZ,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX ix_mfa_backup_codes_user_id ON mfa_backup_codes (user_id);
CREATE INDEX ix_mfa_backup_codes_unused  ON mfa_backup_codes (user_id) WHERE used_at IS NULL;
```

### 000009 — `user_roles` (schema-pre-settle; owner task #8)

Verified against `docs/project/kencleng-erd.md:485-495`.

```sql
CREATE TABLE user_roles (
    id          UUID PRIMARY KEY,
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role        TEXT NOT NULL CHECK (role IN ('admin', 'kurator')),
    granted_by  UUID REFERENCES users(id) ON DELETE SET NULL,
    granted_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (user_id, role)
);
CREATE INDEX ix_user_roles_user_id ON user_roles (user_id);
CREATE INDEX ix_user_roles_role   ON user_roles (role);
```

Note: role-exclusivity invariants (INV-account-09/10) are application-layer, deliberately not expressible as plain constraints — do not add triggers to "improve" this.

## Hard constraints

1. **Additive only** — zero `ALTER` on existing tables (`users`, `auth_identities`, `auth_tokens`, `refresh_tokens`, `user_logs`). Empty-table semantics are load-bearing: they model "nobody has MFA enrolled / nobody has roles" today.
2. **Reversible** — downs are symmetric DROPs (trigger-before-table order where applicable). Not just written symmetrically; verified by round-trip below.
3. Plain SQL, no dynamic anything (R20 discipline extends to migration files).
4. Do not touch numbering of existing 000001–000005 files.

## Verification

```bash
make migrate-up      # applies 000006–000009 cleanly
make migrate-down    # rolls all four back
make migrate-up      # re-applies
```

All three steps exit 0 against the dev Postgres. Then `\d+ login_attempts` (or equivalent) confirms indexes incl. the two partial/composite ones and the BRIN.

## Out of scope

Any writes to the new tables (task #6/#8 territory); any Go code (later tasks); seeding (the Admin bootstrap seed script belongs to task #8).
