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
