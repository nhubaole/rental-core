-- +goose Up
-- +goose StatementBegin
ALTER TABLE "contracts"
ADD COLUMN "signature_a" TEXT NULL,
ADD COLUMN "signature_b" TEXT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "contracts"
DROP COLUMN "signature_a",
DROP COLUMN "signature_b";
-- +goose StatementEnd
