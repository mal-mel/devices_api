-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX vendor_name_idx ON vendor ((lower(name)));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX vendor_name_idx;
-- +goose StatementEnd
