-- drop new items
DROP TABLE shipment_status;
DROP TYPE shipment_status_enum;

-- reverte changes on previous migration
CREATE TYPE shipment_status AS ENUM ('CREATED', 'PICKED_UP', 'IN_WAREHOUSE', 'IN_TRANSIT', 'DELIVERED');
ALTER TABLE shipment ADD COLUMN status shipment_status NOT NULL DEFAULT 'CREATED';
