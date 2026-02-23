package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/ixmael/99minutos/internal/adapters/handlers/events_handlers"
	"github.com/ixmael/99minutos/internal/adapters/handlers/shipment_handlers"
	"github.com/ixmael/99minutos/internal/adapters/handlers/user_handlers"
	"github.com/ixmael/99minutos/internal/adapters/middlewares"
)

func SetupAPI(config *ApplicationConfig, services *ApplicationServices) *http.Server {
	shipmentHandlers := shipment_handlers.NewHTTPShipmentHandler(services.ShipmentService)
	userHandlers := user_handlers.NewHTTPUserHandler(services.UserService)
	eventsHandlers := events_handlers.NewHTTTEventHandler(services.ShipmentService, services.Logger)

	handlers := http.NewServeMux()
	authMiddleware := middlewares.AuthMiddleware()
	paginationMiddleware := middlewares.PaginationMiddleware()
	idempontencyMiddleware := middlewares.IdempontencyMiddleware()

	handlers.Handle("POST /shipments", authMiddleware(shipmentHandlers.Register))
	handlers.HandleFunc("GET /shipments/{shipment_id}", authMiddleware(shipmentHandlers.ListShipmentDetails))
	handlers.HandleFunc("GET /shipments", authMiddleware(paginationMiddleware(shipmentHandlers.ListShipmentWithStatuses)))

	handlers.HandleFunc("POST /account", userHandlers.CreateUser)
	handlers.HandleFunc("POST /auth", userHandlers.Auth)

	handlers.HandleFunc("POST /events", idempontencyMiddleware(eventsHandlers.RegisterEvent))
	handlers.HandleFunc("POST /events/batch", idempontencyMiddleware(eventsHandlers.RegisterBatchEvents))

	handlers.HandleFunc("GET /health", healthStatus(config))

	port := fmt.Sprintf(":%d", config.RestAPI.Port)

	api := &http.Server{
		Addr:    port,
		Handler: handlers,
	}

	return api
}

func healthStatus(config *ApplicationConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name, err := os.Hostname()
		if err != nil {
			name = "unknown"
		}

		okResult := struct {
			Status      string `json:"status"`
			Version     string `json:"version"`
			Environment string `json:"environment"`
			Host        string `json:"host"`
		}{
			Status:      "OK",
			Version:     config.Version,
			Environment: config.Environment,
			Host:        name,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(okResult)
	}
}
