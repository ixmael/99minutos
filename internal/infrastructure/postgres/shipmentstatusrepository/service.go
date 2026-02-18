package shipmentstatusrepository

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/ixmael/99minutos/internal/core/ports"
)

type postgresshipmentstatusrepository struct {
	db *sql.DB
}

// NewPostgresShipmentStatusRepository implements ports.ShipmentStatusRepository
func NewPostgresShipmentStatusRepository(uri string) (ports.ShipmentStatusRepository, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	repo := postgresshipmentstatusrepository{
		db: db,
	}

	return &repo, nil
}
