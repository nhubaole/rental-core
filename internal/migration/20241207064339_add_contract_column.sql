-- +goose Up
-- +goose StatementBegin
ALTER TABLE "contracts"
ADD COLUMN "room_id" integer;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "contracts"
DROP COLUMN "room_id";
-- +goose StatementEnd
