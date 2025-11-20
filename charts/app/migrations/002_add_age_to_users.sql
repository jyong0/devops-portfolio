-- +goose Up
ALTER TABLE users
ADD COLUMN age INT;

-- +goose Down
-- NO OP
