CREATE TABLE persons (
    id uuid PRIMARY KEY,
    issuer text COLLATE "C" NOT NULL,
    subject text COLLATE "C" NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT persons_issuer_subject_unique UNIQUE (issuer, subject),
    CONSTRAINT persons_identity_nonempty CHECK (issuer <> '' AND subject <> '')
);

CREATE TABLE pending_logins (
    state_digest bytea PRIMARY KEY,
    browser_digest bytea NOT NULL,
    nonce text NOT NULL,
    pkce_verifier text NOT NULL,
    expires_at timestamptz NOT NULL,
    CONSTRAINT pending_login_digests CHECK (octet_length(state_digest) = 32 AND octet_length(browser_digest) = 32)
);
CREATE INDEX pending_logins_expires_at_idx ON pending_logins (expires_at);

CREATE TABLE sessions (
    token_digest bytea PRIMARY KEY,
    person_id uuid NOT NULL REFERENCES persons(id) ON DELETE RESTRICT,
    csrf_token text NOT NULL,
    created_at timestamptz NOT NULL,
    expires_at timestamptz NOT NULL,
    revoked_at timestamptz,
    CONSTRAINT session_digest_length CHECK (octet_length(token_digest) = 32),
    CONSTRAINT session_expiry_order CHECK (expires_at > created_at)
);
CREATE INDEX sessions_person_id_idx ON sessions (person_id);
