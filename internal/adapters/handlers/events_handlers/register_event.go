package events_handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (h *HTTTEventHandler) RegisterEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	idempotencyKeyVal := ctx.Value(domain.IdempotencyKey)
	idempontencyKey, ok := idempotencyKeyVal.(string)
	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var webEventRequest struct {
		TrackingNumber string    `json:"tracking_number"`
		Status         string    `json:"status"`
		Timestamp      time.Time `json:"timestamp"`
		Source         string    `json:"source"`
		Location       struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
	}

	if err := json.NewDecoder(r.Body).Decode(&webEventRequest); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	shipmentEventRequest := domain.ShipmentEventRequest{
		IdempotencyKey: idempontencyKey,
		TrackingNumber: webEventRequest.TrackingNumber,
		Status:         webEventRequest.Status,
		Timestamp:      webEventRequest.Timestamp,
		Source:         webEventRequest.Source,
		Location: struct {
			Lat float64
			Lng float64
		}{
			Lat: webEventRequest.Location.Lat,
			Lng: webEventRequest.Location.Lng,
		},
	}

	status, err := h.service.RegisterEventAsync(ctx, &shipmentEventRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(status)
}
