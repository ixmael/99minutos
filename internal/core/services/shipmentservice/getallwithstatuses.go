package shipmentservice

import (
	"context"
	"errors"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (service *shipmentservice) GetAllWithStatuses(ctx context.Context, email string) ([]*domain.ShipmentWithCurrentStatus, error) {
	user, err := service.userrepository.FindByEmail(ctx, email)
	if err != nil {
		service.logger.Error("failed to load user", "error", err)
		return nil, err
	}
	if user == nil {
		service.logger.Error("user not found", "email", email)
		return nil, errors.New("user not found")
	}

	shipmentsWithStatus, err := service.shipmentrepository.GetAllWithStatuses(ctx, user.ID)
	if err != nil {
		service.logger.Error("failed to load all the shipments with the current statuses", "error", err)
		return nil, err
	}

	return shipmentsWithStatus, nil
}
