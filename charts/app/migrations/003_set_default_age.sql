-- +goose Up
UPDATE users
SET age = COALESCE(age, 20);

-- +goose Down
-- NO OP
