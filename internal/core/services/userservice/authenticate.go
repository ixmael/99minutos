package userservice

import (
	"context"
	"errors"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (service *userserviceimpl) Authenticate(ctx context.Context, email string, password string) (*domain.User, error) {
	hashedPassword, err := domain.HashPassword(password)
	if err != nil {
		service.logger.Error("cannot hash password", "error", err)
		return nil, err
	}

	user, err := service.repo.FindByEmailAndHashedPassword(ctx, email, hashedPassword)
	if err != nil {
		service.logger.Error("cannot find user", "error", err)
		return nil, err
	}
	if user == nil {
		service.logger.Error("there isn't a user with that email")
		return nil, errors.New("there isn't a user with that email")
	}

	return user, nil
}
