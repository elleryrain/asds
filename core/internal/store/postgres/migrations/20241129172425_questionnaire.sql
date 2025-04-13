-- +goose Up
CREATE TABLE IF NOT EXISTS questionnaire (
    id BIGSERIAL PRIMARY KEY,       
    organization_id INT NOT NULL UNIQUE REFERENCES organizations(id) ON DELETE CASCADE,   
    answers JSONB NOT NULL,               
    is_completed BOOLEAN DEFAULT FALSE,      
    completed_at TIMESTAMP DEFAULT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS questionnaire;
