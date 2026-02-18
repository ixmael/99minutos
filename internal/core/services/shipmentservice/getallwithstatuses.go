package shipmentservice

import (
	"context"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (service *shipmentservice) GetAllWithStatuses(ctx context.Context) ([]*domain.ShipmentWithCurrentStatus, error) {
	shipmentsWithStatus, err := service.shipmentrepository.GetAllWithStatuses(ctx)
	if err != nil {
		service.logger.Error("failed to load all the shipments with the current statuses", "error", err)
		return nil, err
	}

	return shipmentsWithStatus, nil
}
