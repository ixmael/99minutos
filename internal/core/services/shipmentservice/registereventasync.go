package shipmentservice

import (
	"context"
	"encoding/json"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (service *shipmentservice) RegisterEventAsync(ctx context.Context, shipmentEventRequest *domain.ShipmentEventRequest) (*domain.EventStatusResult, error) {
	currentStatusRaw, err := service.cacheservice.Get(ctx, shipmentEventRequest.IdempotencyKey)
	if err != nil {
		return nil, err
	}

	var requestStatus *domain.EventStatusResult
	if currentStatusRaw != "" {
		var currentStatus domain.EventStatusResult

		err := json.Unmarshal([]byte(currentStatusRaw.(string)), &currentStatus)
		if err != nil {
			return nil, err
		}

		requestStatus = &currentStatus
	} else {
		err := service.queueservice.Publish(ctx, domain.QueueShipmentKey, shipmentEventRequest)
		if err != nil {
			return nil, err
		}

		pendingStatus := domain.EventStatusResult{
			TrackingNumber: shipmentEventRequest.TrackingNumber,
			Status:         "pending",
		}

		err = service.cacheservice.Set(ctx, shipmentEventRequest.IdempotencyKey, &pendingStatus)
		if err != nil {
			return nil, err
		}

		requestStatus = &pendingStatus
	}

	return requestStatus, nil
}
