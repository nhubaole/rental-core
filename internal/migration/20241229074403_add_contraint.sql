-- +goose Up
-- +goose StatementBegin
ALTER TABLE PUBLIC.INDEX
ADD CONSTRAINT index_unique UNIQUE (room_id, month, year);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
