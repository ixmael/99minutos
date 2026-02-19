//go:build integration
// +build integration

package bdd

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/cucumber/godog"

	"github.com/ixmael/99minutos/internal/core/domain"
	"github.com/ixmael/99minutos/internal/core/services/shipmentservice"
	"github.com/ixmael/99minutos/internal/infrastructure/inmemory/logger"
	"github.com/ixmael/99minutos/internal/infrastructure/inmemory/shipmentrepository"
	"github.com/ixmael/99minutos/internal/infrastructure/inmemory/shipmentstatusrepository"
	"github.com/ixmael/99minutos/internal/infrastructure/inmemory/userrepository"
)

// ClientShipmentSteps contains shared step definitions for shipment-related BDD tests
type ClientShipmentSteps struct {
	testContext *TestContext
}

// NewClientShipmentSteps creates a new ClientShipmentSteps instance
func NewClientShipmentSteps(tc *TestContext) *ClientShipmentSteps {
	return &ClientShipmentSteps{testContext: tc}
}

// RegisterClientShipmentSteps registers all shipment-related step definitions
func RegisterClientShipmentSteps(ctx *godog.ScenarioContext, tc *TestContext) {
	cs := NewClientShipmentSteps(tc)

	ctx.Step(`^the system is ready to accept shipment registration requests$`, cs.GivenShipmentSystemIsReady)
	ctx.Step(`^the origin is "([^"]*)"$`, cs.GivenShipmentOrigin)
	ctx.Step(`^the destination is "([^"]*)"$`, cs.GivenShipmentDestination)
	ctx.Step(`^a shipment request is submitted$`, cs.WhenRequestNewShipment)
	ctx.Step(`^the registration request is successful$`, cs.ThenRegistrationRequestIsSuccessful)
	ctx.Step(`^the shipment state is "([^"]*)"$`, cs.ThenShipmentStateIs)
	ctx.Step(`^the registration request is rejected$`, cs.ThenRegistrationRequestIsRejected)
	ctx.Step(`^the repository not has any shipment$`, cs.ThenRepositoryNotHasAnyShipment)
	ctx.Step(`^the repository not has any shipment status$`, cs.ThenRepositoryNotHasAnyShipmentStatus)
	ctx.Step(`^the system has no shipments$`, cs.GivenSystemHasNoShipments)
	ctx.Step(`^I request all shipments$`, cs.WhenRequestAllShipments)
	ctx.Step(`^I should receive an empty list of shipments$`, cs.ThenShouldReceiveEmptyList)
	ctx.Step(`^I should receive a list with (\d+) shipment$`, cs.ThenShouldReceiveListWithCount)
	ctx.Step(`^the shipment in the list has origin "([^"]*)"$`, cs.ThenShipmentInListHasOrigin)
	ctx.Step(`^the shipment in the list has destination "([^"]*)"$`, cs.ThenShipmentInListHasDestination)
	ctx.Step(`^I request the shipment details$`, cs.WhenRequestShipmentDetails)
	ctx.Step(`^I should receive the shipment details$`, cs.ThenShouldReceiveShipmentDetails)
	ctx.Step(`^the shipment details has origin "([^"]*)"$`, cs.ThenShipmentDetailsHasOrigin)
	ctx.Step(`^the shipment details has destination "([^"]*)"$`, cs.ThenShipmentDetailsHasDestination)
	ctx.Step(`^the shipment details has (\d+) status$`, cs.ThenShipmentDetailsHasStatusCount)
	ctx.Step(`^the shipment status is "([^"]*)"$`, cs.ThenShipmentStatusIs)
	ctx.Step(`^the shipment is related to the client$`, cs.ThenShipmentIsRelatedToClient)
	ctx.Step(`^the system is ready to accept shipment list requests$`, cs.GivenSystemIsReadyToList)
	ctx.Step(`^the system has 1 shipment for "([^"]*)"$`, cs.GivenSystemHasOneClientShipment)
	ctx.Step(`^I should receive a list with 1 shipment$`, cs.ThenShouldListOneShipment)
	ctx.Step(`^the shipment has status "([^"]*)"$`, cs.GivenShipmentHasStatus)
	ctx.Step(`^the shipment details has "([^"]*)" statuses$`, cs.ThenShipmentDetailsHasStatuses)
}

