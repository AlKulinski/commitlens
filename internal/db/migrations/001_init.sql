-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS snapshots (
    id UUID PRIMARY KEY,
    path TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    hash TEXT,
    size INTEGER,
    mtime INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS snapshots;
-- +goose StatementEnd