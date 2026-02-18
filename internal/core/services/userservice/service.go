package userservice

import "github.com/ixmael/99minutos/internal/core/ports"

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
