package domain

import (
	"context"
	"errors"
	"time"
)

type ShipmentStatus string

const (
	ShipmentCreatedStatus ShipmentStatus = "CREATED"
)

// Shipment represents a shipment entity.
type Shipment struct {
	ID          string
	Origin      string
	Destination string
	Status      ShipmentStatus
	CreatedAt   *time.Time
}

// NewShipment creates a new shipment with the given origin and destination.
func NewShipment(ctx context.Context, origin, destination string) (*Shipment, error) {
	if origin == "" {
		return nil, errors.New("origin cannot be empty")
	}

	if destination == "" {
		return nil, errors.New("destination cannot be empty")
	}

	now := time.Now().UTC()
	shipment := Shipment{
		Origin:      origin,
		Destination: destination,
		Status:      ShipmentCreatedStatus,
		CreatedAt:   &now,
	}

	return &shipment, nil
}

// NewShipmentRequest represents a request to create a new shipment.
type NewShipmentRequest struct {
	Origin      string
	Destination string
}

// ShipmentCreatedResult represents the result of creating a new shipment.
type ShipmentCreatedResult struct {
	ID string
}
