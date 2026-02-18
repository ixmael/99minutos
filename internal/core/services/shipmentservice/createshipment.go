package shipmentservice

import (
	"context"
	"errors"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (service *shipmentservice) CreateShipment(
	ctx context.Context,
	newShipmentRequest *domain.NewShipmentRequest,
) (*domain.ShipmentCreatedResult, error) {
	if newShipmentRequest == nil {
		service.logger.Error("invalid request")
		return nil, errors.New("invalid request")
	}

	newShipment, err := domain.NewShipment(ctx, newShipmentRequest.Origin, newShipmentRequest.Destination)
	if err != nil {
		service.logger.Error("invalid shipment", "error", err)
		return nil, err
	}

	err = service.shipmentrepository.Save(ctx, newShipment)
	if err != nil {
		service.logger.Error("cannot save shipment", "error", err)
		return nil, err
	}

	shipmentStatus, err := domain.NewCreatedShipmentStatus(newShipment.ID)
	if err != nil {
		service.logger.Error("cannot create shipment status", "error", err)
		return nil, err
	}

	err = service.shipmentstatusrepository.Save(ctx, shipmentStatus)
	if err != nil {
		service.logger.Error("cannot save shipment status", "error", err)
		return nil, err
	}

	shipmentResult := domain.ShipmentCreatedResult{
		ID: newShipment.ID,
	}

	return &shipmentResult, nil
}
