-- +goose Up
CREATE TABLE IF NOT EXISTS cities (
    id          BIGSERIAL PRIMARY KEY,
    region_id   BIGINT REFERENCES regions(id),
    name        TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS cities;
