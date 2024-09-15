-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "phone_number" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "password" varchar NOT NULL,
  "address" varchar,
  "role" integer NOT NULL,
  "created_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "rooms" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar NOT NULL,
  "address" text[] NOT NULL,
  "room_number" integer NOT NULL,
  "room_images" text[] NOT NULL,
  "utilities" text[] NOT NULL,
  "description" varchar NOT NULL,
  "room_type" varchar,
  "owner" integer NOT NULL,
  "capacity" integer NOT NULL,
  "gender" integer,
  "area" float NOT NULL,
  "total_price" float,
  "deposit" float NOT NULL,
  "electricity_cost" float NOT NULL,
  "water_cost" float NOT NULL,
  "internet_cost" float NOT NULL,
  "is_parking" bool NOT NULL,
  "parking_fee" float,
  "status" integer NOT NULL,
  "is_rent" bool NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz 
);

CREATE TABLE "tenants" (
  "id" SERIAL PRIMARY KEY,
  "room_id" integer NOT NULL,
  "tenant_id" integer NOT NULL,
  "begin_date" timestamptz NOT NULL,
  "end_date" timestamptz,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "like" (
  "id" SERIAL PRIMARY KEY,
  "room_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "rental_requests" (
  "id" SERIAL PRIMARY KEY,
  "code" varchar NOT NULL,
  "sender_id" integer NOT NULL,
  "room_id" integer NOT NULL,
  "suggested_price" float,
  "num_of_person" integer,
  "begin_date" timestamptz NOT NULL,
  "end_date" timestamptz NOT NULL,
  "addition_request" varchar,
  "status" integer NOT NULL,
  -- 1: pending
  -- 2: approved
  -- 3. declined
  "created_at" timestamptz,
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE TABLE "process_tracking" (
  "id" SERIAL PRIMARY KEY,
  "actor" integer NOT NULL,
  "action" varchar NOT NULL,
  "issued_at" timestamptz NOT NULL,
  "request_id" integer NOT NULL
);

CREATE TABLE "contracts" (
  "id" SERIAL PRIMARY KEY,
  "code" varchar NOT NULL,
  "party_a" integer NOT NULL,
  "party_b" integer NOT NULL,
  "request_id" integer NOT NULL,
  "room_id" integer NOT NULL,
  "actual_price" float NOT NULL,
  "payment_method" varchar,
  "electricity_method" varchar NOT NULL,
  "electricity_cost" float NOT NULL,
  "water_method" varchar NOT NULL,
  "water_cost" float NOT NULL,
  "internet_cost" float NOT NULL,
  "parking_fee" float,
  "deposit" float NOT NULL,
  "begin_date" date NOT NULL,
  "end_date" date NOT NULL,
  "responsibility_a" varchar NOT NULL,
  "responsibility_b" varchar NOT NULL,
  "general_responsibility" varchar,
  "signature_a" varchar NOT NULL,
  "signed_time_a" timestamptz NOT NULL,
  "signature_b" varchar NOT NULL,
  "signed_time_b" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamp
);

CREATE TABLE "payments" (
  "id" SERIAL PRIMARY KEY,
  "code" varchar NOT NULL,
  "sender_id" integer NOT NULL,
  "bill_id" integer NOT NULL,
  "contract_id" integer NOT NULL,
  "amount" float NOT NULL,
  "status" integer NOT NULL
);

CREATE TABLE "transactions" (
  "id" SERIAL PRIMARY KEY,
  "payment_id" integer NOT NULL,
  "sender_id" integer NOT NULL,
  "amount" float NOT NULL,
  "status" integer NOT NULL,
  "transaction_type" integer,
  "gateway_response" varchar
);

CREATE TABLE "index" (
  "id" SERIAL PRIMARY KEY,
  "water_index" float NOT NULL,
  "electricity_index" float NOT NULL,
  "room_id" integer NOT NULL,
  "month" integer NOT NULL,
  "year" integer NOT NULL
);

CREATE TABLE "billing" (
  "id" SERIAL PRIMARY KEY,
  "code" varchar NOT NULL,
  "contract_id" integer NOT NULL,
  "index_id" integer NOT NULL,
  "addition_fee" integer,
  "addition_note" varchar,
  "total_amount" float NOT NULL,
  "month" integer NOT NULL,
  "year" integer NOT NULL,
  "paid_time" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

ALTER TABLE "rooms" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "tenants" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

ALTER TABLE "tenants" ADD FOREIGN KEY ("tenant_id") REFERENCES "users" ("id");

ALTER TABLE "like" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

ALTER TABLE "like" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "rental_requests" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

ALTER TABLE "rental_requests" ADD FOREIGN KEY ("sender_id") REFERENCES "users" ("id");

ALTER TABLE "process_tracking" ADD FOREIGN KEY ("request_id") REFERENCES "rental_requests" ("id");

ALTER TABLE "contracts" ADD FOREIGN KEY ("party_a") REFERENCES "users" ("id");

ALTER TABLE "contracts" ADD FOREIGN KEY ("party_b") REFERENCES "users" ("id");

ALTER TABLE "contracts" ADD FOREIGN KEY ("request_id") REFERENCES "rental_requests" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("sender_id") REFERENCES "users" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("bill_id") REFERENCES "billing" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("contract_id") REFERENCES "contracts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("payment_id") REFERENCES "payments" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("sender_id") REFERENCES "users" ("id");

ALTER TABLE "billing" ADD FOREIGN KEY ("index_id") REFERENCES "index" ("id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE "billing" DROP CONSTRAINT IF EXISTS "billing_index_id_fkey";
ALTER TABLE "transactions" DROP CONSTRAINT IF EXISTS "transactions_sender_id_fkey";
ALTER TABLE "transactions" DROP CONSTRAINT IF EXISTS "transactions_payment_id_fkey";
ALTER TABLE "payments" DROP CONSTRAINT IF EXISTS "payments_contract_id_fkey";
ALTER TABLE "payments" DROP CONSTRAINT IF EXISTS "payments_bill_id_fkey";
ALTER TABLE "payments" DROP CONSTRAINT IF EXISTS "payments_sender_id_fkey";
ALTER TABLE "contracts" DROP CONSTRAINT IF EXISTS "contracts_request_id_fkey";
ALTER TABLE "contracts" DROP CONSTRAINT IF EXISTS "contracts_party_b_fkey";
ALTER TABLE "contracts" DROP CONSTRAINT IF EXISTS "contracts_party_a_fkey";
ALTER TABLE "process_tracking" DROP CONSTRAINT IF EXISTS "process_tracking_request_id_fkey";
ALTER TABLE "rental_requests" DROP CONSTRAINT IF EXISTS "rental_requests_sender_id_fkey";
ALTER TABLE "rental_requests" DROP CONSTRAINT IF EXISTS "rental_requests_room_id_fkey";
ALTER TABLE "like" DROP CONSTRAINT IF EXISTS "like_user_id_fkey";
ALTER TABLE "like" DROP CONSTRAINT IF EXISTS "like_room_id_fkey";
ALTER TABLE "tenants" DROP CONSTRAINT IF EXISTS "tenants_tenant_id_fkey";
ALTER TABLE "tenants" DROP CONSTRAINT IF EXISTS "tenants_room_id_fkey";
ALTER TABLE "rooms" DROP CONSTRAINT IF EXISTS "rooms_owner_fkey";

DROP TABLE IF EXISTS "billing";
DROP TABLE IF EXISTS "index";
DROP TABLE IF EXISTS "transactions";
DROP TABLE IF EXISTS "payments";
DROP TABLE IF EXISTS "contracts";
DROP TABLE IF EXISTS "process_tracking";
DROP TABLE IF EXISTS "rental_requests";
DROP TABLE IF EXISTS "like";
DROP TABLE IF EXISTS "tenants";
DROP TABLE IF EXISTS "rooms";
DROP TABLE IF EXISTS "users";

-- +goose StatementEnd

