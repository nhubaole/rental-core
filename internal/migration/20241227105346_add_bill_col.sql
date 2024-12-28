-- +goose Up
-- +goose StatementBegin
ALTER TABLE "billing"
ADD COLUMN "payment_id" INT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "billing"
DROP COLUMN "payment_id";
-- +goose StatementEnd
