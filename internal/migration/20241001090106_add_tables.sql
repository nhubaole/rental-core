-- +goose Up
-- +goose StatementBegin
-- Table: return_requests
CREATE TABLE return_requests (
    id SERIAL PRIMARY KEY,
    contract_id INTEGER REFERENCES contracts(id),
    reason VARCHAR,
    return_date TIMESTAMP,
    status INTEGER,
    deduct_amount FLOAT,
    total_return_deposit FLOAT,
    created_user INTEGER REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Table: landlord_ratings
CREATE TABLE landlord_ratings (
    id SERIAL PRIMARY KEY,
    landlord_id INTEGER REFERENCES users(id),
    rated_by INTEGER REFERENCES users(id),
    friendliness_rating INTEGER,
    professionalism_rating INTEGER,
    support_rating INTEGER,
    transparency_rating INTEGER,
    overall_rating INTEGER,
    comments TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: tenant_ratings
CREATE TABLE tenant_ratings (
    id SERIAL PRIMARY KEY,
    tenant_id INTEGER REFERENCES users(id),
    rated_by INTEGER REFERENCES users(id),
    payment_rating INTEGER,
    property_care_rating INTEGER,
    neighborhood_disturbance_rating INTEGER,
    contract_compliance_rating INTEGER,
    overall_rating INTEGER,
    comments TEXT,
    images TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: room_ratings
CREATE TABLE room_ratings (
    id SERIAL PRIMARY KEY,
    room_id INTEGER REFERENCES rooms(id),
    rated_by INTEGER REFERENCES users(id),
    amenities_rating INTEGER,
    location_rating INTEGER,
    cleanliness_rating INTEGER,
    price_rating INTEGER,
    overall_rating INTEGER,
    comments TEXT,
    images TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS room_ratings;
DROP TABLE IF EXISTS tenant_ratings;
DROP TABLE IF EXISTS landlord_ratings;
DROP TABLE IF EXISTS return_requests;
-- +goose StatementEnd
