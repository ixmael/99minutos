package user_handlers

import (
	"net/http"

	"github.com/ixmael/99minutos/internal/core/ports"
)

type HTTPUserHandler struct {
	service ports.UserService
}

func NewHTTPUserHandler(userService ports.UserService) *HTTPUserHandler {
	return &HTTPUserHandler{
		service: userService,
	}
}

func (h *HTTPUserHandler) handleError(w http.ResponseWriter, err error) {
	switch err {
	default:
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
