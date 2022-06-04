-- +goose Up
-- +goose StatementBegin
CREATE TABLE access_tokens(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    token VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deactivated_at TIMESTAMP NULL DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE access_tokens;
-- +goose StatementEnd
