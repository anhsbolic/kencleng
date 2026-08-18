# Kencleng — Entity Relationship Design (ERD)

> File: `docs/project/kencleng-erd.md`
> Status: Draft — Step 3 roadmap output. Derived from
> `kencleng-actors-entities.md`, `kencleng-phase0-detail.md` through
> `kencleng-phase3-detail.md`, plus decisions made during ERD
> discussion sessions (encryption-at-rest, `auth_tokens` unification,
> `login_attempts`, attachment tables, curation table split).
> Last updated: 2026-07-24 (rev 2 — `campaigns.category`/`location`/
> `beneficiary_description` added, org-per-user limit note, audit-log
> scope note for representative management)

## Context

This document is the target schema for `golang-migrate` migrations.
It's one level below the phase-detail docs (which describe business
rules/flow) and one level above actual `.sql` migration files (which
this doc should translate into almost verbatim).

Table names are plural, English, snake_case — conceptual entity names
used in the phase-detail docs (e.g. `Organisasi`, `Campaign`) map
1:1 to their plural English table name (`organizations`, `campaigns`)
unless noted otherwise.

---

## Design Conventions

These apply uniformly across all tables below, agreed during ERD
discussion — stated once here instead of repeated per table.

**Primary keys** — `UUID`, generated in the Go application layer
(UUID **v7**, time-orderable), not via Postgres `uuid-ossp` /
`gen_random_uuid()`. Rationale: v7 keeps B-tree index locality better
than random v4 (inserts land near the end of the index rather than
scattered), and app-generated IDs mean an entity's ID is known before
the INSERT round-trip (useful for cross-entity references built in
the same request, and for tests without a DB).

**Timestamps** — `TIMESTAMPTZ` everywhere (never bare `TIMESTAMP`).
`created_at NOT NULL DEFAULT now()`. `updated_at`, where present, is
maintained by a single shared trigger function rather than relying on
application code to remember to set it on every UPDATE:

```sql
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- attached per-table, e.g.:
CREATE TRIGGER trg_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
```

**Money** — `NUMERIC(19,2)`, never `FLOAT`/`DOUBLE PRECISION`.
Binary floating point cannot represent decimal fractions exactly,
which is unacceptable for financial data even when today's values
happen to be whole numbers. On the Go side, use
`github.com/shopspring/decimal` end-to-end (DB → app → API response)
so precision isn't silently lost re-marshaling into `float64`
somewhere in the middle. This is a new dependency added to
`kencleng-backend-tech-stack.md` — justified because this is
literally the money-handling path, a case where "add complexity only
with a genuine need" is clearly satisfied.

**Enums** — implemented as `TEXT NOT NULL CHECK (col IN (...))`, not
native Postgres `CREATE TYPE ... AS ENUM`. Trade-off: native enums are
marginally more storage-efficient and self-documenting via `\d`, but
altering them (removing a value, or reordering) is awkward — some
operations historically couldn't run inside the same transaction as
other DDL. A `CHECK` constraint is trivially altered by a normal
migration (`ALTER TABLE ... DROP CONSTRAINT ...; ALTER TABLE ... ADD
CONSTRAINT ...`), which matters more for a project still actively
iterating on business rules than the small storage/documentation win.

**PII encryption-at-rest** — applies to `users.primary_email`,
`auth_identities.identifier`, `organizations.npwp`,
`donations.guest_email` (per UU PDP compliance, decided in prior
session). Pattern, applied consistently:
- `{field}` — `BYTEA`, AES-GCM ciphertext (Go stdlib `crypto/aes`, no
  extra dependency)
- `{field}_hash` — `TEXT`, HMAC-SHA256 (deterministic, so it can be
  indexed/uniqued/looked-up — AES-GCM ciphertext itself can't be,
  since it's non-deterministic by design)

This is one layer deeper than the frontend `MaskedField` masking
already agreed in `kencleng-actors-entities.md`: masking protects
against unauthorized *viewing*, encryption-at-rest protects the data
even if the database itself is dumped/exfiltrated.

**Foreign key delete policy** — one rule, applied consistently instead
of deciding per-FK ad hoc:
- FK to a **required, non-nullable actor/owner** (e.g. `created_by`,
  `submitted_by`, `kurator_id`, `assigned_by`, `uploaded_by`) →
  `ON DELETE RESTRICT`. These are provenance-critical; the row that
  references them must never be able to end up pointing at nothing.
- FK to an **optional/nullable actor** (e.g. `reviewed_by`,
  `closed_by`, `granted_by`) → `ON DELETE SET NULL`. Losing the exact
  actor reference on a rare hard-delete is acceptable since the field
  was already optional.
- FK expressing **pure ownership** (child rows meaningless without
  the parent: attachments, report items, join/junction tables, auth
  artifacts like tokens/sessions/MFA secrets tied to a `user`) →
  `ON DELETE CASCADE`.
- FK from **business entities to other business entities**
  (`campaigns.organization_id`, `donations.campaign_id`,
  `disbursement_requests.campaign_id`) → `ON DELETE RESTRICT`. These
  should never be hard-deleted in practice (lifecycle is modeled via
  `status`, not row deletion) — `RESTRICT` makes that assumption
  explicit and fails loudly if ever violated, rather than silently
  cascading away financial history.

**Soft delete** — **not** applied generically. Most core entities
already carry a `status` enum that answers "is this row still
relevant" with clear, specific semantics — a blanket `deleted_at`
would just be a second, competing answer to the same question. Log
tables must never be deletable at all (see below), and token tables
already have more precise lifecycle columns (`used_at`, `revoked_at`,
`expires_at`) than a generic flag could offer. `deleted_at` gets added
**only** to a specific table if a concrete future need appears (e.g.
account deactivation for `users`), not as a default.

**Append-only log tables** — `user_logs`, `organization_logs`,
`campaign_logs`, `donation_logs`, `disbursement_request_logs`,
`fund_usage_report_logs` must be genuinely immutable, enforced at the
database privilege level rather than trusted to application code:

```sql
REVOKE UPDATE, DELETE ON
  user_logs, organization_logs, campaign_logs, donation_logs,
  disbursement_request_logs, fund_usage_report_logs
FROM kencleng_app;
```

(`kencleng_app` — the role the Go backend connects as; only `INSERT`
and `SELECT` remain grantable.)

**"One active row" invariants** — wherever a business rule says "only
one X can be pending/active at a time" (curation assignments,
disbursement requests), this is enforced with a **partial unique
index**, not just an application-level check-then-insert. A DB
constraint holds even under concurrent requests; an app-level check
has a race window between the check and the insert.

