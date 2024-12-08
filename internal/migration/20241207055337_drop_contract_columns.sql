-- +goose Up
-- +goose StatementBegin
ALTER TABLE "contracts"
DROP COLUMN "code",
DROP COLUMN "party_a",
DROP COLUMN "party_b",
DROP COLUMN "request_id",
DROP COLUMN "room_id",
DROP COLUMN "actual_price",
DROP COLUMN "payment_method",
DROP COLUMN "electricity_method",
DROP COLUMN "electricity_cost",
DROP COLUMN "water_method",
DROP COLUMN "water_cost",
DROP COLUMN "internet_cost",
DROP COLUMN "parking_fee",
DROP COLUMN "deposit",
DROP COLUMN "begin_date",
DROP COLUMN "end_date",
DROP COLUMN "responsibility_a",
DROP COLUMN "responsibility_b",
DROP COLUMN "general_responsibility",
DROP COLUMN "signature_a",
DROP COLUMN "signed_time_a",
DROP COLUMN "signature_b",
DROP COLUMN "signed_time_b",
DROP COLUMN "created_at",
DROP COLUMN "updated_at",
DROP COLUMN "deleted_at";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "contracts"
ADD COLUMN "code" varchar NOT NULL,
ADD COLUMN "party_a" integer NOT NULL,
ADD COLUMN "party_b" integer NOT NULL,
ADD COLUMN "request_id" integer NOT NULL,
ADD COLUMN "room_id" integer NOT NULL,
ADD COLUMN "actual_price" float NOT NULL,
ADD COLUMN "payment_method" varchar,
ADD COLUMN "electricity_method" varchar NOT NULL,
ADD COLUMN "electricity_cost" float NOT NULL,
ADD COLUMN "water_method" varchar NOT NULL,
ADD COLUMN "water_cost" float NOT NULL,
ADD COLUMN "internet_cost" float NOT NULL,
ADD COLUMN "parking_fee" float,
ADD COLUMN "deposit" float NOT NULL,
ADD COLUMN "begin_date" date NOT NULL,
ADD COLUMN "end_date" date NOT NULL,
ADD COLUMN "responsibility_a" varchar NOT NULL,
ADD COLUMN "responsibility_b" varchar NOT NULL,
ADD COLUMN "general_responsibility" varchar,
ADD COLUMN "signature_a" varchar NOT NULL,
ADD COLUMN "signed_time_a" timestamptz NOT NULL,
ADD COLUMN "signature_b" varchar,
ADD COLUMN "signed_time_b" timestamptz,
ADD COLUMN "created_at" timestamptz NOT NULL,
ADD COLUMN "updated_at" timestamptz NOT NULL,
ADD COLUMN "deleted_at" timestamp;
-- +goose StatementEnd