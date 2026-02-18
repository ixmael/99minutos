package shipmentservice

import (
	"github.com/ixmael/99minutos/internal/core/ports"
)

// shipmentservice implements the ShipmentService interface.
type shipmentservice struct {
	logger             ports.Logger
	shipmentrepository ports.ShipmentRepository
}

// NewShipmentService creates a new instance of the shipment service.
func NewShipmentService(logger ports.Logger, shipmentrepository ports.ShipmentRepository) (ports.ShipmentService, error) {
	service := shipmentservice{
		logger:             logger,
		shipmentrepository: shipmentrepository,
	}

	return &service, nil
}
