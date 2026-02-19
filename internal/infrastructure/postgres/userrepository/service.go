package userrepository

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/ixmael/99minutos/internal/core/ports"
)

type postgresuserrepository struct {
	db     *sql.DB
	logger ports.Logger
}

// NewPostgresUserStatusRepository implements ports.UserRepository
func NewPostgresUserStatusRepository(uri string, logger ports.Logger) (ports.UserRepository, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	repo := postgresuserrepository{
		db:     db,
		logger: logger,
	}

	return &repo, nil
}

func (repo *postgresuserrepository) Stop() {
	if repo.db != nil {
		err := repo.db.Close()
		if err != nil {
			repo.logger.Error("error closing database", err)
		}
	}
}
