package shipment_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (h *HTTPShipmentHandler) ListShipmentDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	emailVal := ctx.Value(domain.EmailKey)
	email, ok := emailVal.(string)
	if !ok {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	shipmentID := r.PathValue("shipment_id")
	if shipmentID == "" {
		http.Error(w, "Shipment ID is required", http.StatusBadRequest)
		return
	}

	shipmentResult, err := h.service.GetShipmentDetails(r.Context(), email, shipmentID)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shipmentResult)
}
