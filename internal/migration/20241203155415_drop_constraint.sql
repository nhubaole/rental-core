-- +goose Up
-- +goose StatementBegin
ALTER TABLE "return_requests" 
DROP CONSTRAINT IF EXISTS return_requests_contract_id_fkey;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
