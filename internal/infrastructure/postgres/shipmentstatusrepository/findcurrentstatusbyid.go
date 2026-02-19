package shipmentstatusrepository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/ixmael/99minutos/internal/core/domain"
)

func (repo *postgresshipmentstatusrepository) FindCurrentStatusByID(ctx context.Context, shipmentID string) (*domain.ShipmentStatus, error) {
	getShipmentCurrentStatusByIDQuery := sq.
		Select(
			"shipment_id",
			"status",
			"created_at",
		).
		From("shipment_status").
		Where(sq.Eq{"shipment_id": shipmentID}).
		OrderBy("created_at DESC").
		Limit(1)

	var shipmentStatus domain.ShipmentStatus

	err := getShipmentCurrentStatusByIDQuery.RunWith(repo.db).PlaceholderFormat(sq.Dollar).QueryRowContext(ctx).Scan(
		&shipmentStatus.ID,
		&shipmentStatus.Status,
		&shipmentStatus.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &shipmentStatus, nil
}
