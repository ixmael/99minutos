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
		service.logger.Error("invalid shipment")
		return nil, err
	}

	err = service.shipmentrepository.Save(ctx, newShipment)
	if err != nil {
		service.logger.Error("cannot save shipment")
		return nil, err
	}

	shipmentResult := domain.ShipmentCreatedResult{
		ID: newShipment.ID,
	}

	return &shipmentResult, nil
}
