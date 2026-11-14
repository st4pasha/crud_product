-- +goose Up
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY, 
    title TEXT, 
    price REAL);

-- +goose Down
SELECT 'down SQL query';

DROP TABLE products;
