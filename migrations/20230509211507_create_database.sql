-- +goose Up
-- +goose StatementBegin
CREATE TABLE vendor(
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(256) NOT NULL
);

CREATE TABLE device (
    id UUID PRIMARY KEY NOT NULL,
    is_charging BOOL NOT NULL,
    battery_level float4 NOT NULL,
    vendor_id INT NOT NULL,
    tags JSONB DEFAULT NULL,
    CONSTRAINT vendor_fk FOREIGN KEY (vendor_id) REFERENCES vendor (id) ON DELETE CASCADE

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS device;
DROP TABLE IF EXISTS vendor;
-- +goose StatementEnd
