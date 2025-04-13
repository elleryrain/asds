-- +goose Up
CREATE TABLE IF NOT EXISTS measurements (
    id SERIAL PRIMARY KEY, 
    name VARCHAR(255) NOT NULL UNIQUE
);

INSERT INTO measurements (name)
VALUES 
-- Длина
('м'),       
('км'),      

-- Площадь
('м²'),     
('га'),     
('км²'),    

-- Объем
('м³'),      
('л'),       

-- Масса
('кг'),      
('т'),       

-- Время
('мин'),    
('ч'),      
('день'),
('неделя'),
('мес');   

-- +goose Down
DROP TABLE IF EXISTS measurements;