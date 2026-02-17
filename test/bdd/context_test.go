//go:build integration
// +build integration

package bdd

import (
	"github.com/ixmael/99minutos/internal/core/domain"
	"github.com/ixmael/99minutos/internal/core/ports"
)

// TestContext represents the test context for BDD scenarios
type TestContext struct {
	LastError                     error
	LastShipmentCreatedResult     *domain.ShipmentCreatedResult
	NewShipmentOriginRequest      *string
	NewShipmentDestinationRequest *string
	ShipmentRepository            ports.ShipmentRepository
	ShipmentService               ports.ShipmentService
}

// NewTestContext creates a new test context with all required dependencies
func NewTestContext() *TestContext {
	return &TestContext{
		LastError:                     nil,
		LastShipmentCreatedResult:     nil,
		NewShipmentOriginRequest:      nil,
		NewShipmentDestinationRequest: nil,
		ShipmentRepository:            nil,
		ShipmentService:               nil,
	}
}
