-- +goose Up
-- +goose StatementBegin
ALTER TABLE "billing" ALTER COLUMN paid_time DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "billing" ALTER COLUMN paid_time SET NOT NULL;
-- +goose StatementEnd
