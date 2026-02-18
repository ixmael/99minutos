//go:build integration
// +build integration

package bdd

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/ixmael/99minutos/internal/core/domain"
)

// UserSteps contains shared step definitions for user-related BDD tests
type UserSteps struct {
	testContext *TestContext
}

// NewShipmentSteps creates a new ShipmentSteps instance
func NewUserSteps(tc *TestContext) *UserSteps {
	return &UserSteps{testContext: tc}
}

// RegisterUserSteps registers all user-related step definitions
func RegisterUserSteps(ctx *godog.ScenarioContext, tc *TestContext) {
	us := NewUserSteps(tc)

	ctx.Step(`^the client is registered with the email "([^"]*)"$`, us.GivenClientIsRegistered)
}

func (us *UserSteps) GivenClientIsRegistered(email string) error {
	client, err := domain.NewClient(email, "thePazsW0r?")
	if err != nil {
		return err
	}

	err = us.testContext.UserRepository.Register(context.Background(), client)
	if err != nil {
		return err
	}

	us.testContext.UserEmail = &client.Email

	return nil
}
