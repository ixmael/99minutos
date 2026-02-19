package shipmentrepository

import (
	"context"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (repo *postgresshipmentrepository) Save(ctx context.Context, shipment *domain.Shipment) error {
	insertShipmentQuery := sq.
		Insert("shipment").
		Columns(
			"user_id",
			"tracking_number_id",
			"origin",
			"destination",
			"created_at",
		).
		Values(
			shipment.UserID,
			shipment.ID,
			shipment.Origin,
			shipment.Destination,
			shipment.CreatedAt.Format(time.RFC3339),
		)

	result, err := insertShipmentQuery.RunWith(repo.db).PlaceholderFormat(sq.Dollar).ExecContext(ctx)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return errors.New("failed to save shipment")
	}

	return nil
}