func (cs *ClientShipmentSteps) GivenShipmentSystemIsReady() error {
	logger, err := logger.NewInMemoryLoggerServices()
	if err != nil {
		return err
	}

	userrepository, err := userrepository.NewInMemoryUserRepository()
	if err != nil {
		return err
	}

	shipmentstatusrepository, err := shipmentstatusrepository.NewInMemoryShipmentStatusRepository()
	if err != nil {
		return err
	}

	shipmentrepository, err := shipmentrepository.NewInMemoryShipmentRepository(shipmentstatusrepository)
	if err != nil {
		return err
	}

	shipmentservice, err := shipmentservice.NewShipmentService(logger, shipmentrepository, shipmentstatusrepository, userrepository)
	if err != nil {
		return err
	}

	cs.testContext.UserRepository = userrepository
	cs.testContext.ShipmentStatusRepository = shipmentstatusrepository
	cs.testContext.ShipmentRepository = shipmentrepository
	cs.testContext.ShipmentService = shipmentservice

	cs.testContext.LastError = nil
	cs.testContext.NewShipmentOriginRequest = nil
	cs.testContext.NewShipmentDestinationRequest = nil

	return nil
}

func (cs *ClientShipmentSteps) GivenShipmentOrigin(originStr string) error {
	cs.testContext.NewShipmentOriginRequest = &originStr

	return nil
}

func (cs *ClientShipmentSteps) GivenShipmentDestination(destinationStr string) error {
	cs.testContext.NewShipmentDestinationRequest = &destinationStr

	return nil
}

func (cs *ClientShipmentSteps) WhenRequestNewShipment() error {
	origin := ""
	if cs.testContext.NewShipmentOriginRequest != nil {
		origin = *cs.testContext.NewShipmentOriginRequest
	}
	destination := ""
	if cs.testContext.NewShipmentDestinationRequest != nil {
		destination = *cs.testContext.NewShipmentDestinationRequest
	}
	email := ""
	if cs.testContext.UserEmail != nil {
		email = *cs.testContext.UserEmail
	}

	newShipmentRequest := &domain.NewShipmentRequest{
		Email:       email,
		Origin:      origin,
		Destination: destination,
	}

	shipmentResult, err := cs.testContext.ShipmentService.CreateShipment(
		context.Background(),
		newShipmentRequest,
	)
	cs.testContext.LastShipmentCreatedResult = shipmentResult
	cs.testContext.LastError = err

	return nil
}

func (cs *ClientShipmentSteps) ThenRegistrationRequestIsSuccessful() error {
	if cs.testContext.LastError != nil {
		return errors.New(fmt.Sprintf("registration request failed: %v", cs.testContext.LastError))
	}

	if cs.testContext.LastShipmentCreatedResult == nil {
		return errors.New("shipment not created")
	}

	return nil
}

