package shipmentstatusrepository

import (
	"context"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (repo *postgresshipmentstatusrepository) Save(ctx context.Context, shipmentStatus *domain.ShipmentStatus) error {
	insertShipmentStatusQuery := sq.
		Insert("shipment_status").
		Columns(
			"shipment_id",
			"status",
			"created_at",
		).
		Values(
			shipmentStatus.ID,
			shipmentStatus.Status,
			shipmentStatus.CreatedAt.Format(time.RFC3339),
		)

	result, err := insertShipmentStatusQuery.RunWith(repo.db).PlaceholderFormat(sq.Dollar).ExecContext(ctx)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return errors.New("failed to save shipment status")
	}

	return nil
}
