-- +goose Up
-- +goose StatementBegin
CREATE TABLE messages (
  id SERIAL PRIMARY KEY,
  sender_id INTEGER REFERENCES users(id),
  receiver_id INTEGER REFERENCES users(id),
  type INTEGER NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMPTZ 
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS messages;
-- +goose StatementEnd
