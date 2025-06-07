-- -- +goose Up
-- -- +goose StatementBegin
-- SELECT 'up SQL query';
-- -- +goose StatementEnd
--
-- -- +goose Down
-- -- +goose StatementBegin
-- SELECT 'down SQL query';
-- -- +goose StatementEnd

CREATE TABLE IF NOT EXISTS tablets (
                         id SERIAL PRIMARY KEY,
                         title TEXT NOT NULL,
                         price FLOAT NOT NULL,
                         description TEXT NOT NULL,
                         rating INTEGER NOT NULL,
                         image_url TEXT NOT NULL
);