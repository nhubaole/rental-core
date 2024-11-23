-- +goose Up
-- +goose StatementBegin
ALTER TABLE "payments" 
DROP CONSTRAINT IF EXISTS payments_return_request_id_fkey,
DROP CONSTRAINT IF EXISTS payments_contract_id_fkey;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
