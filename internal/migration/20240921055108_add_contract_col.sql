-- +goose Up
-- +goose StatementBegin
ALTER TABLE "contracts"
ADD COLUMN "status" INTEGER;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "contracts"
DROP COLUMN "status";
-- +goose StatementEnd
