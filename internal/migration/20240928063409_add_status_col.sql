-- +goose Up
-- +goose StatementBegin
ALTER TABLE "billing" ADD COLUMN status INTEGER;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "billing" DROP COLUMN status;
-- +goose StatementEnd
