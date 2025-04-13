-- +goose Up
CREATE TABLE IF NOT EXISTS tender_responses (
    id              BIGSERIAL   PRIMARY KEY,
    tender_id       BIGINT      NOT NULL REFERENCES tenders(id),
    organization_id BIGINT      NOT NULL REFERENCES organizations(id),
    price           INT         NOT NULL,
    is_nds_price    BOOLEAN     NOT NULL,
    is_winner       BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_tender_organization UNIQUE (tender_id, organization_id)
);

-- +goose Down
DROP TABLE IF EXISTS tender_responses;
