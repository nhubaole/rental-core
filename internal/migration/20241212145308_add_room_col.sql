-- +goose Up
-- +goose StatementBegin
ALTER TABLE "rooms"
ADD COLUMN "available_from" timestamptz;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "rooms"
DROP COLUMN "available_from";
-- +goose StatementEnd
