-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS conversations (
    id SERIAL PRIMARY KEY,           -- Khóa chính với giá trị tự động tăng
    user_a INTEGER NOT NULL,          -- ID của người dùng A
    user_b INTEGER NOT NULL,          -- ID của người dùng B
    last_message_id INTEGER,          -- ID của tin nhắn cuối cùng
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Thời gian tạo conversation
);

CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,           -- Khóa chính với giá trị tự động tăng
    conversation_id INTEGER NOT NULL, -- ID của conversation mà tin nhắn thuộc về
    sender_id INTEGER NOT NULL,       -- ID của người gửi tin nhắn
    type INTEGER NOT NULL,            -- Loại tin nhắn (ví dụ: văn bản, hình ảnh, video, v.v.)
    content TEXT NOT NULL,            -- Nội dung tin nhắn
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Thời gian gửi tin nhắn
    FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE -- Ràng buộc khóa ngoại
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS messages;

DROP TABLE IF EXISTS conversations;
-- +goose StatementEnd
