package shipmentservice

import (
	"github.com/ixmael/99minutos/internal/core/ports"
)

// shipmentservice implements the ShipmentService interface.
type shipmentservice struct {
	logger                   ports.Logger
	shipmentrepository       ports.ShipmentRepository
	shipmentstatusrepository ports.ShipmentStatusRepository
	userrepository           ports.UserRepository
}

// NewShipmentService creates a new instance of the shipment service.
func NewShipmentService(
	logger ports.Logger,
	shipmentrepository ports.ShipmentRepository,
	shipmentstatusrepository ports.ShipmentStatusRepository,
	userrepository ports.UserRepository,
) (ports.ShipmentService, error) {
	service := shipmentservice{
		logger:                   logger,
		shipmentrepository:       shipmentrepository,
		shipmentstatusrepository: shipmentstatusrepository,
		userrepository:           userrepository,
	}

	return &service, nil
}
