-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "users" (
  id SERIAL PRIMARY KEY,
  phone_number varchar NOT NULL,
  full_name varchar NOT NULL,
  password varchar NOT NULL,
  address varchar,
  role integer NOT NULL,
  created_at timestamp NOT NULL,
  deleted_at timestamp 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
