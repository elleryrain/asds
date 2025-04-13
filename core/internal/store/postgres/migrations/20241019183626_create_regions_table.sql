-- +goose Up
CREATE TABLE IF NOT EXISTS regions (
    id      BIGSERIAL PRIMARY KEY,
    name    TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS regions;
