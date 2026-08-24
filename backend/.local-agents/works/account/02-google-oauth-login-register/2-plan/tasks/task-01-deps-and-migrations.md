# Task 01 — Dependencies + Migrations (refresh_tokens, user_logs)

> Ticket    : 02-google-oauth-login-register
> Sub-task  : 1 of 4
> Axis      : Dependency/sequence chain (primary) + component/layer alignment
> Status    : Ready — no Tier 0 prerequisite this time (`platform/crypto/` encrypt/HMAC functions already shipped in ticket 01)
> Back-ref  : `../techplan.md` — "Tech Plan: Google OAuth Login/Register" (originating contract techplan; cross-check high-level decisions there whenever needed)

---

## 1. Scope

Foundation for the OAuth task: one new Go dependency and two additive
migrations. Nothing else.

**In scope:**
- `backend/go.mod` — add `github.com/golang-jwt/jwt/v5` as a direct dependency
- `backend/migrations/000004_create_refresh_tokens.up.sql` + `.down.sql`
- `backend/migrations/000005_create_user_logs.up.sql` + `.down.sql`

**Out of scope (explicit):**
- Any Go code that uses these tables/entities — Task 03 owns entities,
  repository methods, and the service layer
- Rotation/reuse-detection indexes and constraints on refresh_tokens
  (INV-account-03, INV-account-04) — **task #3 adds those**; this migration
  ships only the minimal schema needed for token issuance (techplan §8 note)
- DB-level immutability constraint on user_logs
  (`REVOKE UPDATE, DELETE ON user_logs FROM kencleng_app`, INV-account-11)
  and any additional action_types/columns — **task #08 owns the full
  user_logs design**; this migration creates only the minimal columns needed
  for the link-intent audit entry (techplan §8 note, §14 Resolved item 2)
- `.env.example` — **no change needed** (verified against current file:
  FRONTEND_URL, GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET already present;
  GOOGLE_REDIRECT_URI already points at `:8090` for dev)

## 2. Dependencies

- **Hard deps:** none. This is the first task.
- **Blocks:** Task 02 (needs the jwt dependency importable),
  Task 03 (tables must exist for inserts and integration tests).

## 3. Files

| File | Change Type |
|---|---|
| `backend/go.mod` (+ `go.sum`) | Edit |
| `backend/migrations/000004_create_refresh_tokens.up.sql` | New |
| `backend/migrations/000004_create_refresh_tokens.down.sql` | New |
| `backend/migrations/000005_create_user_logs.up.sql` | New |
| `backend/migrations/000005_create_user_logs.down.sql` | New |

Current migrations end at `000003_create_auth_tokens` — numbering 000004 /
000005 is correct.

## 4. Implementation detail

### `backend/go.mod`

```
go get github.com/golang-jwt/jwt/v5
```

This is the only new dependency for the whole ticket (techplan §5 Decision
on OAuth client: golang-jwt/jwt/v5 + manual JWKS fetch chosen; coreos/go-oidc
rejected as overkill; raw std-lib rejected as error-prone). Tasks 02–04
consume it; they must not add further dependencies.

### Migration 000004 — refresh_tokens (minimal schema)

Up:

```sql
CREATE TABLE refresh_tokens (
    id              UUID PRIMARY KEY,
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    family_id       UUID NOT NULL,
    token_hash      TEXT NOT NULL,
    expires_at      TIMESTAMPTZ NOT NULL,
    revoked_at      TIMESTAMPTZ,
    replaced_by_id  UUID,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX ux_refresh_tokens_token_hash ON refresh_tokens (token_hash);
CREATE INDEX ix_refresh_tokens_user_id ON refresh_tokens (user_id);
CREATE INDEX ix_refresh_tokens_active ON refresh_tokens (user_id)
    WHERE revoked_at IS NULL AND replaced_by_id IS NULL;
```

Down:

```sql
DROP TABLE IF EXISTS refresh_tokens;
```

