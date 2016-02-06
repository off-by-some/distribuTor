
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE OR REPLACE FUNCTION update_timestamp() RETURNS TRIGGER AS $body$
  BEGIN NEW.updated_at = NOW(); RETURN NEW; END; $body$ language plpgsql;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP FUNCTION update_timestamp;
