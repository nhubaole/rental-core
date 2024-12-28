-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users"
ADD COLUMN "gender" INT,
ADD COLUMN "dob" DATE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users"
DROP COLUMN "gender",
DROP COLUMN "dob";
-- +goose StatementEnd
