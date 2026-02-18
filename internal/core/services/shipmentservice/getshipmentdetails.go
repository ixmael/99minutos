package shipmentservice

import (
	"context"

	"github.com/ixmael/99minutos/internal/core/domain"
)

// GetShipmentDetails retrieves the details of a shipment by its ID.
func (service *shipmentservice) GetShipmentDetails(ctx context.Context, shipmentID string) (*domain.ShipmentDetails, error) {
	shipment, err := service.shipmentrepository.FindByID(ctx, shipmentID)
	if err != nil {
		service.logger.Error("failed to find shipment by id", "error", err)
		return nil, err
	}

	statuses, err := service.shipmentstatusrepository.FindStatusByID(ctx, shipment.ID)
	if err != nil {
		service.logger.Error("failed to find shipment status by id", "error", err)
		return nil, err
	}

	return &domain.ShipmentDetails{
		Shipment: *shipment,
		Statuses: statuses,
	}, nil
}
