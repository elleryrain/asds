-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS employee (
    user_id         BIGINT NOT NULL REFERENCES users(id),
    role            SMALLINT NOT NULL,
    position        TEXT NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

WITH new_user AS (
    INSERT INTO users (email, phone, password_hash, first_name, last_name, middle_name, email_verified, totp_salt, is_banned)
    VALUES ('admin@ubrato.ru', '+79999999999', '$2a$10$B36B0bYuGJLq63MQzlSoxuUvWeFHNG62Z02e6XSdOgyOA522gIZIy', 'Admin', 'Admin', 'Admin', TRUE, gen_random_uuid(), FALSE)
    RETURNING id
)
INSERT INTO employee (user_id, role, position)
SELECT id, 4, 'Admin' FROM new_user;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS employee;
-- +goose StatementEnd
