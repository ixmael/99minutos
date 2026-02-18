package ports

import (
	"context"

	"github.com/ixmael/99minutos/internal/core/domain"
)

// UserService defines the interface for user-related operations.
type UserService interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	// CreateUser(ctx context.Context, newUserRequest *domain.UserRequest) (*domain.UserRegisteredResult, error)
}

// UserRepository defines the interface for user-related operations repository.
type UserRepository interface {
	Register(ctx context.Context, user *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}
