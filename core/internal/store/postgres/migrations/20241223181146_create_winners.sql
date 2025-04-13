-- +goose Up
CREATE TABLE IF NOT EXISTS winners (
    id              BIGSERIAL    PRIMARY KEY,
    organization_id BIGINT       NOT NULL REFERENCES organizations(id),
    tender_id       BIGINT       NOT NULL REFERENCES tenders(id),
    accepted        SMALLINT     NOT NULL DEFAULT 1,
    added_at        TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS winners;
