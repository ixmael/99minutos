package userrepository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (repo *postgresuserrepository) Register(ctx context.Context, user *domain.User) error {
	insertUserQuery := sq.
		Insert("users").
		Columns(
			"email",
			"hashed_password",
			"is_admin",
			"created_at",
		).
		Values(
			user.Email,
			user.HashedPassword,
			user.IsAdmin,
			user.CreatedAt,
		).
		Suffix("RETURNING id")

	err := insertUserQuery.RunWith(repo.db).PlaceholderFormat(sq.Dollar).QueryRowContext(ctx).Scan(&user.ID)
	if err != nil {
		repo.logger.Error("cannot create user", "error", err)
		return err
	}

	return nil
}
