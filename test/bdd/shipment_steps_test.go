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
	"github.com/ixmael/99minutos/internal/infrastructure/inmemory/logger"
	"github.com/ixmael/99minutos/internal/infrastructure/inmemory/shipmentrepository"
	"github.com/ixmael/99minutos/internal/infrastructure/inmemory/shipmentstatusrepository"
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
	ctx.Step(`^the registration request is rejected$`, ss.ThenRegistrationRequestIsRejected)
	ctx.Step(`^the repository not has any shipment$`, ss.ThenRepositoryNotHasAnyShipment)
	ctx.Step(`^the repository not has any shipment status$`, ss.ThenRepositoryNotHasAnyShipmentStatus)
}

func (ss *ShipmentSteps) GivenShipmentSystemIsReady() error {
	logger, err := logger.NewInMemoryLoggerServices()
	if err != nil {
		return err
	}

	shipmentrepository, err := shipmentrepository.NewInMemoryShipmentRepository()
	if err != nil {
		return err
	}

	shipmentstatusrepository, err := shipmentstatusrepository.NewInMemoryShipmentStatusRepository()
	if err != nil {
		return err
	}

	shipmentservice, err := shipmentservice.NewShipmentService(logger, shipmentrepository, shipmentstatusrepository)
	if err != nil {
		return err
	}

	ss.testContext.ShipmentStatusRepository = shipmentstatusrepository
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
	origin := ""
	if ss.testContext.NewShipmentDestinationRequest != nil {
		origin = *ss.testContext.NewShipmentDestinationRequest
	}
	destination := ""
	if ss.testContext.NewShipmentDestinationRequest != nil {
		destination = *ss.testContext.NewShipmentDestinationRequest
	}

	newShipmentRequest := &domain.NewShipmentRequest{
		Origin:      origin,
		Destination: destination,
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

	status, err := ss.testContext.ShipmentStatusRepository.FindStatusByID(
		context.Background(),
		shipment.ID,
	)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to find shipment status: %v", err))
	}

	if len(status) != 1 {
		return errors.New("shipment status not found")
	}

	if status[0].Status != domain.ShipmentStatusList(strings.ToUpper(stateStr)) {
		return errors.New(fmt.Sprintf("shipment status is not %s", stateStr))
	}

	return nil
}

func (ss *ShipmentSteps) ThenRegistrationRequestIsRejected() error {
	if ss.testContext.LastError == nil {
		return errors.New("registration request not rejected")
	}

	return nil
}

func (ss *ShipmentSteps) ThenRepositoryNotHasAnyShipment() error {
	repo, ok := ss.testContext.ShipmentRepository.(*shipmentrepository.InMemoryShipmentRepository)
	if !ok {
		return errors.New("shipment repository is not an InMemoryShipmentRepository")
	}

	if len(repo.GetAll()) != 0 {
		return errors.New("shipment repository has shipments")
	}

	return nil
}

func (ss *ShipmentSteps) ThenRepositoryNotHasAnyShipmentStatus() error {
	repo, ok := ss.testContext.ShipmentStatusRepository.(*shipmentstatusrepository.InMemoryShipmentStatusRepository)
	if !ok {
		return errors.New("shipment status repository is not an InMemoryShipmentRepository")
	}

	if len(repo.GetAll()) != 0 {
		return errors.New("shipment status repository has items")
	}

	return nil
}
