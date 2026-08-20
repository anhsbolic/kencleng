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
