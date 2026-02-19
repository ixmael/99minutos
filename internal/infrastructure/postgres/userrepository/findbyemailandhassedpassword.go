package userrepository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (repo *postgresuserrepository) FindByEmailAndHashedPassword(ctx context.Context, email string, hashedPassword string) (*domain.User, error) {
	getUserByEmailAndHashedPasswordQuery := sq.
		Select(
			"id",
			"email",
			"is_admin",
			"created_at",
		).
		From("users").
		Where(sq.And{
			sq.Eq{"email": email},
			sq.Eq{"hashed_password": hashedPassword},
		}).
		Limit(1)

	var user domain.User

	err := getUserByEmailAndHashedPasswordQuery.RunWith(repo.db).PlaceholderFormat(sq.Dollar).QueryRowContext(ctx).Scan(
		&user.ID,
		&user.Email,
		&user.IsAdmin,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
