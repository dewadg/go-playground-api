-- +goose Up
-- +goose StatementBegin
CREATE TABLE items(
    id VARCHAR(255) PRIMARY KEY,
    data TEXT,
    created_at TEXT,
    assigned_at TEXT DEFAULT NULL
);

CREATE INDEX items_assigned_at ON items(assigned_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE items;
-- +goose StatementEnd
