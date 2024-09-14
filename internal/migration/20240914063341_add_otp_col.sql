-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users"
ADD COLUMN "otp" INTEGER;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users"
DROP COLUMN "otp";
-- +goose StatementEnd
