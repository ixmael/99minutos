package shipment_handlers

import (
	"net/http"

	"github.com/ixmael/99minutos/internal/core/ports"
)

type HTTPShipmentHandler struct {
	service ports.ShipmentService
}

func NewHTTPShipmentHandler(shipmentService ports.ShipmentService) *HTTPShipmentHandler {
	return &HTTPShipmentHandler{
		service: shipmentService,
	}
}

func (h *HTTPShipmentHandler) handleError(w http.ResponseWriter, err error) {
	switch err {
	default:
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
