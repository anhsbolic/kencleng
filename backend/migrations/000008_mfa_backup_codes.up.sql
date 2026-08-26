-- mfa_backup_codes: single-use recovery codes for MFA login
-- (INV-account-06). Owned by account task #6 (generation/enrollment);
-- created here as schema-pre-settle per task #3's approved techplan D1-C.

CREATE TABLE mfa_backup_codes (
    id          UUID PRIMARY KEY,
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    code_hash   TEXT NOT NULL,
    used_at     TIMESTAMPTZ,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ix_mfa_backup_codes_user_id ON mfa_backup_codes (user_id);

-- Partial index for the hot path: "unused codes remaining for this user".
CREATE INDEX ix_mfa_backup_codes_unused  ON mfa_backup_codes (user_id) WHERE used_at IS NULL;
