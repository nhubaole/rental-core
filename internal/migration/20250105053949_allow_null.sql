-- +goose Up
-- +goose StatementBegin
ALTER TABLE rental_requests
  ALTER COLUMN begin_date DROP NOT NULL,
  ALTER COLUMN end_date DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE rental_requests
  ALTER COLUMN begin_date SET NOT NULL,
  ALTER COLUMN end_date SET NOT NULL;
-- +goose StatementEnd
