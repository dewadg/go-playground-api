-- +goose Up
-- +goose StatementBegin
ALTER TABLE items
ADD assigned_at TIMESTAMP NULL DEFAULT NULL,
ADD INDEX items_assigned_at(assigned_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE items DROP COLUMN assigned_at;
-- +goose StatementEnd
