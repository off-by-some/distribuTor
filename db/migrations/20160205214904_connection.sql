
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE "connection" (
    control_port int,
    port int,
    updated_at timestamptz default NOW(),
    PRIMARY KEY(control_port),
    UNIQUE(control_port),
    UNIQUE(port)
);

CREATE TRIGGER update_connection_timestamp
BEFORE UPDATE ON "connection"
FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TRIGGER update_connection_timestamp;
DROP TABLE "connection";
