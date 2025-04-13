-- +goose Up
CREATE TABLE IF NOT EXISTS notifications (
    id                    SERIAL PRIMARY KEY,
    user_id               BIGINT NOT NULl,
    title                 VARCHAR(255) NOT NULL,
    comment               TEXT, 
    action_button_text    VARCHAR(255),
    action_button_url     TEXT,  
    action_button_styled  BOOLEAN,      
    status                INT,               
    status_text           VARCHAR(255),
    is_read               BOOLEAN DEFAULT FALSE,
    created_at            TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS notifications;
