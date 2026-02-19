package userrepository

import (
	"context"
	"errors"
	"sync"

	"github.com/ixmael/99minutos/internal/core/domain"
	"github.com/ixmael/99minutos/internal/core/ports"
)

type InMemoryUserRepository struct {
	usersLock sync.Mutex
	users     []*domain.User
}

func NewInMemoryUserRepository() (ports.UserRepository, error) {
	repo := InMemoryUserRepository{
		users: make([]*domain.User, 0),
	}

	return &repo, nil
}

func (repo *InMemoryUserRepository) Register(ctx context.Context, user *domain.User) error {
	repo.usersLock.Lock()
	defer repo.usersLock.Unlock()

	var userExists *domain.User = nil
	for _, currentUser := range repo.users {
		if currentUser.Email == user.Email {
			userExists = currentUser
			break
		}
	}
	if userExists != nil {
		return errors.New("user already exists")
	}

	repo.users = append(repo.users, user)

	return nil
}

func (repo *InMemoryUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	repo.usersLock.Lock()
	defer repo.usersLock.Unlock()

	var id int64 = 0
	var user *domain.User = nil
	for _, currentUser := range repo.users {
		id = id + 1
		if currentUser.Email == email {
			user = currentUser
			break
		}
	}

	user.ID = id

	return user, nil
}

func (repo *InMemoryUserRepository) FindByEmailAndHashedPassword(ctx context.Context, email, hashedPassword string) (*domain.User, error) {
	return nil, errors.New("not implemented")
}
