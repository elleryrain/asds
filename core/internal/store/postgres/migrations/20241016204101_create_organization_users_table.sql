-- +goose Up
CREATE TABLE IF NOT EXISTS organization_users (
    organization_id BIGINT NOT NULL,
    user_id         BIGINT NOT NULL UNIQUE,
    is_owner        BOOLEAN NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS organization_users;