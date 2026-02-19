package shipmentrepository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/ixmael/99minutos/internal/core/domain"
)

func (repo *postgresshipmentrepository) FindByShipmentID(ctx context.Context, shipmentID string) (*domain.Shipment, error) {
	getShipmentByIDQuery := sq.
		Select(
			"tracking_number_id",
			"origin",
			"destination",
			"created_at",
			"updated_at",
		).
		From("shipment").
		Where(sq.Eq{"tracking_number_id": shipmentID}).
		Limit(1)

	var shipment domain.Shipment

	err := getShipmentByIDQuery.RunWith(repo.db).PlaceholderFormat(sq.Dollar).QueryRowContext(ctx).Scan(
		&shipment.ID,
		&shipment.Origin,
		&shipment.Destination,
		&shipment.CreatedAt,
		&shipment.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &shipment, nil
}
