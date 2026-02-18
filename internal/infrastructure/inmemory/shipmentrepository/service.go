package shipmentrepository

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"

	"github.com/ixmael/99minutos/internal/core/domain"
	"github.com/ixmael/99minutos/internal/core/ports"
)

type InMemoryShipmentRepository struct {
	shimpmentLock sync.Mutex
	shipments     map[string]*domain.Shipment
}

func NewInMemoryShipmentRepository() (ports.ShipmentRepository, error) {
	repo := InMemoryShipmentRepository{
		shipments: make(map[string]*domain.Shipment),
	}

	return &repo, nil
}

func (repo *InMemoryShipmentRepository) Save(ctx context.Context, shipment *domain.Shipment) error {
	repo.shimpmentLock.Lock()
	defer repo.shimpmentLock.Unlock()

	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	shipment.ID = id.String()

	repo.shipments[shipment.ID] = shipment

	return nil
}

func (repo *InMemoryShipmentRepository) FindByID(ctx context.Context, shipmentID string) (*domain.Shipment, error) {
	repo.shimpmentLock.Lock()
	defer repo.shimpmentLock.Unlock()

	shipment, ok := repo.shipments[shipmentID]
	if !ok {
		return nil, errors.New("shipment not found")
	}

	return shipment, nil
}

func (repo *InMemoryShipmentRepository) GetAll() []*domain.Shipment {
	repo.shimpmentLock.Lock()
	defer repo.shimpmentLock.Unlock()

	var shipments []*domain.Shipment
	for _, shipment := range repo.shipments {
		shipments = append(shipments, shipment)
	}

	return shipments
}

func (repo *InMemoryShipmentRepository) GetAllWithStatuses(ctx context.Context) ([]*domain.ShipmentWithCurrentStatus, error) {
	repo.shimpmentLock.Lock()
	defer repo.shimpmentLock.Unlock()

	var shipments []*domain.ShipmentWithCurrentStatus

	return shipments, nil
}
