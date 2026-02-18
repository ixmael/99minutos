ALTER TABLE shipment DROP COLUMN status;
DROP TYPE shipment_status;

CREATE TYPE shipment_status_enum AS ENUM ('CREATED', 'PICKED_UP', 'IN_WAREHOUSE', 'IN_TRANSIT', 'DELIVERED');

CREATE TABLE shipment_status (
    shipment_id VARCHAR(128) NOT NULL,
    status shipment_status_enum NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