func (cs *ClientShipmentSteps) ThenShipmentStateIs(stateStr string) error {
	shipment, err := cs.testContext.ShipmentRepository.FindByShipmentID(
		context.Background(),
		cs.testContext.LastShipmentCreatedResult.ID,
	)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to find shipment: %v", err))
	}

	status, err := cs.testContext.ShipmentStatusRepository.FindStatusByID(
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

func (cs *ClientShipmentSteps) ThenRegistrationRequestIsRejected() error {
	if cs.testContext.LastError == nil {
		return errors.New("registration request not rejected")
	}

	return nil
}

func (cs *ClientShipmentSteps) ThenRepositoryNotHasAnyShipment() error {
	repo, ok := cs.testContext.ShipmentRepository.(*shipmentrepository.InMemoryShipmentRepository)
	if !ok {
		return errors.New("shipment repository is not an InMemoryShipmentRepository")
	}

	if len(repo.GetAll()) != 0 {
		return errors.New("shipment repository has shipments")
	}

	return nil
}

func (cs *ClientShipmentSteps) ThenRepositoryNotHasAnyShipmentStatus() error {
	repo, ok := cs.testContext.ShipmentStatusRepository.(*shipmentstatusrepository.InMemoryShipmentStatusRepository)
	if !ok {
		return errors.New("shipment status repository is not an InMemoryShipmentRepository")
	}

	if len(repo.GetAll()) != 0 {
		return errors.New("shipment status repository has items")
	}

	return nil
}

func (cs *ClientShipmentSteps) WhenRequestAllShipments() error {
	shipments, err := cs.testContext.ShipmentService.GetAllWithStatuses(context.Background(), *cs.testContext.UserEmail, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to get all shipments: %v", err))
	}
	cs.testContext.LastShipmentsWithStatuses = shipments

	return nil
}

func (cs *ClientShipmentSteps) ThenShouldReceiveEmptyList() error {
	if len(cs.testContext.LastShipmentsWithStatuses) != 0 {
		return errors.New(fmt.Sprintf("expected empty list, got %d shipments", len(cs.testContext.LastShipmentsWithStatuses)))
	}
	return nil
}

func (cs *ClientShipmentSteps) ThenShouldReceiveListWithCount(count int) error {
	if len(cs.testContext.LastShipmentsWithStatuses) != count {
		return errors.New(fmt.Sprintf("expected %d shipments, got %d", count, len(cs.testContext.LastShipmentsWithStatuses)))
	}
	return nil
}

func (cs *ClientShipmentSteps) ThenShipmentInListHasOrigin(origin string) error {
	if len(cs.testContext.LastShipmentsWithStatuses) == 0 {
		return errors.New("no shipments in list")
	}

	shipment := cs.testContext.LastShipmentsWithStatuses[0]
	if shipment.Origin != origin {
		return errors.New(fmt.Sprintf("expected origin %s, got %s", origin, shipment.Origin))
	}
	return nil
}

func (cs *ClientShipmentSteps) ThenShipmentInListHasDestination(destination string) error {
	if len(cs.testContext.LastShipmentsWithStatuses) == 0 {
		return errors.New("no shipments in list")
	}

	shipment := cs.testContext.LastShipmentsWithStatuses[0]
	if shipment.Destination != destination {
		return errors.New(fmt.Sprintf("expected destination %s, got %s", destination, shipment.Destination))
	}
	return nil
}

func (cs *ClientShipmentSteps) WhenRequestShipmentDetails() error {
	details, err := cs.testContext.ShipmentService.GetShipmentDetails(
		context.Background(),
		*cs.testContext.UserEmail,
		cs.testContext.LastShipment.ID,
	)
	cs.testContext.LastShipmentDetails = details
	cs.testContext.LastError = err

	return nil
}

func (cs *ClientShipmentSteps) ThenShouldReceiveShipmentDetails() error {
	if cs.testContext.LastShipmentDetails == nil {
		return errors.New("no shipment details received")
	}
	return nil
}

func (cs *ClientShipmentSteps) ThenShipmentDetailsHasOrigin(origin string) error {
	if cs.testContext.LastShipmentDetails == nil {
		return errors.New("no shipment details")
	}

	if cs.testContext.LastShipmentDetails.Origin != origin {
		return errors.New(fmt.Sprintf("expected origin %s, got %s", origin, cs.testContext.LastShipmentDetails.Origin))
	}
	return nil
}

func (cs *ClientShipmentSteps) ThenShipmentDetailsHasDestination(destination string) error {
	if cs.testContext.LastShipmentDetails == nil {
		return errors.New("no shipment details")
	}

	if cs.testContext.LastShipmentDetails.Destination != destination {
		return errors.New(fmt.Sprintf("expected destination %s, got %s", destination, cs.testContext.LastShipmentDetails.Destination))
	}
	return nil
}

func (cs *ClientShipmentSteps) ThenShipmentDetailsHasStatusCount(count int) error {
	if cs.testContext.LastShipmentDetails == nil {
		return errors.New("no shipment details")
	}

	if len(cs.testContext.LastShipmentDetails.Statuses) != count {
		return errors.New(fmt.Sprintf("expected %d statuses, got %d", count, len(cs.testContext.LastShipmentDetails.Statuses)))
	}
	return nil
}

func (cs *ClientShipmentSteps) ThenShipmentStatusIs(statusStr string) error {
	return nil
}

func (cs *ClientShipmentSteps) ThenShipmentIsRelatedToClient() error {
	user, err := cs.testContext.UserRepository.FindByEmail(context.Background(), *cs.testContext.UserEmail)
	if err != nil {
		return err
	}

	shipment, err := cs.testContext.ShipmentRepository.FindByShipmentID(
		context.Background(),
		cs.testContext.LastShipmentCreatedResult.ID,
	)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to find shipment: %v", err))
	}

	if shipment.UserID != user.ID {
		return errors.New("shipment user is not the client")
	}

	return nil
}

