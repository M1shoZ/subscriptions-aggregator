-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id) VALUES
  (uuid_generate_v4()),
  (uuid_generate_v4()),
  (uuid_generate_v4());

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM languages
-- +goose StatementEnd
