package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Shipment represents a shipment entity.
type Shipment struct {
	ID          string
	UserID      int64
	Origin      string
	Destination string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

// NewShipment creates a new shipment with the given origin and destination.
func NewShipment(ctx context.Context, userID int64, origin, destination string) (*Shipment, error) {
	if origin == "" {
		return nil, errors.New("origin cannot be empty")
	}

	if destination == "" {
		return nil, errors.New("destination cannot be empty")
	}

	now := time.Now().UTC()
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	shipment := Shipment{
		ID:          id.String(),
		UserID:      userID,
		Origin:      origin,
		Destination: destination,
		CreatedAt:   &now,
	}

	return &shipment, nil
}

// NewShipmentRequest represents a request to create a new shipment.
type NewShipmentRequest struct {
	Email       string
	Role        string
	Origin      string
	Destination string
}

// ShipmentCreatedResult represents the result of creating a new shipment.
type ShipmentCreatedResult struct {
	ID string
}

// ShipmentDetails represents the details of a shipment.
type ShipmentDetails struct {
	Shipment
	Statuses []*ShipmentStatus
}

// ShipmentWithCurrentStatus represents a shipment with its current status.
type ShipmentWithCurrentStatus struct {
	ID          string
	Origin      string
	Destination string
	Status      ShipmentStatusList
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

// ShipmentEventRequest represents a request to update the status of a shipment.
type ShipmentEventRequest struct {
	TrackingNumber string
	Status         string
	Timestamp      time.Time
	Source         string
	Location       struct {
		Lat float64
		Lng float64
	}
}

type EventStatusResult struct {
	Status string
}
