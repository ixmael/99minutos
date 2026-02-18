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
	shimpmentLock    sync.Mutex
	shipments        map[string]*domain.Shipment
	statusRepository ports.ShipmentStatusRepository
}

func NewInMemoryShipmentRepository(statusRepo ports.ShipmentStatusRepository) (ports.ShipmentRepository, error) {
	repo := InMemoryShipmentRepository{
		shipments:        make(map[string]*domain.Shipment),
		statusRepository: statusRepo,
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

func (repo *InMemoryShipmentRepository) GetAllWithStatuses(ctx context.Context, userID string) ([]*domain.ShipmentWithCurrentStatus, error) {
	repo.shimpmentLock.Lock()
	defer repo.shimpmentLock.Unlock()

	var shipments []*domain.ShipmentWithCurrentStatus

	for _, shipment := range repo.shipments {
		if shipment.UserID != userID {
			continue
		}

		statuses, err := repo.statusRepository.FindStatusByID(ctx, shipment.ID)
		var currentStatus domain.ShipmentStatusList
		if err == nil && len(statuses) > 0 {
			currentStatus = statuses[len(statuses)-1].Status
		}

		shipments = append(shipments, &domain.ShipmentWithCurrentStatus{
			ID:          shipment.ID,
			Origin:      shipment.Origin,
			Destination: shipment.Destination,
			Status:      currentStatus,
			CreatedAt:   shipment.CreatedAt,
			UpdatedAt:   shipment.UpdatedAt,
		})
	}

	return shipments, nil
}

func (repo *InMemoryShipmentRepository) GetAllByUserID(ctx context.Context, userID string) ([]*domain.Shipment, error) {
	repo.shimpmentLock.Lock()
	defer repo.shimpmentLock.Unlock()

	var shipments []*domain.Shipment

	for _, shipment := range repo.shipments {
		if shipment.UserID != userID {
			continue
		}

		shipments = append(shipments, shipment)
	}

	return shipments, nil
}
