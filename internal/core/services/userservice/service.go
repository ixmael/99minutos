package userservice

import (
	"github.com/ixmael/99minutos/internal/core/ports"
)

type userserviceimpl struct {
	repo   ports.UserRepository
	logger ports.Logger
}

func NewUserService(logger ports.Logger, repo ports.UserRepository) (ports.UserService, error) {
	service := userserviceimpl{
		repo:   repo,
		logger: logger,
	}

	return &service, nil
}
