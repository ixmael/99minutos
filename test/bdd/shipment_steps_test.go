//go:build integration
// +build integration

package bdd

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cucumber/godog"

	"github.com/ixmael/99minutos/internal/core/domain"
	"github.com/ixmael/99minutos/internal/core/services/shipmentservice"
	"github.com/ixmael/99minutos/internal/infrastructure/inmemory/shipmentrepository"
)

// ShipmentSteps contains shared step definitions for shipment-related BDD tests
type ShipmentSteps struct {
	testContext *TestContext
}

// NewShipmentSteps creates a new ShipmentSteps instance
func NewShipmentSteps(tc *TestContext) *ShipmentSteps {
	return &ShipmentSteps{testContext: tc}
}

// RegisterShipmentSteps registers all shipment-related step definitions
func RegisterShipmentSteps(ctx *godog.ScenarioContext, tc *TestContext) {
	ss := NewShipmentSteps(tc)

	ctx.Step(`^the system is ready to accept shipment registration requests$`, ss.GivenShipmentSystemIsReady)
	ctx.Step(`^the origin is "([^"]*)"$`, ss.GivenShipmentOrigin)
	ctx.Step(`^the destination is "([^"]*)"$`, ss.GivenShipmentDestination)
	ctx.Step(`^a shipment request is submitted$`, ss.WhenRequestNewShipment)
	ctx.Step(`^the registration request is successful$`, ss.ThenRegistrationRequestIsSuccessful)
	ctx.Step(`^the shipment state is "([^"]*)"$`, ss.ThenShipmentStateIs)
}

func (ss *ShipmentSteps) GivenShipmentSystemIsReady() error {
	shipmentrepository, err := shipmentrepository.NewInMemoryShipmentRepository()
	if err != nil {
		return err
	}

	shipmentservice, err := shipmentservice.NewShipmentService(shipmentrepository)
	if err != nil {
		return err
	}

	ss.testContext.ShipmentRepository = shipmentrepository
	ss.testContext.ShipmentService = shipmentservice

	ss.testContext.LastError = nil
	ss.testContext.NewShipmentOriginRequest = nil
	ss.testContext.NewShipmentDestinationRequest = nil

	return nil
}

func (ss *ShipmentSteps) GivenShipmentOrigin(originStr string) error {
	ss.testContext.NewShipmentOriginRequest = &originStr

	return nil
}

func (ss *ShipmentSteps) GivenShipmentDestination(destinationStr string) error {
	ss.testContext.NewShipmentDestinationRequest = &destinationStr

	return nil
}

func (ss *ShipmentSteps) WhenRequestNewShipment() error {
	newShipmentRequest := &domain.NewShipmentRequest{
		Origin:      *ss.testContext.NewShipmentOriginRequest,
		Destination: *ss.testContext.NewShipmentDestinationRequest,
	}

	shipmentResult, err := ss.testContext.ShipmentService.CreateShipment(
		context.Background(),
		newShipmentRequest,
	)
	ss.testContext.LastShipmentCreatedResult = shipmentResult
	ss.testContext.LastError = err

	return nil
}

func (ss *ShipmentSteps) ThenRegistrationRequestIsSuccessful() error {
	if ss.testContext.LastError != nil {
		return errors.New(fmt.Sprintf("registration request failed: %v", ss.testContext.LastError))
	}

	if ss.testContext.LastShipmentCreatedResult == nil {
		return errors.New("shipment not created")
	}

	return nil
}

func (ss *ShipmentSteps) ThenShipmentStateIs(stateStr string) error {
	shipment, err := ss.testContext.ShipmentRepository.FindByID(
		context.Background(),
		ss.testContext.LastShipmentCreatedResult.ID,
	)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to find shipment: %v", err))
	}

	if strings.ToLower(string(shipment.Status)) != strings.ToLower(stateStr) {
		return errors.New(fmt.Sprintf("shipment state is not %s", stateStr))
	}

	return nil
}
