package shipmentservice

import "github.com/ixmael/99minutos/internal/core/ports"

// shipmentservice implements the ShipmentService interface.
type shipmentservice struct {
	shipmentrepository ports.ShipmentRepository
}

// NewShipmentService creates a new instance of the shipment service.
func NewShipmentService(shipmentrepository ports.ShipmentRepository) (ports.ShipmentService, error) {
	service := shipmentservice{
		shipmentrepository: shipmentrepository,
	}

	return &service, nil
}
