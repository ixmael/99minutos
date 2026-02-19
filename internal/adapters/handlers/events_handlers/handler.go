package events_handlers

import (
	"net/http"

	"github.com/ixmael/99minutos/internal/core/ports"
)

type HTTTEventHandler struct {
	service ports.ShipmentService
}

func NewHTTTEventHandler(shipmentservice ports.ShipmentService) *HTTTEventHandler {
	return &HTTTEventHandler{
		service: shipmentservice,
	}
}

func (h *HTTTEventHandler) handleError(w http.ResponseWriter, err error) {
	switch err {
	default:
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
