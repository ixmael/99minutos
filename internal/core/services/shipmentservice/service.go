package shipmentservice

import (
	"github.com/ixmael/99minutos/internal/core/domain"
	"github.com/ixmael/99minutos/internal/core/ports"
)

// shipmentservice implements the ShipmentService interface.
type shipmentservice struct {
	logger                   ports.Logger
	shipmentrepository       ports.ShipmentRepository
	shipmentstatusrepository ports.ShipmentStatusRepository
	userrepository           ports.UserRepository
	queueservice             ports.QueueService
	cacheservice             ports.CacheService
}

// NewShipmentService creates a new instance of the shipment service.
func NewShipmentService(
	logger ports.Logger,
	shipmentrepository ports.ShipmentRepository,
	shipmentstatusrepository ports.ShipmentStatusRepository,
	userrepository ports.UserRepository,
	queueservice ports.QueueService,
	cacheservice ports.CacheService,
) (ports.ShipmentService, error) {
	err := queueservice.RegisterQueue(domain.QueueShipmentKey)
	if err != nil {
		return nil, err
	}

	service := shipmentservice{
		logger:                   logger,
		shipmentrepository:       shipmentrepository,
		shipmentstatusrepository: shipmentstatusrepository,
		userrepository:           userrepository,
		queueservice:             queueservice,
		cacheservice:             cacheservice,
	}

	return &service, nil
}
