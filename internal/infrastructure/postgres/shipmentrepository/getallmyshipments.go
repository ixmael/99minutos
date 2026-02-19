package shipmentrepository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"github.com/ixmael/99minutos/internal/core/domain"
)

func (repo *postgresshipmentrepository) GetAllMyShipmentsWithStatuses(ctx context.Context, userID int64, pagination *domain.PaginationRequest) ([]*domain.ShipmentWithCurrentStatus, error) {
	statusSubquery := sq.Select("DISTINCT ON (shipment_id) shipment_id, status, created_at").
		From("shipment_status").
		OrderBy("shipment_id", "created_at ASC")
	subQuerySql, subArgs, err := statusSubquery.ToSql()
	if err != nil {
		return nil, err
	}

	offset := uint64((pagination.Page - 1) * pagination.Limit)
	getMyShipmentWithStatusesQuery := sq.Select(
		"S.tracking_number_id AS id",
		"S.origin AS origin",
		"S.destination AS destination",
		"S.created_at AS created_at",
		"ST.status AS status",
		"ST.created_at AS updated_at",
	).From("shipment S").
		Join(fmt.Sprintf("(%s) st ON S.tracking_number_id = ST.shipment_id", subQuerySql), subArgs...).
		Where(sq.Eq{"user_id": userID}).
		Limit(uint64(pagination.Limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := getMyShipmentWithStatusesQuery.RunWith(repo.db).PlaceholderFormat(sq.Dollar).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	shipmentsWithStatuses := []*domain.ShipmentWithCurrentStatus{}
	for rows.Next() {
		var id string
		var origin string
		var destination string
		var statusStr string
		var createdAt time.Time
		var updatedAt time.Time

		rows.Scan(
			&id,
			&origin,
			&destination,
			&createdAt,
			&statusStr,
			&updatedAt,
		)

		status := domain.ShipmentStatusList(strings.ToUpper(statusStr))
		shipmentWithCurrentStatus := &domain.ShipmentWithCurrentStatus{
			ID:          id,
			Origin:      origin,
			Destination: destination,
			Status:      status,
			CreatedAt:   &createdAt,
			UpdatedAt:   &updatedAt,
		}

		shipmentsWithStatuses = append(shipmentsWithStatuses, shipmentWithCurrentStatus)
	}

	return shipmentsWithStatuses, nil
}
