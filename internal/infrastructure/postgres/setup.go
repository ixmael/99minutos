package postgres

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/ixmael/99minutos/internal/core/ports"
)

func SetupPostgresRepository(logger ports.Logger, uri, pathMigrations string) error {
	return nil
}
