package shipmentservice

import (
	"context"
	"errors"

	"github.com/ixmael/99minutos/internal/core/domain"
)

// GetShipmentDetails retrieves the details of a shipment by its ID.
func (service *shipmentservice) GetShipmentDetails(ctx context.Context, email, shipmentID string) (*domain.ShipmentDetails, error) {
	user, err := service.userrepository.FindByEmail(ctx, email)
	if err != nil {
		service.logger.Error("failed to load user", "error", err)
		return nil, err
	}
	if user == nil {
		service.logger.Error("user not found", "email", email)
		return nil, errors.New("user not found")
	}

	shipment, err := service.shipmentrepository.FindByShipmentIDAndUser(ctx, user.ID, shipmentID, user.IsAdmin)
	if err != nil {
		service.logger.Error("failed to find shipment by id", "error", err)
		return nil, err
	}
	if shipment == nil {
		service.logger.Error("shipment not found", "shipmentID", shipmentID)
		return nil, errors.New("shipment not found")
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
