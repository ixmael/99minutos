package shipment_handlers

import (
	"encoding/json"
	"net/http"
)

func (h *HTTPShipmentHandler) ListShipmentDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	shipmentID := r.PathValue("shipment_id")
	if shipmentID == "" {
		http.Error(w, "Shipment ID is required", http.StatusBadRequest)
		return
	}

	shipmentResult, err := h.service.GetShipmentDetails(r.Context(), shipmentID)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shipmentResult)
}
