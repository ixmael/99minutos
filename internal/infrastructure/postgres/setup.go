package postgres

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/ixmael/99minutos/internal/core/ports"
)

func SetupPostgresRepository(uri string, logger ports.Logger) error {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./internal/infrastructure/postgres/migrations",
		"postgres", driver)
	if err != nil {
		logger.Error("error on setup the migrations", "error", err)
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.Error("error on applying the migrations", "error", err)
		return err
	}

	return nil
}
