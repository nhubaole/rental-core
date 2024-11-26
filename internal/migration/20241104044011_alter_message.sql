-- +goose Up
-- +goose StatementBegin
ALTER TABLE messages
    ALTER COLUMN content DROP NOT NULL,                         -- Allow content to be NULL
    ADD COLUMN rent_auto_content JSONB DEFAULT '{}'::jsonb;     -- Add new column with default empty JSON object
        

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE messages
ALTER COLUMN content SET NOT NULL,  
DROP COLUMN rent_auto_content;         
-- +goose StatementEnd
