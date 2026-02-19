package domain

import (
	"fmt"
	"time"
)

const (
	ClientRole = "client"
	AdminRole  = "admin"
)

// User represents a user in the system.
type User struct {
	ID             int64      `json:"id"`
	Email          string     `json:"email"`
	HashedPassword string     `json:"hashed_password"`
	IsAdmin        bool       `json:"is_admin"`
	CreatedAt      *time.Time `json:"created_at"`
}

// UserRequest represents a user request in the system.
type UserRequest struct {
	Email         string `json:"email"`
	PlainPassword string `json:"password"`
	IsAdmin       bool   `json:"is_admin"`
}

// UserRegisteredResult represents the result of a user registration.
type UserRegisteredResult struct {
	ID int64 `json:"id"`
}

func NewClient(email, plainPassword string) (*User, error) {
	hashedPassword, err := hashedPassword(plainPassword)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	u := User{
		Email:          email,
		HashedPassword: hashedPassword,
		IsAdmin:        false,
		CreatedAt:      &now,
	}

	return &u, nil
}

func NewUser(email, plainPassword string, isAdmin bool) (*User, error) {
	hashedPassword, err := hashedPassword(plainPassword)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	u := User{
		Email:          email,
		HashedPassword: hashedPassword,
		IsAdmin:        isAdmin,
		CreatedAt:      &now,
	}

	return &u, nil
}

func HashPassword(plainPassword string) (string, error) {
	return hashedPassword(plainPassword)
}

func hashedPassword(plainPassword string) (string, error) {
	return fmt.Sprintf("hashed_%s", plainPassword), nil
}
