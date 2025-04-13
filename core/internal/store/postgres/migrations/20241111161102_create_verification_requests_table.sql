-- +goose Up
CREATE TABLE IF NOT EXISTS verification_requests (
    id                  BIGSERIAL PRIMARY KEY,
    reviewer_user_id    BIGINT REFERENCES users(id),
    object_type         SMALLINT NOT NULL,
    object_id           BIGINT NOT NULL,
    content             TEXT NULL,
    attachments         JSONB,
    status              SMALLINT NOT NULL,
    review_comment      TEXT NULL,
    created_at          TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    reviewed_at         TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS verification_requests;