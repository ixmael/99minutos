package shipmentservice

import (
	"context"
	"encoding/json"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (service *shipmentservice) RegisterEventsBatchAsync(ctx context.Context, idempotencyKey string, shipmentEventRequests []*domain.ShipmentEventRequest) (*domain.EventStatusResult, error) {
	currentStatusRaw, err := service.cacheservice.Get(ctx, idempotencyKey)
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
		for _, shipmentEventRequest := range shipmentEventRequests {
			err := service.queueservice.Publish(ctx, domain.QueueShipmentKey, shipmentEventRequest)
			if err != nil {
				return nil, err
			}

			pendingStatus := domain.EventStatusResult{
				Status: "pending",
			}

			err = service.cacheservice.Set(ctx, shipmentEventRequest.IdempotencyKey, &pendingStatus)
			if err != nil {
				return nil, err
			}

			requestStatus = &pendingStatus
		}
	}

	return requestStatus, nil
}
