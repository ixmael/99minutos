package main

import (
	"github.com/ixmael/99minutos/internal/core/ports"
	"github.com/ixmael/99minutos/internal/core/services/shipmentservice"
	"github.com/ixmael/99minutos/internal/core/services/userservice"
	"github.com/ixmael/99minutos/internal/infrastructure/postgres"
	"github.com/ixmael/99minutos/internal/infrastructure/postgres/shipmentrepository"
	"github.com/ixmael/99minutos/internal/infrastructure/postgres/shipmentstatusrepository"
	"github.com/ixmael/99minutos/internal/infrastructure/postgres/userrepository"
	"github.com/ixmael/99minutos/internal/infrastructure/rabbitmq"
	"github.com/ixmael/99minutos/internal/infrastructure/valkey"
	"github.com/ixmael/99minutos/internal/infrastructure/zaplogger"
)

type ApplicationServices struct {
	Logger          ports.Logger
	ShipmentService ports.ShipmentService
	UserService     ports.UserService

	userRepository           ports.UserRepository
	shipmentRepository       ports.ShipmentRepository
	shipmentStatusRepository ports.ShipmentStatusRepository
	queueService             ports.QueueService
	cacheService             ports.CacheService
}

// SetupServices initializes and returns the application services.
func SetupServices(config *ApplicationConfig) (*ApplicationServices, error) {
	zaplogger, err := zaplogger.NewZapLogger(config.Environment)
	if err != nil {
		return nil, err
	}

	err = postgres.SetupPostgresRepository(zaplogger, config.Repository.PostgresURL, config.Repository.PostgresMigrationsPath)
	if err != nil {
		zaplogger.Error("Failed to setup Postgres repository", "error", err)
		return nil, err
	}

	userpostgresrepository, err := userrepository.NewPostgresUserStatusRepository(config.Repository.PostgresURL, zaplogger)
	if err != nil {
		zaplogger.Error("Failed to start Postgres User repository", "error", err)
		return nil, err
	}

	shipmentpostgresrepository, err := shipmentrepository.NewPostgresShipmentRepository(config.Repository.PostgresURL, zaplogger)
	if err != nil {
		zaplogger.Error("Failed to start Postgres Shipment repository", "error", err)
		return nil, err
	}

	shipmentstatuspostgresrepository, err := shipmentstatusrepository.NewPostgresShipmentStatusRepository(config.Repository.PostgresURL, zaplogger)
	if err != nil {
		zaplogger.Error("Failed to start Postgres ShipmentStatus repository", "error", err)
		return nil, err
	}

	queueService, err := rabbitmq.NewRabbitMQQueue(config.Queue.RabbitMQURL, zaplogger)
	if err != nil {
		zaplogger.Error("Failed to start RabbitMQ queue", "error", err)
		return nil, err
	}

	valkeyService, err := valkey.NewValkeyService(config.Cache.ValkeyURL, zaplogger)
	if err != nil {
		zaplogger.Error("Failed to start Valkey cache", "error", err)
		return nil, err
	}

	shipmentService, err := shipmentservice.NewShipmentService(
		zaplogger,
		shipmentpostgresrepository,
		shipmentstatuspostgresrepository,
		userpostgresrepository,
		queueService,
		valkeyService,
	)
	if err != nil {
		zaplogger.Error("Failed to start shipment service", "error", err)
		return nil, err
	}

	userService, err := userservice.NewUserService(zaplogger, userpostgresrepository)
	if err != nil {
		zaplogger.Error("Failed to start user service", "error", err)
		return nil, err
	}

	services := ApplicationServices{
		Logger:                   zaplogger,
		ShipmentService:          shipmentService,
		UserService:              userService,
		userRepository:           userpostgresrepository,
		shipmentRepository:       shipmentpostgresrepository,
		shipmentStatusRepository: shipmentstatuspostgresrepository,
		queueService:             queueService,
		cacheService:             valkeyService,
	}

	return &services, nil
}

func (services *ApplicationServices) Stop() {
	services.Logger.Stop()
	services.userRepository.Stop()
	services.shipmentRepository.Stop()
	services.shipmentStatusRepository.Stop()
}
