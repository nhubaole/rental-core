-- +goose Up
-- +goose StatementBegin
ALTER TABLE "billing"
ADD COLUMN "old_water_index" INTEGER,
ADD COLUMN "old_electricity_index" INTEGER,
ADD COLUMN "new_water_index" INTEGER,
ADD COLUMN "new_electricity_index" INTEGER,
ADD COLUMN "total_water_cost" FLOAT,
ADD COLUMN "total_electricity_cost" FLOAT,
DROP COLUMN "index_id";

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "billing"
DROP COLUMN "old_water_index",
DROP COLUMN "old_electricity_index",
DROP COLUMN "new_water_index",
DROP COLUMN "new_electricity_index",
DROP COLUMN "total_water_cost",
DROP COLUMN "total_electricity_cost";
-- +goose StatementEnd
