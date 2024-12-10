-- +goose Up
-- +goose StatementBegin
CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,                    
    user_id INT NOT NULL,                     
    reference_id INT NOT NULL,                
    reference_type VARCHAR(50) NOT NULL,       
    title VARCHAR(255) NOT NULL,               
    is_read BOOLEAN DEFAULT FALSE,             
    created_at TIMESTAMP DEFAULT now(),        
    updated_at TIMESTAMP DEFAULT now()        
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notifications;
-- +goose StatementEnd
