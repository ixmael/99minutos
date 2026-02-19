package shipmentrepository

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"github.com/ixmael/99minutos/internal/core/domain"
)

func (repo *postgresshipmentrepository) GetAllByUserID(ctx context.Context, userID int64) ([]*domain.Shipment, error) {
	getShipmentQuery := sq.Select(
		"tracking_number_id",
		"origin",
		"destination",
		"created_at",
	).From("shipment").
		Where(sq.Eq{"user_id": userID}).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := getShipmentQuery.RunWith(repo.db).PlaceholderFormat(sq.Dollar).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	shipments := []*domain.Shipment{}
	for rows.Next() {
		var id string
		var origin string
		var destination string
		var createdAt time.Time

		rows.Scan(
			&id,
			&origin,
			&destination,
			&createdAt,
		)
		shipment := &domain.Shipment{
			ID:          id,
			Origin:      origin,
			Destination: destination,
			CreatedAt:   &createdAt,
		}

		shipments = append(shipments, shipment)
	}

	return shipments, nil
}
