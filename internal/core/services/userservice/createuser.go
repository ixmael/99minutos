package userservice

import (
	"context"
	"errors"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (service *userserviceimpl) CreateUser(ctx context.Context, newUserRequest *domain.UserRequest) (*domain.UserRegisteredResult, error) {
	existsUser, err := service.repo.FindByEmail(ctx, newUserRequest.Email)
	if err != nil {
		service.logger.Error("cannot find user", "error", err)
		return nil, err
	}
	if existsUser != nil {
		service.logger.Error("there is already a user with that email")
		return nil, errors.New("there is already a user with that email")
	}

	user, err := domain.NewUser(newUserRequest.Email, newUserRequest.PlainPassword, newUserRequest.IsAdmin)
	if err != nil {
		service.logger.Error("cannot create user", "error", err)
		return nil, err
	}

	err = service.repo.Register(ctx, user)
	if err != nil {
		service.logger.Error("cannot save user", "error", err)
		return nil, err
	}

	userRegisteredResult := domain.UserRegisteredResult{
		ID: user.ID,
	}

	return &userRegisteredResult, nil
}
