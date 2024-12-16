-- +goose Up
-- +goose StatementBegin
ALTER TABLE "banks"
ADD COLUMN "short_name" TEXT,
ADD COLUMN "logo" TEXT,
DROP COLUMN "country";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "banks"
DROP COLUMN "short_name",
DROP COLUMN "logo",
ADD COLUMN "country" TEXT;
-- +goose StatementEnd
