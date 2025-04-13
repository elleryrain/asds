-- +goose Up
CREATE TABLE IF NOT EXISTS organizations (
    id                    BIGSERIAL PRIMARY KEY,
    brand_name            TEXT NOT NULL,
    full_name             TEXT NOT NULL,
    short_name            TEXT NOT NULL,
    inn                   TEXT NOT NULL UNIQUE,
    okpo                  TEXT NOT NULL,
    ogrn                  TEXT NOT NULL,
    kpp                   TEXT NOT NULL,
    tax_code              TEXT NOT NULL,
    address               TEXT NOT NULL,
    avatar_url            TEXT NULL,
    emails                JSONB NOT NULL,
    phones                JSONB NOT NULL,
    messengers            JSONB NOT NULL,
    verification_status   SMALLINT DEFAULT 1,
    is_contractor         BOOLEAN DEFAULT FALSE,
    is_banned             BOOLEAN DEFAULT FALSE,
    customer_info         JSONB NOT NULL,
    contractor_info       JSONB NOT NULL,
    created_at            TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at            TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS organizations;
