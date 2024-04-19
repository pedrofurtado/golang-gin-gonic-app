-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS products(
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR,
  description TEXT,
  price DECIMAL,
  quantity INTEGER,
  active BOOLEAN NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
DROP EXTENSION IF EXISTS "uuid-ossp";
-- +goose StatementEnd
