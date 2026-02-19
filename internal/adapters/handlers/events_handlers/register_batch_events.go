package events_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (h *HTTTEventHandler) RegisterBatchEvents(w http.ResponseWriter, r *http.Request) {
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

	fmt.Printf("ID: %s\n", idempontencyKey)

	type webEventRequest struct {
		TrackingNumber string    `json:"tracking_number"`
		Status         string    `json:"status"`
		Timestamp      time.Time `json:"timestamp"`
		Source         string    `json:"source"`
		Location       struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
	}

	var webEventRequests []webEventRequest
	if err := json.NewDecoder(r.Body).Decode(&webEventRequests); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	events := []*domain.ShipmentEventRequest{}
	for _, webEventRequest := range webEventRequests {
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

		events = append(events, &shipmentEventRequest)
	}

	status, err := h.service.RegisterEventsBatchAsync(ctx, idempontencyKey, events)
	if err != nil {
		h.logger.Error("error registering event", "error", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(status)
}
