package user_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ixmael/99minutos/internal/core/domain"
)

func (h *HTTPUserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var webUserRequest struct {
		Email         string `json:"email"`
		PlainPassword string `json:"password"`
		IsAdmin       bool   `json:"is_admin,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&webUserRequest); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	newUserRequest := domain.UserRequest{
		Email:         webUserRequest.Email,
		PlainPassword: webUserRequest.PlainPassword,
		IsAdmin:       webUserRequest.IsAdmin,
	}

	userResult, err := h.service.CreateUser(r.Context(), &newUserRequest)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResult)
}
