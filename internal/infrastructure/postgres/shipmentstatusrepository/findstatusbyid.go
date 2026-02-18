package shipmentstatusrepository

import (
	"context"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/ixmael/99minutos/internal/core/domain"
)

func (repo *postgresshipmentstatusrepository) FindStatusByID(ctx context.Context, shipmentID string) ([]*domain.ShipmentStatus, error) {
	getShipmentStatusByIDQuery := sq.
		Select(
			"shipment_id",
			"status",
			"created_at",
		).
		From("shipment_status").
		Where(sq.Eq{"shipment_id": shipmentID}).
		OrderBy("created_at DESC")

	rows, err := getShipmentStatusByIDQuery.RunWith(repo.db).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	shipmentStatusList := []*domain.ShipmentStatus{}
	for rows.Next() {
		var shipment_id string
		var statusStr string
		var createdAt time.Time

		rows.Scan(&shipment_id, &statusStr, &createdAt)

		status := domain.ShipmentStatusList(strings.ToUpper(statusStr))
		shipmentStatus := domain.ShipmentStatus{
			ID:        shipment_id,
			Status:    status,
			CreatedAt: &createdAt,
		}

		shipmentStatusList = append(shipmentStatusList, &shipmentStatus)
	}

	return shipmentStatusList, nil
}