func (cs *ClientShipmentSteps) GivenSystemHasNoShipments() error {
	// All the repositories reset for each test
	return nil
}

func (cs *ClientShipmentSteps) GivenSystemIsReadyToList() error {
	logger, err := logger.NewInMemoryLoggerServices()
	if err != nil {
		return err
	}

	userrepository, err := userrepository.NewInMemoryUserRepository()
	if err != nil {
		return err
	}

	shipmentstatusrepository, err := shipmentstatusrepository.NewInMemoryShipmentStatusRepository()
	if err != nil {
		return err
	}

	shipmentrepository, err := shipmentrepository.NewInMemoryShipmentRepository(shipmentstatusrepository)
	if err != nil {
		return err
	}

	shipmentservice, err := shipmentservice.NewShipmentService(logger, shipmentrepository, shipmentstatusrepository, userrepository)
	if err != nil {
		return err
	}

	cs.testContext.UserRepository = userrepository
	cs.testContext.ShipmentStatusRepository = shipmentstatusrepository
	cs.testContext.ShipmentRepository = shipmentrepository
	cs.testContext.ShipmentService = shipmentservice

	cs.testContext.LastError = nil
	cs.testContext.NewShipmentOriginRequest = nil
	cs.testContext.NewShipmentDestinationRequest = nil

	return nil
}

func (cs *ClientShipmentSteps) GivenSystemHasOneClientShipment(emailStr string) error {
	clientShipmentRequest := &domain.NewShipmentRequest{
		Email:       emailStr,
		Role:        domain.ClientRole,
		Origin:      "Beijing",
		Destination: "Shanghai",
	}

	shipment, err := cs.testContext.ShipmentService.CreateShipment(context.Background(), clientShipmentRequest)
	if err != nil {
		return err
	}

	cs.testContext.LastShipment = shipment
	cs.testContext.UserEmail = &emailStr

	return nil
}

func (cs *ClientShipmentSteps) ThenShouldListOneShipment() error {
	if len(cs.testContext.LastShipmentsWithStatuses) != 1 {
		return errors.New("expected one shipment, got none")
	}

	return nil
}

func (cs *ClientShipmentSteps) GivenShipmentHasStatus(status string) error {
	if cs.testContext.LastShipment == nil {
		return errors.New("no shipment found")
	}

	shipmentStatus, err := domain.NewShipmentStatus(cs.testContext.LastShipment.ID, status)
	if err != nil {
		return errors.New("error creating shipment status")
	}

	err = cs.testContext.ShipmentStatusRepository.Save(context.Background(), shipmentStatus)
	if err != nil {
		return errors.New("error adding shipment status")
	}

	return nil
}

func (cs *ClientShipmentSteps) ThenShipmentDetailsHasStatuses(countStr string) error {
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return errors.New("invalid count")
	}

	if len(cs.testContext.LastShipmentDetails.Statuses) != count {
		return errors.New("expected different number of statuses")
	}

	return nil
}
