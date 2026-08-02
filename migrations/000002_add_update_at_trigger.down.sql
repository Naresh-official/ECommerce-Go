DROP TRIGGER IF EXISTS users_set_updated_at ON users;
DROP TRIGGER IF EXISTS addresses_set_updated_at ON addresses;

DROP FUNCTION IF EXISTS set_updated_at();