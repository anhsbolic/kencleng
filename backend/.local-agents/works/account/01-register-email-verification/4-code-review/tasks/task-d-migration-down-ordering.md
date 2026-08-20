# Task D — Down-Migration Ordering Fix

> Ticket    : 01-register-email-verification
> Sub-task  : D of F (post-review remediation)
> Finding   : S5 (DROP FUNCTION fails while triggers depend on it)
> Blocking  : yes
> Back-ref  : `../report.md` §1 (S5); `postgresql/migrations-safety.md`

---

## 1. Scope

`golang-migrate` runs down migrations in reverse order: 000003 →
000002 → 000001. `000003_create_auth_tokens.down.sql` ends with
`DROP FUNCTION IF EXISTS set_updated_at();`. When that runs, the
triggers `trg_users_updated_at` (users) and
`trg_auth_identities_updated_at` (auth_identities) still exist and
reference `set_updated_at()`, so Postgres rejects the `DROP FUNCTION`
(no `CASCADE`) with "cannot drop function because other objects depend
on it." The migration half-applies and the `schema_migrations` state
becomes "dirty."

The comment in `000001_create_users.down.sql:3-4` says the function is
"dropped only in the final down migration" but points to `000003 down`
— which runs *first* (the author confused file-number order with
reverse-run order).

**In scope:**
- Move `DROP FUNCTION IF EXISTS set_updated_at();` to
  `000001_create_users.down.sql` (the last down to run, after both
  triggers are dropped).
- Remove it from `000003_create_auth_tokens.down.sql`.
- Fix the comments.
- Verify by running `make migrate-down && make migrate-up` from a clean
  applied state.

**Out of scope:**
- Any up migration (they're correct and already applied in dev).
- The `set_updated_at()` function body (correct).
- `auth_tokens` not having an `updated_at` column (intentional —
  tokens are immutable once created except `used_at`/`revoked_at`, which
  are set directly; no trigger needed).

## 2. Dependencies

- **Hard deps:** none.
- **Soft deps:** none.
- **Blocks:** none.

## 3. Files

| File | Change Type | Why |
|---|---|---|
| `migrations/000003_create_auth_tokens.down.sql` | Edit | remove the `DROP FUNCTION` line |
| `migrations/000001_create_users.down.sql` | Edit | add the `DROP FUNCTION` after the trigger+table drop |

## 4. Implementation detail

### `migrations/000001_create_users.down.sql`

```sql
DROP TRIGGER IF EXISTS trg_users_updated_at ON users;
DROP TABLE IF EXISTS users;
-- set_updated_at() is shared by users + auth_identities triggers.
-- Both are gone by the time this (the last) down migration runs:
--   000003 down drops auth_tokens (no trigger on it),
--   000002 down drops trg_auth_identities_updated_at + auth_identities,
--   000001 down drops trg_users_updated_at + users (above),
-- so it is now safe to drop the function.
DROP FUNCTION IF EXISTS set_updated_at();
```

### `migrations/000003_create_auth_tokens.down.sql`

```sql
DROP INDEX IF EXISTS ix_auth_tokens_valid;
DROP INDEX IF EXISTS ix_auth_tokens_user_purpose;
DROP INDEX IF EXISTS ux_auth_tokens_token_hash;
DROP TABLE IF EXISTS auth_tokens;
-- set_updated_at() is NOT dropped here: this is the first down to run
-- (reverse order), and triggers on users/auth_identities still depend
-- on it. It is dropped in 000001 down, after all triggers are gone.
```

### Reverse-run order (the key invariant)

```
000003 down:  drop auth_tokens indexes + table           (no trigger on auth_tokens)
000002 down:  drop trg_auth_identities_updated_at + auth_identities
000001 down:  drop trg_users_updated_at + users + set_updated_at()   ← safe now
```

After this change, at every `DROP FUNCTION` moment, zero triggers
reference it.

## 5. Verification (manual, against dev Postgres)

From a clean applied state (all three up migrations applied):

```bash
# 1. Confirm clean state
make migrate-up          # should say "no change"

# 2. Roll everything back — must succeed without "cannot drop function" error
make migrate-down        # runs down -all (000003 → 000002 → 000001)

# 3. Confirm the function is gone and tables are gone
#    (psql: \df, \dt — expect no set_updated_at, no users/auth_identities/auth_tokens)

# 4. Re-apply — must succeed
make migrate-up

# 5. Schema sanity
#    psql: \d users, \d auth_identities, \d auth_tokens — expect all back
```

If step 2 fails, **do not edit the applied migrations further blindly** —
capture the exact error and the `schema_migrations` table state
(`SELECT * FROM schema_migrations`) before proceeding; a dirty state
needs `migrate force` to recover, which is a manual ops step.

## 6. Tests

No Go test for this — migration reversibility is verified by the manual
`migrate-down && migrate-up` round-trip above. The
`repository_db_integration_test.go` suite already covers the schema in
its up state; the down path is operational, not unit-testable here.

## 7. Risk note

- **Assumptions made:** no environment has the *old* `000003.down` (with
  the function drop) applied as a recorded migration version. Dev was
  applied up-only; if any other env ran the broken down and is now
  dirty, it needs manual `migrate force` recovery **before** this fix
  is applied there.
- **Edge cases intentionally NOT handled:** none.
- **Concurrency assumptions:** none (migrations are run by an operator,
  not concurrently).
- **What is not tested, and why:** the down round-trip is the test;
  Go-level testing of DDL reversibility is not idiomatic here.

## 8. Non-goals

- Do not edit up migrations (already applied in dev; per
  `migrations-safety.md`, applied migrations are never edited —
  corrections go into a new migration. The down edits here are
  acceptable because the down path was never successfully applied).
- Do not add `CASCADE` to `DROP FUNCTION` (masks the dependency; the
  fix is to drop in the right order, not to force).
