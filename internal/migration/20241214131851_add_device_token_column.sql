-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users"
ADD COLUMN "device_token" TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users"
DROP COLUMN "device_token";
-- +goose StatementEnd
