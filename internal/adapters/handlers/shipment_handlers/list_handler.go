package shipment_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (h *HTTPShipmentHandler) ListShipmentWithStatuses(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	emailVal := ctx.Value(domain.EmailKey)
	email, ok := emailVal.(string)
	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	pageVal := ctx.Value(domain.PageKey)
	page, ok := pageVal.(int)
	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	limitVal := ctx.Value(domain.LimitKey)
	limit, ok := limitVal.(int)
	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	pagination := &domain.PaginationRequest{
		Page:  page,
		Limit: limit,
	}

	shipmentsStatusResult, err := h.service.GetAllWithStatuses(r.Context(), email, pagination)
	if err != nil {
		h.handleError(w, err)
		return
	}

	paginationResult := domain.PaginationResult[*domain.ShipmentWithCurrentStatus]{
		Data: shipmentsStatusResult,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(paginationResult)
}
