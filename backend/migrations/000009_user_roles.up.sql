-- user_roles: platform role assignments ('admin', 'kurator'). Owned by
-- account task #8 (role assignment API); created here as schema-pre-settle
-- per task #3's approved techplan D1-C.
--
-- Role-exclusivity rules (INV-account-09 Admin ⊥ Kurator, INV-account-10
-- Admin ⊥ Representative) are enforced at the application layer on
-- assignment — this table structurally allows one user_id to hold both
-- rows by design (see kencleng-erd.md).

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
