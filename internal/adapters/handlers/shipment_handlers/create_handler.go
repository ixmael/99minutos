package shipment_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (h *HTTPShipmentHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	roleVal := ctx.Value(domain.RoleKey)
	emailVal := ctx.Value(domain.EmailKey)

	role, ok := roleVal.(string)
	if !ok {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	email, ok := emailVal.(string)
	if !ok {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	var webShipmentRequest struct {
		Origin      string `json:"origin"`
		Destination string `json:"destination"`
	}

	if err := json.NewDecoder(r.Body).Decode(&webShipmentRequest); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	newShipmentRequest := domain.NewShipmentRequest{
		Origin:      webShipmentRequest.Origin,
		Destination: webShipmentRequest.Destination,
		Email:       email,
		Role:        role,
	}

	shipmentResult, err := h.service.CreateShipment(r.Context(), &newShipmentRequest)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"shipment_id": shipmentResult.ID,
	})
}
