-- +goose Up
CREATE TABLE IF NOT EXISTS favourites (
    id              BIGSERIAL PRIMARY KEY,
    organization_id BIGINT NOT NULL REFERENCES organizations(id),
    object_type     TEXT NOT NULL,
    object_id       BIGINT NOT NULL,
    CONSTRAINT unique_favourites UNIQUE (organization_id, object_type, object_id)
);

-- +goose Down
DROP TABLE IF EXISTS favourites;