-- +goose Up
CREATE TABLE IF NOT EXISTS additions (
    id                  BIGSERIAL PRIMARY KEY,
    tender_id           BIGINT REFERENCES tenders(id),
    title               TEXT NOT NULL,
    content             TEXT NOT NULL,
    attachments         TEXT[],
    verification_status SMALLINT NOT NULL,
    created_at          TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS additions;