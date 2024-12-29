-- +goose Up
-- +goose StatementBegin
ALTER TABLE "index" ALTER COLUMN water_index DROP NOT NULL;
ALTER TABLE "index" ALTER COLUMN electricity_index DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
