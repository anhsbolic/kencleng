-- login_attempts: append-only record of credential-verification
-- outcomes, backing the persistent lockout mechanism (Fitur 2C).
-- Two lockout stages with different keys:
--   password stage → keyed by identifier_hash (identity not yet known)
--   mfa stage      → keyed by user_id     (identity known via validated mfa_pending_token)
-- See docs/spec/1-account/features/03-login-session-management.md Assumption C.

CREATE TABLE login_attempts (
    id               UUID PRIMARY KEY,
    identifier_hash  TEXT NOT NULL,
    user_id          UUID REFERENCES users(id) ON DELETE SET NULL,
    stage            TEXT NOT NULL DEFAULT 'password'
                       CHECK (stage IN ('password', 'mfa')),
    success          BOOLEAN NOT NULL,
    attempted_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Password-stage lockout hot query: "how many failed attempts for this
-- identifier in the trailing window".
CREATE INDEX ix_login_attempts_identifier_time
    ON login_attempts (identifier_hash, attempted_at DESC);

-- MFA-stage lockout: identity already known via the validated
-- mfa_pending_token, so lookup is keyed by user_id + stage instead.
-- Partial: most rows are password-stage with NULL user_id.
CREATE INDEX ix_login_attempts_user_stage_time
    ON login_attempts (user_id, stage, attempted_at DESC)
    WHERE user_id IS NOT NULL;

-- Append-only, naturally time-ordered, high-insert table: BRIN on the
-- timestamp column keeps occasional "attempts in date range" audit
-- queries cheap without a second B-tree's write cost.
CREATE INDEX ix_login_attempts_attempted_at_brin
    ON login_attempts USING BRIN (attempted_at);
