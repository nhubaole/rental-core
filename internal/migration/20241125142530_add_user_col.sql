-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users"
ADD COLUMN "private_key_hex" TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users"
DROP COLUMN "private_key_hex";
-- +goose StatementEnd
