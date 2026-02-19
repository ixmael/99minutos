//go:build integration
// +build integration

package bdd

import (
	"errors"

	"github.com/cucumber/godog"
)

// AdminSteps contains shared step definitions for admin-related BDD tests
type AdminSteps struct {
	testContext *TestContext
}

// NewAdminSteps creates a new AdminSteps instance
func NewAdminSteps(tc *TestContext) *AdminSteps {
	return &AdminSteps{testContext: tc}
}

// RegisterUserSteps registers all user-related step definitions
func RegisterAdminSteps(ctx *godog.ScenarioContext, tc *TestContext) {
	as := NewAdminSteps(tc)

	ctx.Step(`^I should receive a list with 2 shipments$`, as.ThenShouldReceiveListWithTwoShipments)
}

func (as *AdminSteps) ThenShouldReceiveListWithTwoShipments() error {
	if len(as.testContext.LastShipmentsWithStatuses) != 2 {
		return errors.New("expected two shipments, got none")
	}

	return nil
}
