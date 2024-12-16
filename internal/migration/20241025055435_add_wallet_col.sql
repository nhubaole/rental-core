-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users"
ADD COLUMN "wallet_address" TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users"
DROP COLUMN "wallet_address";
-- +goose StatementEnd
