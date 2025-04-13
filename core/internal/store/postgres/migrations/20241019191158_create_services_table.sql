-- +goose Up
CREATE TABLE IF NOT EXISTS services (
    id          BIGSERIAL PRIMARY KEY,
    parent_id   BIGINT REFERENCES services(id) ON DELETE CASCADE,
    name        TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS services;
