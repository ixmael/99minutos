package shipmentservice

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (service *shipmentservice) ProcessEventsAsync(ctx context.Context) error {
	return service.queueservice.Consume(ctx, domain.QueueShipmentKey, func(rawData []byte) error {
		var shipmentEventRequest domain.ShipmentEventRequest
		err := json.Unmarshal(rawData, &shipmentEventRequest)
		if err != nil {
			service.logger.Error("error on unmarshaling shipment event request", "error", err)
			return err
		}

		shipment, err := service.shipmentrepository.FindByShipmentID(ctx, shipmentEventRequest.TrackingNumber)
		if err != nil {
			service.logger.Error("error on finding shipment by id", "error", err)
			return err
		}

		currentShipmentStatus, err := service.shipmentstatusrepository.FindCurrentStatusByID(ctx, shipment.ID)
		if err != nil {
			service.logger.Error("error on finding shipment status by id", "error", err)
			return err
		}

		nextStatus := domain.ShipmentStatusList(strings.ToUpper(shipmentEventRequest.Status))
		if !domain.IsValidTransition(currentShipmentStatus.Status, nextStatus) {
			service.logger.Error("the new status is not valid")
			return nil
		}

		now := time.Now()
		newShipmentStatus := domain.ShipmentStatus{
			ID:        shipment.ID,
			Status:    nextStatus,
			CreatedAt: &now,
		}

		err = service.shipmentstatusrepository.Save(ctx, &newShipmentStatus)
		if err != nil {
			service.logger.Error("error on creating the new shipment status", "error", err)
			return err
		}

		return nil
	})
}
