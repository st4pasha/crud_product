-- +goose Up
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS basket (
    product_id INTEGER,
    count INTEGER
);
-- +goose Down
SELECT 'down SQL query';

DROP TABLE IF EXISTS basket
