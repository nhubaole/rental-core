-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_sockets (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id),
  socket_id INTEGER,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(), 
  UNIQUE(user_id, socket_id) 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_sockets;
-- +goose StatementEnd
