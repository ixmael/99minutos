package userrepository

import (
	"context"
	"time"

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

	rows, err := getUserByEmailAndHashedPasswordQuery.RunWith(repo.db).PlaceholderFormat(sq.Dollar).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user *domain.User = nil
	for rows.Next() {
		var id int64
		var emailRow string
		var isAdmin bool
		var createdAt time.Time

		rows.Scan(&id, &emailRow, &isAdmin, &createdAt)

		user = &domain.User{
			ID:        id,
			Email:     emailRow,
			IsAdmin:   isAdmin,
			CreatedAt: &createdAt,
		}
	}

	return user, nil
}
