package workers

import (
	"context"

	"github.com/ixmael/99minutos/internal/core/ports"
)

type ShipmentEventWorker struct {
	service ports.ShipmentService
	logger  ports.Logger
	queue   ports.QueueService
}

func NewShipmentEventWorker(
	logger ports.Logger,
	service ports.ShipmentService,
) (*ShipmentEventWorker, error) {
	worker := ShipmentEventWorker{
		logger:  logger,
		service: service,
	}

	return &worker, nil
}

func (w *ShipmentEventWorker) Start(ctx context.Context) {
	err := w.service.ProcessEventsAsync(ctx)
	if err != nil {
		w.logger.Error("failed to process events", "error", err)
		return
	}
}
