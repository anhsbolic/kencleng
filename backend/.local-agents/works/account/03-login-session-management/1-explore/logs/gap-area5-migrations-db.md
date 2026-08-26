# Gap Analysis — Area 5: migrations + DB layer

> Files: `backend/migrations/000001-000005`, `docs/project/kencleng-erd.md`,
> `docs/spec/1-account/tasks.md` (ownership/sequencing context)

## Current state

- Migrations `000001`–`000005`: `users`, `auth_identities`, `auth_tokens`,
  `refresh_tokens`, `user_logs`. Plain-SQL golang-migrate pairs; next number
  **000006**.
- **`refresh_tokens` is rotation-ready already** (`000004.up.sql`):
  `family_id UUID NOT NULL`, unique `token_hash`, `revoked_at`,
  `replaced_by_id`, and partial index
  `ix_refresh_tokens_active (user_id) WHERE revoked_at IS NULL AND
  replaced_by_id IS NULL` — exactly matching INV-03/04's state machine.
- `auth_identities`: `credential_secret TEXT` (bcrypt), `verified_at`,
  `identifier BYTEA` + `identifier_hash TEXT`, unique
  `(provider_type, identifier_hash)`.
- **Missing tables:** `login_attempts`, `mfa_totp_secrets`,
  `mfa_backup_codes`, `user_roles`.
- **ERD already contains the updated `login_attempts` DDL**
  (`kencleng-erd.md:615-650`) incl. `stage CHECK ('password','mfa')`, both
  lookup indexes, BRIN — annotated `[ADDED — 2026-08-05, see …/
  03-login-session-management.md Assumption C]`. ERD also defines
  `mfa_totp_secrets` (:584), `mfa_backup_codes` (:603), `user_roles` (:485).
- **tasks.md ownership:** task #6 (MFA) owns mfa tables (Group B); task #8
  owns `user_roles` (Group C); both suggested to run in parallel with S1
  "once their independent tables' migrations are settled." Task #3 (this)
  sits mid-S1. Status tracker says all tasks "not started."

## Requirement

New migration creating `login_attempts` per spec/ERD DDL; read access to
`mfa_totp_secrets.enabled_at` (branch decision + `mfa_enabled`),
backup-code verification (`/auth/login/mfa`), and `user_roles`
(`roles[]` in every LoginResponse).

## Gap

1. `login_attempts` migration missing — direct add, no conflicts.
2. `mfa_*` + `user_roles` migrations missing while this feature's behavior
   reads them; owned by unstarted tasks #6/#8. No doc resolves how #3
   proceeds → Stage 3 D1 (largest open sequencing gap).
3. Status tracker stale: claims #1–#8 all "not started," but #1/#2 are
   demonstrably shipped (handlers, tests, full `.local-agents/works/account/
   01-*`/`02-*` explore→build→review cycles).

## Sniffing findings

- **Risk:** if this task creates the `mfa_*` migrations itself, it collides
  with #6's Group-B table ownership and breaks parallelization premise; if it
  doesn't, two of four endpoints + one response field are unimplementable.
  Resolved by human decision in Stage 3 (Option C: schema-pre-settle).
- **Misleading signal:** ERD shows complete DDL for five tables absent from
  `migrations/` — reading the ERD alone suggests schema is settled *and
  present*; it is neither present nor wholly this task's to create.
- **Inconsistency:** feature spec (:229-232) says the ERD update "needs to be
  applied as a follow-up" — it has already been applied; spec prose stale,
  not wrong.
- **Inconsistency #2:** tasks.md status tracker staleness (flagged only;
  not silently fixed).
- **Edge case worth keeping:** `login_attempts.user_id ON DELETE SET NULL` —
  lockout history keyed by `identifier_hash` survives user deletion; an
  anti-enumeration property, not a bug.
