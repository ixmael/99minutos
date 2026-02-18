package ports

import (
	"context"

	"github.com/ixmael/99minutos/internal/core/domain"
)

// ShipmentService defines the interface for shipment-related operations.
type ShipmentService interface {
	CreateShipment(ctx context.Context, newShipmentRequest *domain.NewShipmentRequest) (*domain.ShipmentCreatedResult, error)
	GetShipmentDetails(ctx context.Context, shipmentID string) (*domain.ShipmentDetails, error)
	GetAllWithStatuses(ctx context.Context, email string) ([]*domain.ShipmentWithCurrentStatus, error)
}

type ShipmentRepository interface {
	Save(ctx context.Context, shipment *domain.Shipment) error
	FindByID(ctx context.Context, shipmentID string) (*domain.Shipment, error)
	GetAllWithStatuses(ctx context.Context, userID string) ([]*domain.ShipmentWithCurrentStatus, error)
	GetAllByUserID(ctx context.Context, userID string) ([]*domain.Shipment, error)
}
