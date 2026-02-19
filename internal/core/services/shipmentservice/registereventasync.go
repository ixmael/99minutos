package shipmentservice

import (
	"context"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (service *shipmentservice) RegisterEventAsync(ctx context.Context, shipmentEventRequest *domain.ShipmentEventRequest) (*domain.EventStatusResult, error) {
	err := service.queueservice.Publish(ctx, domain.QueueShipmentKey, shipmentEventRequest)
	if err != nil {
		return nil, err
	}

	return &domain.EventStatusResult{Status: "success"}, nil
}
