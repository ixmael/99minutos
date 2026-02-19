package userrepository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (repo *postgresuserrepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	getUserByEmailQuery := sq.
		Select(
			"id",
			"email",
			"is_admin",
			"created_at",
		).
		From("users").
		Where(sq.Eq{"email": email}).
		Limit(1)

	rows, err := getUserByEmailQuery.RunWith(repo.db).PlaceholderFormat(sq.Dollar).QueryContext(ctx)
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
