DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS sellers;

DROP TRIGGER IF EXISTS seller_set_updated_at ON sellers;
DROP TRIGGER IF EXISTS products_set_updated_at ON products;
