package shipmentrepository

import (
	"context"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/ixmael/99minutos/internal/core/domain"
)

func (repo *postgresshipmentrepository) FindByShipmentIDAndUser(ctx context.Context, userID int64, shipmentID string, userIsAdmin bool) (*domain.Shipment, error) {
	getShipmentByIDQuery := sq.
		Select(
			"tracking_number_id",
			"origin",
			"destination",
			"created_at",
			"updated_at",
		).
		From("shipment").
		Limit(1)

	if userIsAdmin {
		getShipmentByIDQuery = getShipmentByIDQuery.Where(sq.Eq{"tracking_number_id": shipmentID})
	} else {
		getShipmentByIDQuery = getShipmentByIDQuery.Where(
			sq.And{
				sq.Eq{"tracking_number_id": shipmentID},
				sq.Eq{"user_id": userID},
			},
		)
	}

	rows, err := getShipmentByIDQuery.RunWith(repo.db).PlaceholderFormat(sq.Dollar).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shipment *domain.Shipment = nil
	for rows.Next() {
		var trackingNumberId string
		var origin string
		var destination string
		var createdAt time.Time
		var updatedAt time.Time

		rows.Scan(&trackingNumberId, &origin, &destination, &createdAt, &updatedAt)

		shipment = &domain.Shipment{
			ID:          trackingNumberId,
			Origin:      origin,
			Destination: destination,
			CreatedAt:   &createdAt,
			UpdatedAt:   &updatedAt,
		}
	}

	if shipment == nil {
		return nil, errors.New("shipment not found")
	}

	return shipment, nil
}
