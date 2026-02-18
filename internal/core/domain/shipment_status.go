package domain

import (
	"time"
)

type ShipmentStatusList string

const (
	ShipmentCreatedStatus     ShipmentStatusList = "CREATED"
	ShipmentPickedUpStatus    ShipmentStatusList = "PICKED_UP"
	ShipmentInWarehouseStatus ShipmentStatusList = "IN_WAREHOUSE"
	ShipmentInTransitStatus   ShipmentStatusList = "IN_TRANSIT"
	ShipmentDeliveredStatus   ShipmentStatusList = "DELIVERED"
	ShipmentCancelledStatus   ShipmentStatusList = "CANCELLED"
)

// ShipmentStatus represents a shipment status.
type ShipmentStatus struct {
	ID        string
	Status    ShipmentStatusList
	CreatedAt *time.Time
}

// NewCreatedShipmentStatus creates a new shipment status with the CREATED status.
func NewCreatedShipmentStatus(shipmentID string) (*ShipmentStatus, error) {
	now := time.Now()
	shipmentstatus := &ShipmentStatus{
		ID:        shipmentID,
		Status:    ShipmentCreatedStatus,
		CreatedAt: &now,
	}

	return shipmentstatus, nil
}
