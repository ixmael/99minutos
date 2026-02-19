//go:build integration
// +build integration

package bdd

import (
	"github.com/ixmael/99minutos/internal/core/domain"
	"github.com/ixmael/99minutos/internal/core/ports"
)

// TestContext represents the test context for BDD scenarios
type TestContext struct {
	UserEmail                     *string
	LastError                     error
	LastShipment                  *domain.ShipmentCreatedResult
	LastShipmentCreatedResult     *domain.ShipmentCreatedResult
	NewShipmentOriginRequest      *string
	NewShipmentDestinationRequest *string
	ShipmentRepository            ports.ShipmentRepository
	ShipmentStatusRepository      ports.ShipmentStatusRepository
	UserRepository                ports.UserRepository
	ShipmentService               ports.ShipmentService
	LastShipmentsWithStatuses     []*domain.ShipmentWithCurrentStatus
	LastShipmentDetails           *domain.ShipmentDetails
}

// NewTestContext creates a new test context with all required dependencies
func NewTestContext() *TestContext {
	return &TestContext{
		UserEmail:                     nil,
		LastError:                     nil,
		LastShipmentCreatedResult:     nil,
		NewShipmentOriginRequest:      nil,
		NewShipmentDestinationRequest: nil,
		UserRepository:                nil,
		ShipmentRepository:            nil,
		ShipmentStatusRepository:      nil,
		ShipmentService:               nil,
		LastShipmentsWithStatuses:     nil,
		LastShipmentDetails:           nil,
		LastShipment:                  nil,
	}
}
