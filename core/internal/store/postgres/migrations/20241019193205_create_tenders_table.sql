-- +goose Up
CREATE TABLE IF NOT EXISTS tenders (
    id                      BIGSERIAL       PRIMARY KEY,
    organization_id         BIGINT          NOT NULL REFERENCES organizations(id),
    city_id                 INT             NOT NULL REFERENCES cities(id),
    services_ids            INT[]           NOT NULL,
    objects_ids             INT[]           NOT NULL,
    name                    TEXT            NOT NULL,
    price                   INT             NOT NULL,
    is_contract_price       BOOLEAN         NOT NULL,
    is_nds_price            BOOLEAN         NOT NULL,
    floor_space             INT             NOT NULL,
    description             TEXT,
    wishes                  TEXT,
    specification           TEXT            NOT NULL,
    attachments             TEXT[],
    status                  SMALLINT        NOT NULL,
    verification_status     SMALLINT        NOT NULL,
    is_draft                BOOLEAN         NOT NULL,
    reception_start         TIMESTAMPTZ     NOT NULL,
    reception_end           TIMESTAMPTZ     NOT NULL,
    work_start              TIMESTAMPTZ     NOT NULL,
    work_end                TIMESTAMPTZ     NOT NULL,
    created_at              TIMESTAMPTZ     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMPTZ     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS tenders;
