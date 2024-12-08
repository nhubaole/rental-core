-- +goose Up
-- +goose StatementBegin
ALTER TABLE "payments" 
ADD COLUMN "transfer_content" TEXT NULL,
ADD COLUMN "evidence_image" TEXT NULL,
ADD COLUMN "paid_time" timestamptz NOT NULL;

ALTER TABLE "billing"
DROP COLUMN "paid_time"; 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "payments" 
DROP COLUMN "transfer_content",
DROP COLUMN "evidence_image",
DROP COLUMN "paid_time"

ALTER TABLE "billing"
ADD COLUMN "paid_time" timestamptz NOT NULL; 
-- +goose StatementEnd
