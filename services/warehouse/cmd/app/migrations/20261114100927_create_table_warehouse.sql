-- +goose Up
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS warehouse (id INTEGER, count INTEGER);
-- +goose Down
SELECT 'down SQL query';
DROP TABLE warehouse;