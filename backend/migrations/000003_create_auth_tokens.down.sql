DROP INDEX IF EXISTS ix_auth_tokens_valid;
DROP INDEX IF EXISTS ix_auth_tokens_user_purpose;
DROP INDEX IF EXISTS ux_auth_tokens_token_hash;
DROP TABLE IF EXISTS auth_tokens;
-- 000003 is the last table using set_updated_at(); drop the function now.
DROP FUNCTION IF EXISTS set_updated_at();
