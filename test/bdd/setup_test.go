//go:build integration
// +build integration

package bdd

import (
	"testing"

	"github.com/cucumber/godog"
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"../features"},
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

// InitializeScenario initializes the scenario context with all step definitions
func InitializeScenario(ctx *godog.ScenarioContext) {
	tc := NewTestContext()

	RegisterClientShipmentSteps(ctx, tc)
	RegisterUserSteps(ctx, tc)
	RegisterAdminSteps(ctx, tc)
}
