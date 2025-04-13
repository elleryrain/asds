-- +goose Up
CREATE TABLE IF NOT EXISTS objects (
    id          BIGSERIAL PRIMARY KEY,
    parent_id   BIGINT REFERENCES objects(id) ON DELETE CASCADE,
    name        TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS objects;
