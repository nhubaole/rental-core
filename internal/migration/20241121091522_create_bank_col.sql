-- +goose Up
-- +goose StatementBegin
-- Tạo bảng banks
CREATE TABLE banks (
    id SERIAL PRIMARY KEY,
    bank_name TEXT NOT NULL,
    bank_code TEXT NOT NULL,
    country TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tạo bảng user_banks
CREATE TABLE user_banks (
    user_id INTEGER NOT NULL,
    bank_id INTEGER NOT NULL,
    account_number TEXT NOT NULL,
    account_name TEXT NOT NULL,
    card_number TEXT,
    currency TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, bank_id),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_bank FOREIGN KEY (bank_id) REFERENCES banks (id) ON DELETE CASCADE
);

-- chỉnh sửa khóa ngoại
ALTER TABLE conversations
ADD CONSTRAINT fk_user_a FOREIGN KEY (user_a) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE conversations
ADD CONSTRAINT fk_user_b FOREIGN KEY (user_b) REFERENCES users(id) ON DELETE CASCADE;

-- Thêm khóa ngoại vào bảng messages
ALTER TABLE messages
ADD CONSTRAINT fk_sender FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Xóa bảng user_banks
DROP TABLE IF EXISTS user_banks;

-- Xóa bảng banks
DROP TABLE IF EXISTS banks;
-- +goose StatementEnd
