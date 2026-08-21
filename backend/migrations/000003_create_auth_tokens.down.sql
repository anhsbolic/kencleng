DROP INDEX IF EXISTS ix_auth_tokens_valid;
DROP INDEX IF EXISTS ix_auth_tokens_user_purpose;
DROP INDEX IF EXISTS ux_auth_tokens_token_hash;
DROP TABLE IF EXISTS auth_tokens;
-- set_updated_at() is NOT dropped here: golang-migrate runs down in
-- reverse order (000003 → 000002 → 000001), and at this point triggers
-- on users + auth_identities still depend on the function. Dropping it
-- here would fail with "cannot drop function because other objects
-- depend on it" and leave the schema in a dirty state. The function is
-- dropped in 000001 down, after all triggers using it are gone.
