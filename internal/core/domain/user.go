package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system.
type User struct {
	ID             string     `json:"id"`
	Email          string     `json:"email"`
	HashedPassword string     `json:"hashed_password"`
	IsAdmin        bool       `json:"is_admin"`
	CreatedAt      *time.Time `json:"created_at"`
}

// UserRequest represents a user request in the system.
type UserRequest struct {
	Email         string `json:"email"`
	PlainPassword string `json:"password"`
}

// UserRegisteredResult represents the result of a user registration.
type UserRegisteredResult struct {
	ID string `json:"id"`
}

func NewClient(email, plainPassword string) (*User, error) {
	hashedPassword, err := hashedPassword(plainPassword)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	u := User{
		ID:             id.String(),
		Email:          email,
		HashedPassword: hashedPassword,
		IsAdmin:        false,
		CreatedAt:      &now,
	}

	return &u, nil
}

func hashedPassword(plainPassword string) (string, error) {
	return fmt.Sprintf("hashed_%s", plainPassword), nil
}
