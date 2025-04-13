-- +goose Up
CREATE TABLE IF NOT EXISTS external_id_amo_id (
    external_id INTEGER,
    amo_id INTEGER,
    entity TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (external_id, amo_id, entity)
);

-- +goose Down
DROP TABLE IF EXISTS external_id_amo_id;