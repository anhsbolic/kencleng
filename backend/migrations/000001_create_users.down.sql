DROP TRIGGER IF EXISTS trg_users_updated_at ON users;
DROP TABLE IF EXISTS users;
-- NOTE: set_updated_at() dropped only in the final down migration,
-- after all triggers using it are gone. See 000003 down.
