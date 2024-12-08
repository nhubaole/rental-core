-- +goose Up
-- +goose StatementBegin
ALTER TABLE "contracts"
DROP COLUMN "contract_template_id",
DROP COLUMN "status";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "contracts"
ADD COLUMN "contract_template_id" integer NOT NULL,
ADD COLUMN "status" integer NOT NULL;
-- +goose StatementEnd
