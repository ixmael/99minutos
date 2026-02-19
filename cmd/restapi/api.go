package main

import (
	"fmt"
	"net/http"

	"github.com/ixmael/99minutos/internal/adapters/handlers/events_handlers"
	"github.com/ixmael/99minutos/internal/adapters/handlers/shipment_handlers"
	"github.com/ixmael/99minutos/internal/adapters/handlers/user_handlers"
	"github.com/ixmael/99minutos/internal/adapters/middlewares"
)

func SetupAPI(config *ApplicationConfig, services *ApplicationServices) *http.Server {
	shipmentHandlers := shipment_handlers.NewHTTPShipmentHandler(services.ShipmentService)
	userHandlers := user_handlers.NewHTTPUserHandler(services.UserService)
	eventsHandlers := events_handlers.NewHTTTEventHandler(services.ShipmentService)

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

	handlers.HandleFunc("GET /health", healthStatus)

	port := fmt.Sprintf(":%d", config.RestAPI.Port)

	api := &http.Server{
		Addr:    port,
		Handler: handlers,
	}

	return api
}

func healthStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
