-- +goose Up
-- +goose StatementBegin
CREATE TABLE "contract_templates" (
  "id" SERIAL PRIMARY KEY,
  "party_a" integer NOT NULL,
  "address" varchar[] NOT NULL,
  "electricity_method" varchar NOT NULL,
  "electricity_cost" float NOT NULL,
  "water_method" varchar NOT NULL,
  "water_cost" float NOT NULL,
  "internet_cost" float NOT NULL,
  "parking_fee" float NOT NULL,
  "responsibility_a" varchar NOT NULL,
  "responsibility_b" varchar NOT NULL,
  "general_responsibility" varchar NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

-- Foreign key linking contracts to contract_templates
ALTER TABLE "contracts"
ADD COLUMN "contract_template_id" integer,
ADD CONSTRAINT fk_contract_template
  FOREIGN KEY ("contract_template_id") REFERENCES "contract_templates"("id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Remove foreign key from contracts table first
ALTER TABLE "contracts" 
DROP CONSTRAINT IF EXISTS fk_contract_template,
DROP COLUMN IF EXISTS "contract_template_id";

DROP TABLE IF EXISTS "contract_templates";

-- +goose StatementEnd