**Reconciliation strict-match** (`fund_usage_report_items.amount` sum
must equal `disbursed_amount`) is a cross-row aggregate invariant that
a plain `CHECK` constraint can't express (`CHECK` only sees one row at
a time). Recommended: enforce it with a Postgres **constraint
trigger** that runs on insert/update/delete of `fund_usage_report_items`
and re-validates the sum against the parent's disbursed amount within
the same transaction — this keeps the correctness guarantee at the DB
level (matching the project's "concurrency-safe, correctness-critical
backend code" learning goal) rather than trusting application code to
always re-check before commit. Sketch:

```sql
CREATE OR REPLACE FUNCTION check_fund_usage_reconciliation()
RETURNS TRIGGER AS $$
DECLARE
  v_report_id UUID;
  v_disbursed NUMERIC(19,2);
  v_total NUMERIC(19,2);
BEGIN
  v_report_id := COALESCE(NEW.fund_usage_report_id, OLD.fund_usage_report_id);

  SELECT dr.requested_amount INTO v_disbursed
  FROM fund_usage_reports fur
  JOIN disbursement_requests dr ON dr.id = fur.disbursement_request_id
  WHERE fur.id = v_report_id;

  SELECT COALESCE(SUM(amount), 0) INTO v_total
  FROM fund_usage_report_items
  WHERE fund_usage_report_id = v_report_id;

  IF v_total > v_disbursed THEN
    RAISE EXCEPTION 'fund_usage_report_items total (%) exceeds disbursed amount (%) for report %',
      v_total, v_disbursed, v_report_id;
  END IF;

  RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE CONSTRAINT TRIGGER trg_fund_usage_reconciliation
AFTER INSERT OR UPDATE OR DELETE ON fund_usage_report_items
DEFERRABLE INITIALLY DEFERRED
FOR EACH ROW EXECUTE FUNCTION check_fund_usage_reconciliation();
```

Note: the trigger only blocks *exceeding* the disbursed amount per-row
change; the final **exact-match** check (total == disbursed, not just
≤) still needs an explicit application-level check at submit time
(when `fund_usage_reports.status` moves to `pending_verification`),
since "exact match" is only meaningful once the Owner declares the
breakdown complete — a work-in-progress draft with a partial total is
expected to not match yet. The trigger's job is narrower: catch
*runaway* totals immediately, as a safety net under the app-level
check.

---

## Mermaid ER Diagram

```mermaid
erDiagram
    users ||--o{ user_roles : "has"
    users ||--o{ auth_identities : "has"
    users ||--o{ refresh_tokens : "has"
    users ||--o{ auth_tokens : "has"
    users ||--o| mfa_totp_secrets : "has"
    users ||--o{ mfa_backup_codes : "has"
    users ||--o{ login_attempts : "matches (optional)"
    users ||--o{ otps : "has"
    users ||--o{ user_terms_agreements : "agrees"
    terms_versions ||--o{ user_terms_agreements : "agreed via"
    users ||--o{ user_logs : "target of"

    users ||--o{ organization_representatives : "represents"
    organizations ||--o{ organization_representatives : "has"
    organizations ||--o{ organization_curation_assignments : "curated via"
    users ||--o{ organization_curation_assignments : "kurator"
    organizations ||--o{ organization_attachments : "has"
    organizations ||--o{ organization_logs : "logged"
    organizations ||--o{ campaigns : "owns"

    campaigns ||--o{ campaign_curation_assignments : "curated via"
    users ||--o{ campaign_curation_assignments : "kurator"
    campaigns ||--o{ campaign_attachments : "has"
    campaigns ||--o{ campaign_logs : "logged"
    campaigns ||--o{ campaign_events : "promoted via"
    events ||--o{ campaign_events : "promotes"
    users ||--o{ events : "created by"

    campaigns ||--o{ donations : "receives"
    events ||--o{ donations : "context (optional)"
    users ||--o{ donations : "donor (optional)"
    donations ||--o{ donation_logs : "logged"

    campaigns ||--o{ disbursement_requests : "requested for"
    users ||--o{ disbursement_requests : "requested by"
    disbursement_requests ||--o{ disbursement_request_logs : "logged"
    disbursement_requests ||--o{ fund_usage_reports : "reconciled by"

    users ||--o{ fund_usage_reports : "submitted by"
    fund_usage_reports ||--o{ fund_usage_report_items : "has"
    fund_usage_report_items ||--o{ fund_usage_report_item_attachments : "has"
    fund_usage_reports ||--o{ fund_usage_report_verification_assignments : "verified via"
    users ||--o{ fund_usage_report_verification_assignments : "kurator"
    fund_usage_reports ||--o{ fund_usage_report_logs : "logged"

    users ||--o{ notifications : "recipient (optional)"

    users {
        uuid id PK
        text name
        bytea primary_email
        text primary_email_hash
    }
    user_roles {
        uuid id PK
        uuid user_id FK
        text role
    }
    auth_identities {
        uuid id PK
        uuid user_id FK
        text provider_type
        bytea identifier
        text identifier_hash
        timestamptz verified_at
    }
    refresh_tokens {
        uuid id PK
        uuid user_id FK
        uuid family_id
        uuid replaced_by_id FK
    }
    auth_tokens {
        uuid id PK
        uuid user_id FK
        text purpose
        timestamptz revoked_at
    }
    mfa_totp_secrets {
        uuid user_id PK_FK
        timestamptz enabled_at
    }
    mfa_backup_codes {
        uuid id PK
        uuid user_id FK
        timestamptz used_at
    }
    login_attempts {
        uuid id PK
        text identifier_hash
        uuid user_id FK
        text stage
        boolean success
    }
    otps {
        uuid id PK
        uuid user_id FK
        text purpose
    }
    terms_versions {
        uuid id PK
        text version_number
    }
    user_terms_agreements {
        uuid user_id PK_FK
        uuid terms_version_id PK_FK
    }
    user_logs {
        uuid id PK
        uuid user_id FK
        text action_type
    }

    organizations {
        uuid id PK
        text name
        bytea npwp
        text npwp_hash
        text status
        boolean has_overdue_report
    }
    organization_representatives {
        uuid id PK
        uuid user_id FK
        uuid organization_id FK
        text level
    }
    organization_curation_assignments {
        uuid id PK
        uuid organization_id FK
        uuid kurator_id FK
        text decision
    }
    organization_attachments {
        uuid id PK
        uuid organization_id FK
        text type
    }
    organization_logs {
        uuid id PK
        uuid organization_id FK
    }

    campaigns {
        uuid id PK
        uuid organization_id FK
        text title
        numeric target_amount
        numeric max_amount
        numeric collected_amount
        text status
        text report_narrative
    }
    campaign_curation_assignments {
        uuid id PK
        uuid campaign_id FK
        uuid kurator_id FK
        text decision
    }
    campaign_attachments {
        uuid id PK
        uuid campaign_id FK
    }
    campaign_logs {
        uuid id PK
        uuid campaign_id FK
    }
    events {
        uuid id PK
        text name
        timestamptz event_datetime
    }
    campaign_events {
        uuid campaign_id PK_FK
        uuid event_id PK_FK
    }

    donations {
        uuid id PK
        uuid campaign_id FK
        uuid event_id FK
        uuid donor_user_id FK
        bytea guest_email
        text guest_email_hash
        numeric amount
        text status
        text status_token
        timestamptz claimed_at
    }
    donation_logs {
        uuid id PK
        uuid donation_id FK
    }

    disbursement_requests {
        uuid id PK
        uuid campaign_id FK
        uuid requested_by FK
        numeric requested_amount
        text status
        timestamptz disbursed_at
    }
    disbursement_request_logs {
        uuid id PK
        uuid disbursement_request_id FK
    }

    fund_usage_reports {
        uuid id PK
        uuid disbursement_request_id FK
        uuid submitted_by FK
        text status
    }
    fund_usage_report_items {
        uuid id PK
        uuid fund_usage_report_id FK
        text category
        numeric amount
    }
    fund_usage_report_item_attachments {
        uuid id PK
        uuid fund_usage_report_item_id FK
    }
    fund_usage_report_verification_assignments {
        uuid id PK
        uuid fund_usage_report_id FK
        uuid kurator_id FK
        text decision
    }
    fund_usage_report_logs {
        uuid id PK
        uuid fund_usage_report_id FK
    }

    notifications {
        uuid id PK
        uuid recipient_user_id FK
        bytea recipient_email
        text channel
        text type
        timestamptz expires_at
    }
```

---

## 1. Identity & Auth

### `users`

```sql
CREATE TABLE users (
    id                  UUID PRIMARY KEY,
    name                TEXT NOT NULL,
    primary_email       BYTEA NOT NULL,   -- AES-GCM ciphertext
    primary_email_hash  TEXT NOT NULL,    -- HMAC-SHA256, for lookup/uniqueness
    created_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX ux_users_primary_email_hash ON users (primary_email_hash);

CREATE TRIGGER trg_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
```

No `platform_role` column — role membership lives entirely in
`user_roles` (see below), which is the multi-role-ready design chosen
over a single enum column.

### `user_roles`

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

Business rules "Admin cannot also be Kurator/Representative" and "one
Kurator can't hold `kurator` role twice" are enforced at the
**application layer** on assignment (this table structurally *allows*
a user_id to have both rows — the invariant is behavioral, not
representable as a plain SQL constraint without a cross-table trigger,
and the assignment flow is already gated through Admin-only endpoints
per `kencleng-phase0-detail.md` Fitur 5).

### `auth_identities`

```sql
CREATE TABLE auth_identities (
    id                 UUID PRIMARY KEY,
    user_id            UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    provider_type      TEXT NOT NULL CHECK (provider_type IN ('email_password', 'google', 'phone_otp')),
    identifier         BYTEA NOT NULL,   -- AES-GCM ciphertext (email or phone)
    identifier_hash    TEXT NOT NULL,    -- HMAC-SHA256
    credential_secret  TEXT,             -- password hash; NULL for google/phone_otp
    verified_at        TIMESTAMPTZ,
    created_at         TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at         TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- uniqueness is per provider_type namespace, not global
CREATE UNIQUE INDEX ux_auth_identities_provider_identifier
    ON auth_identities (provider_type, identifier_hash);

CREATE INDEX ix_auth_identities_user_id ON auth_identities (user_id);

CREATE TRIGGER trg_auth_identities_updated_at
BEFORE UPDATE ON auth_identities
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
```

### `refresh_tokens`

```sql
CREATE TABLE refresh_tokens (
    id              UUID PRIMARY KEY,
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash      TEXT NOT NULL,
    family_id       UUID NOT NULL,
    issued_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    expires_at      TIMESTAMPTZ NOT NULL,
    revoked_at      TIMESTAMPTZ,
    replaced_by_id  UUID REFERENCES refresh_tokens(id) ON DELETE SET NULL
);

CREATE UNIQUE INDEX ux_refresh_tokens_token_hash ON refresh_tokens (token_hash);
CREATE INDEX ix_refresh_tokens_user_id   ON refresh_tokens (user_id);
CREATE INDEX ix_refresh_tokens_family_id ON refresh_tokens (family_id);

-- speeds up the rotate-on-use check: "is this token still active?"
CREATE INDEX ix_refresh_tokens_active ON refresh_tokens (user_id)
    WHERE revoked_at IS NULL AND replaced_by_id IS NULL;
```

### `auth_tokens`

Unified table for email-verification and password-reset tokens
(identical shape, different `purpose` + durations set by the app at
creation time). Replaces two separate tables originally sketched.

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

-- fast "does a still-valid token exist" check (e.g. before issuing a resend)
CREATE INDEX ix_auth_tokens_valid ON auth_tokens (user_id, purpose)
    WHERE used_at IS NULL AND revoked_at IS NULL;
```

### `mfa_totp_secrets`

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

One-to-one with `users` — `user_id` is the PK directly, no surrogate
`id` needed since there's never more than one active secret per user.

### `mfa_backup_codes`

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

### `login_attempts`

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

-- password-stage lockout: identity not yet known at check time,
-- keyed by identifier_hash (the "how many failed attempts for this
-- identifier in the last N minutes" hot query)
CREATE INDEX ix_login_attempts_identifier_time
    ON login_attempts (identifier_hash, attempted_at DESC);

-- MFA-stage lockout [ADDED — 2026-08-05, see
-- docs/spec/account/features/03-login-session-management.md
-- Assumption C]: identity is already known via the validated
-- mfa_pending_token by the time this check runs, so it's keyed by
-- user_id instead of identifier_hash
CREATE INDEX ix_login_attempts_user_stage_time
    ON login_attempts (user_id, stage, attempted_at DESC)
    WHERE user_id IS NOT NULL;

-- BRIN: high insert volume, naturally time-ordered, rarely queried by
-- anything other than identifier_hash/user_id (above) — a BRIN index
-- on the append-only timestamp column is far cheaper to maintain than
-- a second B-tree, useful for occasional "attempts in date range"
-- audit queries without bloating write cost
CREATE INDEX ix_login_attempts_attempted_at_brin
    ON login_attempts USING BRIN (attempted_at);
```

### `otps` *(placeholder — not used by any endpoint in v1)*

```sql
CREATE TABLE otps (
    id          UUID PRIMARY KEY,
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    code_hash   TEXT NOT NULL,
    purpose     TEXT NOT NULL CHECK (purpose IN ('phone_verification', 'phone_login')),
    expires_at  TIMESTAMPTZ NOT NULL,
    used_at     TIMESTAMPTZ,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ix_otps_user_purpose ON otps (user_id, purpose);
```

Physical table exists now so the migration history matches
`auth_identities.provider_type = 'phone_otp'` being modeled from day
one; the phone+OTP *flow* itself remains deferred to a later version.

### `terms_versions` / `user_terms_agreements`

```sql
CREATE TABLE terms_versions (
    id              UUID PRIMARY KEY,
    version_number  TEXT NOT NULL,
    content         TEXT NOT NULL,
    published_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX ux_terms_versions_version_number ON terms_versions (version_number);

CREATE TABLE user_terms_agreements (
    user_id           UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    terms_version_id  UUID NOT NULL REFERENCES terms_versions(id) ON DELETE RESTRICT,
    agreed_at         TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (user_id, terms_version_id)
);
```

### `user_logs`

```sql
CREATE TABLE user_logs (
    id             UUID PRIMARY KEY,
    user_id        UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    actor_user_id  UUID REFERENCES users(id) ON DELETE SET NULL,
    action_type    TEXT NOT NULL,
    metadata       JSONB,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ix_user_logs_user_id           ON user_logs (user_id, created_at DESC);
CREATE INDEX ix_user_logs_created_at_brin   ON user_logs USING BRIN (created_at);

REVOKE UPDATE, DELETE ON user_logs FROM kencleng_app;
```

Covers: role assign/revoke, MFA enable/disable, account linking
(Google), reveal of PII belonging to the user themself
(`primary_email`).

---

## 2. Organization & Representation

### `organizations`

```sql
CREATE TABLE organizations (
    id                  UUID PRIMARY KEY,
    name                TEXT NOT NULL,
    description         TEXT,
    contact             TEXT,
    npwp                BYTEA NOT NULL,   -- AES-GCM ciphertext
    npwp_hash           TEXT NOT NULL,    -- HMAC-SHA256, unique
    status              TEXT NOT NULL DEFAULT 'pending_verification'
                        CHECK (status IN ('pending_verification', 'verified', 'rejected')),
    has_overdue_report  BOOLEAN NOT NULL DEFAULT false,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX ux_organizations_npwp_hash ON organizations (npwp_hash);
CREATE INDEX ix_organizations_status ON organizations (status);

-- speeds up the scheduler/query that lists orgs currently blocked from new campaigns
CREATE INDEX ix_organizations_overdue ON organizations (id) WHERE has_overdue_report = true;

CREATE TRIGGER trg_organizations_updated_at
BEFORE UPDATE ON organizations
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
```

Legal documents (Akta, SK Kemenkumham, Izin PUB) are **not** columns
here — they live in `organization_attachments` below. `name`,
`description`, `contact` are the "operasional" field class (freely
editable); `npwp` + the attachment rows are the "legal/identitas"
class whose edit triggers re-curation (app-level logic on the
`UPDATE`/attachment-replace path, not a DB constraint).

**NPWP format validation [RESOLVED — NEW]**: validated at the
application layer as a **format check only** (regex against the
standard `XX.XXX.XXX.X-XXX.XXX`, 15-digit pattern) before encryption —
not a DB constraint, since it applies to the plaintext value prior to
`BYTEA` encryption. No external verification against DJP/Ditjen Pajak
records — that's out of scope for a sandbox project and would require
a government API integration. Actual legitimacy is still verified
manually by the assigned Kurator reviewing the uploaded legal
documents.

### `organization_representatives`

```sql
CREATE TABLE organization_representatives (
    id               UUID PRIMARY KEY,
    user_id          UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    organization_id  UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    level            TEXT NOT NULL CHECK (level IN ('owner', 'staff')),
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (user_id, organization_id)
);

CREATE INDEX ix_org_reps_organization_id ON organization_representatives (organization_id);
CREATE INDEX ix_org_reps_user_id         ON organization_representatives (user_id);

-- speeds up the "org must always have >=1 owner" check before a downgrade/removal
CREATE INDEX ix_org_reps_owners ON organization_representatives (organization_id)
    WHERE level = 'owner';
```

**Org-per-user limit [RESOLVED — NEW]**: a user may register at most
**5 organisasi**. Enforced at the application layer (`COUNT(*) FROM
organization_representatives WHERE user_id = ? AND level = 'owner'`
before allowing a new `organizations` insert), not a DB constraint —
already covered by the existing `ix_org_reps_user_id` index, no new
index needed. Chosen as a round, generous-but-bounded number to
prevent abuse (e.g. spam organisasi registration) without blocking
legitimate multi-organisasi involvement; not derived from concrete
usage data, since this is a sandbox project.

**Representative invite [RESOLVED — NEW]**: invite is direct-add, no
accept/consent step — an owner adds an existing, verified user
directly as `level = 'staff'` via this table, no additional
`invited_at`/`accepted_at` columns needed. See
`kencleng-roadmap-next-steps.md` and `kencleng-ux-page-map.md` for the
full rationale.

### `organization_curation_assignments`

```sql
CREATE TABLE organization_curation_assignments (
    id               UUID PRIMARY KEY,
    organization_id  UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    kurator_id       UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    assigned_by      UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    assigned_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    decision         TEXT NOT NULL DEFAULT 'pending'
                     CHECK (decision IN ('pending', 'approved', 'rejected')),
    decision_note    TEXT,
    decided_at       TIMESTAMPTZ
);

CREATE INDEX ix_org_curation_organization_id ON organization_curation_assignments (organization_id);
CREATE INDEX ix_org_curation_kurator_id      ON organization_curation_assignments (kurator_id);

-- enforces "only one active assignment per organization" at the DB level
CREATE UNIQUE INDEX ux_org_curation_one_pending
    ON organization_curation_assignments (organization_id) WHERE decision = 'pending';
```

### `organization_attachments`

```sql
CREATE TABLE organization_attachments (
    id               UUID PRIMARY KEY,
    organization_id  UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    type             TEXT NOT NULL CHECK (type IN ('akta_notaris', 'sk_kemenkumham', 'izin_pub')),
    original_name    TEXT NOT NULL,
    stored_name      TEXT NOT NULL,
    path             TEXT NOT NULL,
    size_bytes       BIGINT NOT NULL,
    content_type     TEXT NOT NULL,
    uploaded_by      UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ix_org_attachments_organization_id ON organization_attachments (organization_id, type);
```

Bucket visibility (private, per `kencleng-phase0-detail.md` Fitur 7)
is determined by the table itself, not a per-row `is_public` flag —
every row here is private-bucket by definition.

### `organization_logs`

```sql
CREATE TABLE organization_logs (
    id               UUID PRIMARY KEY,
    organization_id  UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    actor_user_id    UUID REFERENCES users(id) ON DELETE SET NULL,
    action_type      TEXT NOT NULL,
    metadata         JSONB,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ix_organization_logs_organization_id ON organization_logs (organization_id, created_at DESC);
CREATE INDEX ix_organization_logs_created_at_brin ON organization_logs USING BRIN (created_at);

REVOKE UPDATE, DELETE ON organization_logs FROM kencleng_app;
```

Covers: curation decisions, `has_overdue_report` set/clear, and
**representative management actions (invite/remove/promote/demote)
[RESOLVED — NEW]** — added to audit-log scope since these actions
determine who holds authority over the organisasi.

---

## 3. Campaign, Curation & Event

### `campaigns`

```sql
CREATE TABLE campaigns (
    id                        UUID PRIMARY KEY,
    organization_id           UUID NOT NULL REFERENCES organizations(id) ON DELETE RESTRICT,
    title                     TEXT NOT NULL,
    description               TEXT,
    category                  TEXT NOT NULL CHECK (category IN (
                                'bencana_alam', 'kesehatan', 'pendidikan', 'sosial', 'lainnya'
                              )),
    location                  TEXT,
    beneficiary_description   TEXT,
    target_amount             NUMERIC(19,2) NOT NULL CHECK (target_amount > 0),
    max_amount                NUMERIC(19,2) CHECK (max_amount IS NULL OR max_amount >= target_amount),
    deadline                  TIMESTAMPTZ NOT NULL,
    status                    TEXT NOT NULL DEFAULT 'draft'
                              CHECK (status IN (
                                'draft', 'pending_curation', 'approved', 'rejected',
                                'scheduled', 'published', 'unpublished', 'closed'
                              )),
    collected_amount          NUMERIC(19,2) NOT NULL DEFAULT 0 CHECK (collected_amount >= 0),
    publish_at                TIMESTAMPTZ,
    published_at              TIMESTAMPTZ,
    unpublish_reason          TEXT CHECK (unpublish_reason IN ('owner_manual', 'organization_re_verification')),
    closed_at                 TIMESTAMPTZ,
    closed_reason             TEXT CHECK (closed_reason IN ('max_amount_reached', 'deadline_reached', 'admin_force_closed')),
    closed_by                 UUID REFERENCES users(id) ON DELETE SET NULL,
    decision_note             TEXT,
    report_narrative          TEXT CHECK (report_narrative IS NULL OR char_length(report_narrative) <= 5000),
    created_by                UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    created_at                TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at                TIMESTAMPTZ NOT NULL DEFAULT now(),
    CHECK (closed_reason IS DISTINCT FROM 'admin_force_closed' OR closed_by IS NOT NULL)
);

CREATE INDEX ix_campaigns_organization_id ON campaigns (organization_id);
CREATE INDEX ix_campaigns_status          ON campaigns (status);

-- supports category filter on the public /campaign list page
CREATE INDEX ix_campaigns_category ON campaigns (category) WHERE status = 'published';

-- tailored to the most common public read: browse currently-live campaigns
CREATE INDEX ix_campaigns_published ON campaigns (published_at DESC) WHERE status = 'published';

-- scheduler job: campaigns due to auto-publish
CREATE INDEX ix_campaigns_scheduled_due ON campaigns (publish_at) WHERE status = 'scheduled';

-- scheduler job: campaigns due to auto-close by deadline
CREATE INDEX ix_campaigns_deadline_due ON campaigns (deadline) WHERE status = 'published';

CREATE TRIGGER trg_campaigns_updated_at
BEFORE UPDATE ON campaigns
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
```

`decision_note` is reused for both manual-unpublish reason and
force-close reason (never set for both at once, per the state
machine) — full history of *why*, across every occurrence, still
lives in `campaign_logs`, so nothing is lost by not keeping a separate
column per action type.

**`category`, `location`, `beneficiary_description` [RESOLVED — NEW]**:
- `category` — required enum, used for filtering on the public
  `/campaign` list page.
- `location` — optional free-text (city/province), not a structured
  geo/lat-long field — sufficient for coarse filtering/display without
  a mapping dependency.
- `beneficiary_description` — optional free-text, resolved as a simple
  field rather than a dedicated `Beneficiary` entity (see
  `kencleng-roadmap-next-steps.md`, beneficiary entity discussion): no
  current feature needs to track a beneficiary as a reusable entity
  across campaigns or verify it independently of the organisasi.

### `campaign_curation_assignments`

Same shape as `organization_curation_assignments`:

```sql
CREATE TABLE campaign_curation_assignments (
    id             UUID PRIMARY KEY,
    campaign_id    UUID NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,
    kurator_id     UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    assigned_by    UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    assigned_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    decision       TEXT NOT NULL DEFAULT 'pending'
                   CHECK (decision IN ('pending', 'approved', 'rejected')),
    decision_note  TEXT,
    decided_at     TIMESTAMPTZ
);

CREATE INDEX ix_campaign_curation_campaign_id ON campaign_curation_assignments (campaign_id);
CREATE INDEX ix_campaign_curation_kurator_id  ON campaign_curation_assignments (kurator_id);

CREATE UNIQUE INDEX ux_campaign_curation_one_pending
    ON campaign_curation_assignments (campaign_id) WHERE decision = 'pending';
```

### `campaign_attachments`

```sql
CREATE TABLE campaign_attachments (
    id             UUID PRIMARY KEY,
    campaign_id    UUID NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,
    original_name  TEXT NOT NULL,
    stored_name    TEXT NOT NULL,
    path           TEXT NOT NULL,
    size_bytes     BIGINT NOT NULL,
    content_type   TEXT NOT NULL,
    uploaded_by    UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ix_campaign_attachments_campaign_id ON campaign_attachments (campaign_id);
```

Public bucket by definition (campaign media), unlike
`organization_attachments` — no `type` enum needed yet since v1 only
has one media purpose; add one later if e.g. cover-vs-gallery
distinction becomes a real need.

### `campaign_logs`

```sql
CREATE TABLE campaign_logs (
    id             UUID PRIMARY KEY,
    campaign_id    UUID NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,
    actor_user_id  UUID REFERENCES users(id) ON DELETE SET NULL,
    action_type    TEXT NOT NULL,
    metadata       JSONB,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ix_campaign_logs_campaign_id  ON campaign_logs (campaign_id, created_at DESC);
CREATE INDEX ix_campaign_logs_created_brin ON campaign_logs USING BRIN (created_at);

REVOKE UPDATE, DELETE ON campaign_logs FROM kencleng_app;
```

Covers: curation decisions, publish/unpublish (manual + auto),
force-close.

### `events` / `campaign_events`

```sql
CREATE TABLE events (
    id              UUID PRIMARY KEY,
    name            TEXT NOT NULL,
    event_datetime  TIMESTAMPTZ NOT NULL,
    location        TEXT,
    description     TEXT,
    created_by      UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ix_events_datetime ON events (event_datetime);

CREATE TABLE campaign_events (
    campaign_id  UUID NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,
    event_id     UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    PRIMARY KEY (campaign_id, event_id)
);

-- PK covers campaign_id-first lookups; this covers the reverse direction
CREATE INDEX ix_campaign_events_event_id ON campaign_events (event_id);
```

---

## 4. Donation

### `donations`

```sql
CREATE TABLE donations (
    id                 UUID PRIMARY KEY,
    campaign_id        UUID NOT NULL REFERENCES campaigns(id) ON DELETE RESTRICT,
    event_id           UUID REFERENCES events(id) ON DELETE SET NULL,
    donor_user_id      UUID REFERENCES users(id) ON DELETE SET NULL,
    guest_name         TEXT,
    guest_email        BYTEA,   -- AES-GCM ciphertext, nullable
    guest_email_hash   TEXT,    -- HMAC-SHA256, nullable
    amount             NUMERIC(19,2) NOT NULL CHECK (amount >= 5000),
    payment_method     TEXT NOT NULL CHECK (payment_method IN
                       ('transfer', 'debit', 'gopay', 'shopeepay', 'ovo', 'qris')),
    is_anonymous       BOOLEAN NOT NULL DEFAULT false,
    status             TEXT NOT NULL DEFAULT 'pending'
                       CHECK (status IN ('pending', 'success', 'failed')),
    status_token       TEXT NOT NULL,
    claimed_at         TIMESTAMPTZ,
    created_at         TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX ux_donations_status_token ON donations (status_token);

-- covers both the progress-page read and the "list donations for campaign X, filter by status" query
CREATE INDEX ix_donations_campaign_status ON donations (campaign_id, status);

CREATE INDEX ix_donations_donor_user_id ON donations (donor_user_id) WHERE donor_user_id IS NOT NULL;

-- the guest-claim flow query: WHERE guest_email_hash = ? AND donor_user_id IS NULL
CREATE INDEX ix_donations_guest_claim ON donations (guest_email_hash)
    WHERE donor_user_id IS NULL AND guest_email_hash IS NOT NULL;
```

`campaign_id` FK is `RESTRICT` (business entity to business entity,
per the FK policy above) — a campaign with donations attached must
never be hard-deletable. The atomic-increment pattern for
`collected_amount` (`UPDATE campaigns SET collected_amount =
collected_amount + :amount WHERE id = :id AND status = 'published'
RETURNING collected_amount`) already documented in
`kencleng-phase2-detail.md` needs no additional index beyond the
`campaigns` PK — Postgres row-level locking on that single `UPDATE`
already serializes concurrent donations correctly.

### `donation_logs`

```sql
CREATE TABLE donation_logs (
    id             UUID PRIMARY KEY,
    donation_id    UUID NOT NULL REFERENCES donations(id) ON DELETE CASCADE,
    actor_user_id  UUID REFERENCES users(id) ON DELETE SET NULL,
    action_type    TEXT NOT NULL,
    metadata       JSONB,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ix_donation_logs_donation_id ON donation_logs (donation_id, created_at DESC);

REVOKE UPDATE, DELETE ON donation_logs FROM kencleng_app;
```

Covers: reveal of `guest_email` (PII) by Admin/Kurator.

---

## 5. Post-Campaign (Disbursement & Fund Usage)

### `disbursement_requests`

```sql
CREATE TABLE disbursement_requests (
    id                UUID PRIMARY KEY,
    campaign_id       UUID NOT NULL REFERENCES campaigns(id) ON DELETE RESTRICT,
    requested_by      UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    requested_amount  NUMERIC(19,2) NOT NULL CHECK (requested_amount > 0),
    status            TEXT NOT NULL DEFAULT 'pending'
                      CHECK (status IN ('pending', 'approved', 'rejected', 'disbursed')),
    reviewed_by       UUID REFERENCES users(id) ON DELETE SET NULL,
    decision_note     TEXT,
    decided_at        TIMESTAMPTZ,
    disbursed_at      TIMESTAMPTZ,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ix_disbursement_requests_campaign_id ON disbursement_requests (campaign_id);

-- enforces "only one active request per campaign" at the DB level
CREATE UNIQUE INDEX ux_disbursement_requests_one_active
    ON disbursement_requests (campaign_id) WHERE status IN ('pending', 'approved', 'disbursed');
```

### `disbursement_request_logs`

```sql
CREATE TABLE disbursement_request_logs (
    id                        UUID PRIMARY KEY,
    disbursement_request_id  UUID NOT NULL REFERENCES disbursement_requests(id) ON DELETE CASCADE,
    actor_user_id             UUID REFERENCES users(id) ON DELETE SET NULL,
    action_type               TEXT NOT NULL,
    metadata                  JSONB,
    created_at                TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ix_disbursement_request_logs_request_id
    ON disbursement_request_logs (disbursement_request_id, created_at DESC);

REVOKE UPDATE, DELETE ON disbursement_request_logs FROM kencleng_app;
```

### `fund_usage_reports`

```sql
CREATE TABLE fund_usage_reports (
    id                        UUID PRIMARY KEY,
    disbursement_request_id  UUID NOT NULL REFERENCES disbursement_requests(id) ON DELETE RESTRICT,
    submitted_by              UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    status                    TEXT NOT NULL DEFAULT 'pending_verification'
                              CHECK (status IN ('pending_verification', 'verified', 'rejected')),
    submitted_at               TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ix_fund_usage_reports_disbursement_request_id
    ON fund_usage_reports (disbursement_request_id);
CREATE INDEX ix_fund_usage_reports_status ON fund_usage_reports (status);
```

Linked to `disbursement_request_id` (not `campaign_id` directly, per
your confirmed change) — this is what makes the strict-match
reconciliation and the 30-day deadline unambiguous: both are computed
directly off `disbursement_requests.requested_amount` /
`disbursed_at` via a single join, rather than an implicit "find the
relevant disbursement for this campaign" lookup.

### `fund_usage_report_items`

```sql
CREATE TABLE fund_usage_report_items (
    id                     UUID PRIMARY KEY,
    fund_usage_report_id  UUID NOT NULL REFERENCES fund_usage_reports(id) ON DELETE CASCADE,
    category               TEXT NOT NULL,
    amount                 NUMERIC(19,2) NOT NULL CHECK (amount > 0),
    description            TEXT,
    created_at             TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ix_fund_usage_report_items_report_id ON fund_usage_report_items (fund_usage_report_id);
```

Covered by the reconciliation constraint trigger described in Design
Conventions above.

### `fund_usage_report_item_attachments`

```sql
CREATE TABLE fund_usage_report_item_attachments (
    id                         UUID PRIMARY KEY,
    fund_usage_report_item_id  UUID NOT NULL REFERENCES fund_usage_report_items(id) ON DELETE CASCADE,
    original_name              TEXT NOT NULL,
    stored_name                 TEXT NOT NULL,
    path                        TEXT NOT NULL,
    size_bytes                  BIGINT NOT NULL,
    content_type                TEXT NOT NULL,
    uploaded_by                 UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    created_at                  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ix_fund_usage_item_attachments_item_id
    ON fund_usage_report_item_attachments (fund_usage_report_item_id);
```

One-to-many (an item can have more than one supporting
nota/foto/dokumen) — private bucket by definition, same as
`organization_attachments`.

### `fund_usage_report_verification_assignments`

Same shape as the other two curation-assignment tables:

```sql
CREATE TABLE fund_usage_report_verification_assignments (
    id                     UUID PRIMARY KEY,
    fund_usage_report_id  UUID NOT NULL REFERENCES fund_usage_reports(id) ON DELETE CASCADE,
    kurator_id             UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    assigned_by            UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    assigned_at            TIMESTAMPTZ NOT NULL DEFAULT now(),
    decision               TEXT NOT NULL DEFAULT 'pending'
                           CHECK (decision IN ('pending', 'approved', 'rejected')),
    decision_note          TEXT,
    decided_at             TIMESTAMPTZ
);

CREATE INDEX ix_fund_usage_verif_report_id  ON fund_usage_report_verification_assignments (fund_usage_report_id);
CREATE INDEX ix_fund_usage_verif_kurator_id ON fund_usage_report_verification_assignments (kurator_id);

CREATE UNIQUE INDEX ux_fund_usage_verif_one_pending
    ON fund_usage_report_verification_assignments (fund_usage_report_id) WHERE decision = 'pending';
```

### `fund_usage_report_logs`

```sql
CREATE TABLE fund_usage_report_logs (
    id                     UUID PRIMARY KEY,
    fund_usage_report_id  UUID NOT NULL REFERENCES fund_usage_reports(id) ON DELETE CASCADE,
    actor_user_id           UUID REFERENCES users(id) ON DELETE SET NULL,
    action_type             TEXT NOT NULL,
    metadata                JSONB,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ix_fund_usage_report_logs_report_id
    ON fund_usage_report_logs (fund_usage_report_id, created_at DESC);

REVOKE UPDATE, DELETE ON fund_usage_report_logs FROM kencleng_app;
```

Covers: verification decisions.

---

## 6. Cross-Cutting

### `notifications`

```sql
CREATE TABLE notifications (
    id                     UUID PRIMARY KEY,
    recipient_user_id      UUID REFERENCES users(id) ON DELETE CASCADE,
    recipient_email        BYTEA,   -- AES-GCM ciphertext, nullable (guest recipient)
    recipient_email_hash   TEXT,    -- HMAC-SHA256, nullable
    channel                TEXT NOT NULL CHECK (channel IN ('in_app', 'email')),
    type                   TEXT NOT NULL,
    payload                JSONB,
    read_at                TIMESTAMPTZ,
    expires_at             TIMESTAMPTZ NOT NULL,
    created_at             TIMESTAMPTZ NOT NULL DEFAULT now(),
    CHECK (recipient_user_id IS NOT NULL OR recipient_email IS NOT NULL)
);

-- notification-center unread badge + list query, per recipient
CREATE INDEX ix_notifications_recipient_unread
    ON notifications (recipient_user_id, read_at) WHERE recipient_user_id IS NOT NULL;

-- both the soft-hide filter (WHERE expires_at > now()) and the weekly
-- hard-delete worker (WHERE expires_at < now()) use this same index —
-- a plain B-tree, not BRIN, since this one IS on the hot read path
-- (evaluated on every notification-center page load), unlike the log
-- tables above which are write-heavy/rarely-range-scanned
CREATE INDEX ix_notifications_expires_at ON notifications (expires_at);
```

---

## Performance & Optimization Summary

Recap of the non-obvious indexing choices made throughout, gathered in
one place:

| Pattern | Where used | Why |
|---|---|---|
| **Partial unique index for "one active row"** | `organization_curation_assignments`, `campaign_curation_assignments`, `fund_usage_report_verification_assignments`, `disbursement_requests` | DB-enforced invariant, no race window between check and insert — stronger than an app-level check-then-insert |
| **Partial index scoped to hot status** | `campaigns` (published/scheduled/deadline-due), `donations` (guest-claim eligible rows), `organizations` (overdue), `refresh_tokens` (active) | Index only covers the subset of rows actually queried in the hot path, keeping it small and fast to maintain relative to a full-table index |
| **BRIN on append-only timestamp columns** | `login_attempts`, all `*_logs` tables | These tables grow indefinitely and are written far more than range-queried; BRIN is orders of magnitude cheaper to maintain than a B-tree for naturally time-ordered data, at the cost of slightly coarser range-scan precision — an acceptable trade for tables that are mostly point-looked-up by FK (already covered by a separate B-tree) rather than scanned by time range |
| **Composite index matching exact query shape** | `donations (campaign_id, status)`, `auth_identities (provider_type, identifier_hash)`, `login_attempts (identifier_hash, attempted_at DESC)` | Matches the actual `WHERE`/`ORDER BY` combination used by the flows in the phase-detail docs, rather than single-column indexes that would force a bitmap-and or extra sort |
| **DB-level immutability via `REVOKE`** | all `*_logs` tables | Stronger guarantee than "the app never issues an UPDATE/DELETE" — holds even against a compromised or buggy app instance, matching the append-only requirement from `kencleng-phase0-detail.md` Fitur 9 |
| **Constraint trigger for cross-row invariant** | `fund_usage_report_items` reconciliation | The one place a plain `CHECK` genuinely can't express the rule (it needs a `SUM()` across sibling rows) — deferred trigger keeps it correct even under multi-statement edits within one transaction |
| **Atomic conditional `UPDATE ... RETURNING`** (already established in `kencleng-phase2-detail.md`, restated here for completeness) | `campaigns.collected_amount` increment | Postgres row-level lock on the single `UPDATE` serializes concurrent donations without app-level locking — no new index needed beyond the PK |

---

## Open Items Carried Into Migration Writing

- Exact `kencleng_app` DB role/privilege setup (the `REVOKE` statements
  above assume this role exists — needs to be created as part of the
  first migration or deploy script)
- ~~`login_attempts` lockout threshold & window~~ → **resolved: 5
  failed attempts / 15 minute window** — see
  `kencleng-phase0-detail.md` Fitur 2C **[RESOLVED]**
- ~~Encryption key management (`ENCRYPTION_KEY`/`HMAC_KEY` env vars,
  rotation strategy)~~ → **resolved: 2 separate keys, env vars, no
  rotation in v1** — see `kencleng-backend-tech-stack.md`, "Encryption
  Key Management" **[RESOLVED]**
- ~~Whether `notifications.recipient_email_hash` is ever actually
  queried by anything~~ → **resolved: keep it**, for future-proofing
  (symmetry with the encryption pattern, and to allow future
  admin-search-by-email without a schema migration) even though
  unused by any endpoint in v1 **[RESOLVED — NEW]**
- **New `type` values needed for `notifications` [NEW]**: extending
  the enum to cover curation-lifecycle events (Admin notified on new
  queue item, Kurator notified on assignment, Owner notified on
  fund-usage-report/disbursement decisions) — see
  `kencleng-roadmap-next-steps.md` notification-mechanism discussion.
  No schema change needed (`type` is already a plain `TEXT`, not a
  `CHECK`-constrained enum), just new string values used by the
  application.