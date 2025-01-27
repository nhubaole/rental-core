-- +goose Up
-- +goose StatementBegin
ALTER TABLE "rooms"
ADD COLUMN "latitude" FLOAT NULL,
ADD COLUMN "longitude" FLOAT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "rooms"
DROP COLUMN "latitude",
DROP COLUMN "longitude";
-- +goose StatementEnd
