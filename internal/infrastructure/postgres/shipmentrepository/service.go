package shipmentrepository

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/ixmael/99minutos/internal/core/ports"
)

type postgresshipmentrepository struct {
	db     *sql.DB
	logger ports.Logger
}

// NewPostgresShipmentRepository
func NewPostgresShipmentRepository(uri string, logger ports.Logger) (ports.ShipmentRepository, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	repo := postgresshipmentrepository{
		db:     db,
		logger: logger,
	}

	return &repo, nil
}
