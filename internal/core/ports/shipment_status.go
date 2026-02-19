package ports

import (
	"context"

	"github.com/ixmael/99minutos/internal/core/domain"
)

// ShipmentService defines the interface for shipment-related operations.
type ShipmentStatusService interface {
	UpdateShipment(ctx context.Context, newShipmentRequest *domain.NewShipmentRequest) (*domain.ShipmentCreatedResult, error)
}

type ShipmentStatusRepository interface {
	Save(ctx context.Context, shipmentStatus *domain.ShipmentStatus) error
	FindStatusByID(ctx context.Context, shipmentID string) ([]*domain.ShipmentStatus, error)
	FindCurrentStatusByID(ctx context.Context, shipmentID string) (*domain.ShipmentStatus, error)
	Stop()
}
