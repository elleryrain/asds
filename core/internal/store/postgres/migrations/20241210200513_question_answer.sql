-- +goose Up
CREATE TABLE IF NOT EXISTS question_answer (
    id                     BIGSERIAL PRIMARY KEY,
    tender_id              BIGINT    NOT NULL REFERENCES tenders(id) ON DELETE CASCADE,
    author_organization_id BIGINT    NOT NULL REFERENCES organizations(id),
    parent_id              BIGINT    UNIQUE REFERENCES question_answer(id) ON DELETE CASCADE,
    type                   SMALLINT  NOT NULL CHECK (type IN (1, 2)),
    content                TEXT      NOT NULL,
    verification_status    SMALLINT  NOT NULL,
    created_at             TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CHECK (type != 1 OR parent_id IS NULL)
);

-- +goose Down
DROP TABLE IF EXISTS questions_answers;
