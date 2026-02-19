package ports

import (
	"context"

	"github.com/ixmael/99minutos/internal/core/domain"
)

// UserService defines the interface for user-related operations.
type UserService interface {
	CreateUser(ctx context.Context, newUserRequest *domain.UserRequest) (*domain.UserRegisteredResult, error)
	Authenticate(ctx context.Context, email string, password string) (*domain.User, error)
}

// UserRepository defines the interface for user-related operations repository.
type UserRepository interface {
	Register(ctx context.Context, user *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByEmailAndHashedPassword(ctx context.Context, email, hashedPassword string) (*domain.User, error)
	Stop()
}
