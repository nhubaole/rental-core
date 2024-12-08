-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users"
ALTER COLUMN "wallet_address" DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users"
ALTER COLUMN "wallet_address" SET NOT NULL;
-- +goose StatementEnd
