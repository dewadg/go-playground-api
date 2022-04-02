-- +goose Up
-- +goose StatementBegin
CREATE TABLE items(
    id VARCHAR(255) PRIMARY KEY,
    data JSON,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE items;
-- +goose StatementEnd
