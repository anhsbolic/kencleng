DROP TRIGGER IF EXISTS trg_users_updated_at ON users;
DROP TABLE IF EXISTS users;
-- set_updated_at() is shared by the users + auth_identities triggers.
-- golang-migrate runs down in reverse order (000003 → 000002 → 000001),
-- so by the time this (the last) down migration runs, both triggers are
-- already gone:
--   000003 down drops auth_tokens (no trigger on it),
--   000002 down drops trg_auth_identities_updated_at + auth_identities,
--   000001 down drops trg_users_updated_at + users (above),
-- so it is now safe to drop the function.
DROP FUNCTION IF EXISTS set_updated_at();