Deliberate design notes (do not "improve" these away):
- This is the schema **as specified in techplan §8** — verbatim.
- Task #2 (this ticket's consumer) only ever INSERTs into refresh_tokens; it
  does not query them. The indexes above are the minimal issuance-supporting
  set. Rotation/reuse queries (INV-account-03/04) arrive with task #3, which
  adds its own indexes via a later migration.
- Table starts empty — no backfill, no data migration.

### Migration 000005 — user_logs (minimal schema)

Up:

```sql
CREATE TABLE user_logs (
    id           UUID PRIMARY KEY,
    user_id      UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    action_type  TEXT NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ix_user_logs_user_id ON user_logs (user_id);
```

Down:

```sql
DROP TABLE IF EXISTS user_logs;
```

Notes:
- Minimal columns only: id, user_id, action_type, created_at (techplan §10).
- No REVOKE constraint here — INV-account-11 immutability is deferred to
  task #08 together with the full action_type vocabulary and trigger points
  across tasks #08/#06/#05. Do not add a partial action_type CHECK constraint
  now either; task #08 owns that vocabulary (frontend sign-off still pending,
  techplan §14 Active item 2).
- The canonical first literal to be written by this ticket's code is
  `action_type = 'account_linking'` (Fitur 9's "account linking baru") —
  written by Task 03's link branch, enabled by this table.

### Backward compatibility (techplan §6, applies to both migrations)

- Both migrations are purely additive: no existing column altered, no
  existing row touched.
- Down migrations drop the new tables cleanly.
- Existing email/password users are unaffected.

## 5. Rules covered (this task's slice)

From techplan §4:

- **R15 (concurrent duplicate registration)** — schema-level guarantee lives
  on `auth_identities (provider_type, identifier_hash)`, whose unique index
  already exists from ticket 01's migration 000002 — this task creates no
  index for it but must verify it is present when running the round-trip
  test below. The clean-error-handling half of R15 belongs to Task 03.
- Everything else this task produces is enabling infrastructure:
  - refresh_tokens ← consumed by `IssueTokens` (R7/R8 token issuance, Task 03)
  - user_logs ← consumed by link-intent audit entry (R11, Task 03)
  - jwt dependency ← consumed by JWKS verification (Task 02) and inline
    ES256 session verification (Task 04, R25)

Full rule-to-task mapping: see `manifest.md`.

## 6. Testing checklist (this task's slice)

- [ ] `make migrate-up` applies 000004 and 000005 cleanly on a fresh database.
- [ ] `make migrate-down` rolls both back cleanly; re-running up/down cycles
      repeatedly does not error (round-trip).
- [ ] After up: `\d refresh_tokens` shows all three indexes, including the
      partial index `ix_refresh_tokens_active` with
      `WHERE revoked_at IS NULL AND replaced_by_id IS NULL`.
- [ ] After up: `auth_identities` still carries the unique index on
      `(provider_type, identifier_hash)` from migration 000002 (R15
      precondition check).
- [ ] FK cascade sanity: inserting a row referencing a non-existent user_id
      fails; deleting a user removes their refresh_tokens/user_logs rows.
- [ ] `go build ./...` succeeds with the new direct dependency present;
      `go mod tidy` produces no diff afterwards.

## 7. Common mistakes to avoid

| Mistake | Fix |
|---|---|
| Adding rotation/reuse indexes "while we're here" | Deferred to task #3 per techplan §8 — scope discipline (root AGENTS.md §7). |
| Adding a REVOKE or CHECK constraint on user_logs | Deferred to task #08; vocabulary not yet signed off (§14 Active item 2). |
| Making token_hash nullable or non-unique | Unique index on token_hash is required (hash lookup target); keep NOT NULL. |
| Editing `.env.example` | Already correct per techplan §10 — leave untouched. |

## 8. Risk note (to fill in the PR)

- Assumptions made: ...
- Edge cases intentionally NOT handled (and why): ...
- Concurrency assumptions: ...
- What is not tested, and why: ...
