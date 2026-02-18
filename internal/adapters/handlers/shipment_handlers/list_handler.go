package shipment_handlers

import (
	"encoding/json"
	"net/http"
)

func (h *HTTPShipmentHandler) ListShipmentWithStatuses(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	shipmentsStatusResult, err := h.service.GetAllWithStatuses(r.Context())
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shipmentsStatusResult)
}
