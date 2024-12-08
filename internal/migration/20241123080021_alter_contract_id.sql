-- +goose Up
-- +goose StatementBegin
ALTER TABLE "payments" 
ALTER COLUMN "contract_id" DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "payments" 
ALTER COLUMN "contract_id" SET NOT NULL;
-- +goose StatementEnd
