package ports

import (
	"context"

	"github.com/ixmael/99minutos/internal/core/domain"
)

// ShipmentService defines the interface for shipment-related operations.
type ShipmentService interface {
	CreateShipment(ctx context.Context, newShipmentRequest *domain.NewShipmentRequest) (*domain.ShipmentCreatedResult, error)
	GetShipmentDetails(ctx context.Context, email, shipmentID string) (*domain.ShipmentDetails, error)
	GetAllWithStatuses(ctx context.Context, email string) ([]*domain.ShipmentWithCurrentStatus, error)
}

type ShipmentRepository interface {
	Save(ctx context.Context, shipment *domain.Shipment) error
	FindByShipmentID(ctx context.Context, shipmentID string) (*domain.Shipment, error)
	FindByShipmentIDAndUser(ctx context.Context, userID int64, shipmentID string, userIsAdmin bool) (*domain.Shipment, error)
	GetAllWithStatuses(ctx context.Context) ([]*domain.ShipmentWithCurrentStatus, error)
	GetAllMyShipmentsWithStatuses(ctx context.Context, userID int64) ([]*domain.ShipmentWithCurrentStatus, error)
	GetAllByUserID(ctx context.Context, userID int64) ([]*domain.Shipment, error)
}
