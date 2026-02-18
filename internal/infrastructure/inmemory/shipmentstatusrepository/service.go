package shipmentstatusrepository

import (
	"context"
	"sync"

	"github.com/ixmael/99minutos/internal/core/domain"
	"github.com/ixmael/99minutos/internal/core/ports"
)

type InMemoryShipmentStatusRepository struct {
	statusLock sync.Mutex
	status     []*domain.ShipmentStatus
}

func NewInMemoryShipmentStatusRepository() (ports.ShipmentStatusRepository, error) {
	repo := InMemoryShipmentStatusRepository{
		status: make([]*domain.ShipmentStatus, 0),
	}

	return &repo, nil
}

func (repo *InMemoryShipmentStatusRepository) Save(ctx context.Context, shipmentStatus *domain.ShipmentStatus) error {
	repo.statusLock.Lock()
	defer repo.statusLock.Unlock()

	repo.status = append(repo.status, shipmentStatus)

	return nil
}

func (repo *InMemoryShipmentStatusRepository) FindStatusByID(ctx context.Context, shipmentID string) ([]*domain.ShipmentStatus, error) {
	repo.statusLock.Lock()
	defer repo.statusLock.Unlock()

	status := []*domain.ShipmentStatus{}
	for _, currentShipmentStatus := range repo.status {
		if currentShipmentStatus.ID == shipmentID {
			status = append(status, currentShipmentStatus)
		}
	}

	return status, nil
}
