-- mfa_totp_secrets: one row per user (user_id IS the PK — upserted in
-- place on re-enrollment, never deleted). Owned by account task #6
-- (enrollment logic); created here as schema-pre-settle per task #3's
-- approved techplan D1-C. secret_encrypted follows the established PII
-- encryption pattern (AES-GCM ciphertext via platform/crypto).

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
