# Task 01: Migration 000010 + contract bundle regeneration (schema & contract pre-settle)

> Back-reference : `.local-agents/works/account/05-account-linking/2-plan/techplan.md` (Status: Approved) — sections 6 (Backward Compatibility), 8 (DB Schema + API changes), 9 (steps 1 & 6), 5 (D7, D9)
> Depends on    : nothing (first task in the chain)
> Model         : DeepSeek V4 Pro (precision DDL + reversibility subtleties; bundle step is mechanical)

## Objective

Settle the two non-Go prerequisites every later task assumes: (a) migration 000010 widening `auth_tokens.purpose` so Branch-1 tokens can carry `email_verification_link`, proven reversible in both directions; (b) regenerate the stale generated bundle `api/openapi.yaml` so `components.securitySchemes.bearerAuth` (defined in sources at `api/openapi/common.yaml:2`) reappears.

## Files

| File | Change |
|---|---|
| `backend/migrations/000010_widen_auth_tokens_purpose.up.sql` | New |
| `backend/migrations/000010_widen_auth_tokens_purpose.down.sql` | New |
| `api/openapi.yaml` | Regenerated (never hand-edited) |

## Migration DDL (verbatim from techplan §8 — do not redesign)

```sql
-- up
ALTER TABLE auth_tokens DROP CONSTRAINT auth_tokens_purpose_check;
ALTER TABLE auth_tokens ADD CONSTRAINT auth_tokens_purpose_check
    CHECK (purpose IN ('email_verification', 'email_verification_link', 'password_reset'));

-- down (re-map FIRST: redemption is purpose-blind, so re-pointed rows stay semantically valid)
UPDATE auth_tokens SET purpose = 'email_verification'
    WHERE purpose = 'email_verification_link';
ALTER TABLE auth_tokens DROP CONSTRAINT auth_tokens_purpose_check;
ALTER TABLE auth_tokens ADD CONSTRAINT auth_tokens_purpose_check
    CHECK (purpose IN ('email_verification', 'password_reset'));
```

`TBD — verify`: constraint name `auth_tokens_purpose_check` assumes PostgreSQL's automatic naming for the inline CHECK in `000003_create_auth_tokens.up.sql`. Confirm with `\d auth_tokens` after `make migrate-up`; if the actual name differs, adjust both scripts before proceeding.

## Steps

1. Write both migration files.
2. `make migrate-up`; verify constraint name + `\d auth_tokens` shows the 3-value CHECK.
3. Insert a probe row with `purpose='email_verification_link'`, run `make migrate-down` (must succeed via the re-map), confirm the row now reads `email_verification` and the CHECK is 2-value, delete probe row, `make migrate-up` again. Both directions proven per `postgresql/migrations-safety`.
4. `cd ../api && npm run bundle`; commit regenerated `openapi.yaml`.
5. **STOP condition (techplan D9-B)**: if the regenerated bundle still lacks `securitySchemes`, do NOT hand-edit the bundle — report the bundler behavior instead of "fixing" it.

## Rules covered

None directly (this task enables R1/R14 by making the new purpose value legal). Its correctness gates every later task.

## Common mistakes (from techplan §13, applicable subset)

| Mistake | Fix |
|---|---|
| Hand-editing `api/openapi.yaml` to restore `securitySchemes` | Regenerate via `npm run bundle` only (D9-B) |
| Writing a down migration that fails when `email_verification_link` rows exist | Re-map before restoring the 2-value CHECK |

## Verification

- `make migrate-down && make migrate-up` exits 0 twice consecutively
- `\d auth_tokens` shows expected CHECK post-up
- `grep -n "securitySchemes" api/openapi.yaml` finds the restored block
