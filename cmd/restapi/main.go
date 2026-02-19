package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pelletier/go-toml/v2"

	"github.com/ixmael/99minutos/internal/adapters/handlers/shipment_handlers"
	"github.com/ixmael/99minutos/internal/adapters/handlers/user_handlers"
	"github.com/ixmael/99minutos/internal/adapters/middlewares"
	"github.com/ixmael/99minutos/internal/core/services/shipmentservice"
	"github.com/ixmael/99minutos/internal/core/services/userservice"
	"github.com/ixmael/99minutos/internal/infrastructure/postgres/shipmentrepository"
	"github.com/ixmael/99minutos/internal/infrastructure/postgres/shipmentstatusrepository"
	"github.com/ixmael/99minutos/internal/infrastructure/postgres/userrepository"
	"github.com/ixmael/99minutos/internal/infrastructure/zaplogger"
)

// ApplicationConfig represents the configuration for the application.
type ApplicationConfig struct {
	Environment string `toml:"environment"`
	RestAPI     struct {
		Port int `toml:"port"`
	} `toml:"restapi"`
	Repository struct {
		PostgresURL string `toml:"postgres"`
	} `toml:"repository"`
}

func main() {
	configFilePathFlag := flag.String("config", ".env.toml", "The configuration file path")

	flag.Parse()

	configData, err := os.ReadFile(*configFilePathFlag)
	if err != nil {
		log.Println("error on reading the config file")
		os.Exit(1)
		return
	}

	var cnfg ApplicationConfig
	err = toml.Unmarshal(configData, &cnfg)
	if err != nil {
		log.Println("error on parsing the config file")
		os.Exit(1)
		return
	}

	zaplogger, err := zaplogger.NewZapLogger(cnfg.Environment)
	if err != nil {
		log.Println("error on initializing the logger")
		os.Exit(1)
		return
	}

	userpostgresrepository, err := userrepository.NewPostgresUserStatusRepository(cnfg.Repository.PostgresURL, zaplogger)
	if err != nil {
		log.Println("error on initializing the user repository")
		os.Exit(1)
		return
	}

	shipmentpostgresrepository, err := shipmentrepository.NewPostgresShipmentRepository(cnfg.Repository.PostgresURL, zaplogger)
	if err != nil {
		log.Println("error on initializing the shipment repository")
		os.Exit(1)
		return
	}

	shipmentstatuspostgresrepository, err := shipmentstatusrepository.NewPostgresShipmentStatusRepository(cnfg.Repository.PostgresURL, zaplogger)
	if err != nil {
		log.Println("error on initializing the shipment repository")
		os.Exit(1)
		return
	}

	shipmentService, err := shipmentservice.NewShipmentService(
		zaplogger,
		shipmentpostgresrepository,
		shipmentstatuspostgresrepository,
		userpostgresrepository,
	)
	if err != nil {
		log.Println("error on setup the shipment service")
		os.Exit(1)
		return
	}

	userService, err := userservice.NewUserService(zaplogger, userpostgresrepository)
	if err != nil {
		log.Println("error on setup the shipment service")
		os.Exit(1)
		return
	}

	shipmentHandlers := shipment_handlers.NewHTTPShipmentHandler(shipmentService)
	userHandlers := user_handlers.NewHTTPUserHandler(userService)

	handlers := http.NewServeMux()
	authMiddleware := middlewares.AuthMiddleware()

	handlers.Handle("POST /shipments", authMiddleware(shipmentHandlers.Register))
	handlers.HandleFunc("GET /shipments/{shipment_id}", authMiddleware(shipmentHandlers.ListShipmentDetails))
	handlers.HandleFunc("GET /shipments", authMiddleware(shipmentHandlers.ListShipmentWithStatuses))

	handlers.HandleFunc("POST /account", userHandlers.CreateUser)
	handlers.HandleFunc("POST /auth", userHandlers.Auth)

	handlers.HandleFunc("GET /health", healthStatus)

	port := fmt.Sprintf(":%d", cnfg.RestAPI.Port)
	api := &http.Server{
		Addr:    port,
		Handler: handlers,
	}

	go func() {
		zaplogger.Info(fmt.Sprintf("API is running on %s", port))
		if err := api.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutdown signal received")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := api.Shutdown(ctx); err != nil {
		log.Printf("Shutdown error: %v", err)
		os.Exit(0)
	}

	log.Println("Server stopped")
	os.Exit(0)
}

func healthStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
