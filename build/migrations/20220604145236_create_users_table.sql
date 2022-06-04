-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

CREATE INDEX users_email_deleted_at ON users(email, deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
