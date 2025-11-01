-- +goose Up
CREATE UNIQUE INDEX IF NOT EXISTS unique_flight ON voucher(flight_number, flight_date);

-- +goose Down
DROP INDEX IF EXISTS unique_flight;
