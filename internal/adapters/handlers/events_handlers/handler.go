package events_handlers

import (
	"net/http"

	"github.com/ixmael/99minutos/internal/core/ports"
)

type HTTTEventHandler struct {
	service ports.ShipmentService
	logger  ports.Logger
}

func NewHTTTEventHandler(shipmentservice ports.ShipmentService, logger ports.Logger) *HTTTEventHandler {
	return &HTTTEventHandler{
		service: shipmentservice,
		logger:  logger,
	}
}

func (h *HTTTEventHandler) handleError(w http.ResponseWriter, err error) {
	switch err {
	default:
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
