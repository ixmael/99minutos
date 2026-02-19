package user_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/ixmael/99minutos/internal/core/domain"
)

const (
	secretKey = "tu_clave_secreta"
)

func (h *HTTPUserHandler) Auth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var webUserAuthRequest struct {
		Email         string `json:"email"`
		PlainPassword string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&webUserAuthRequest); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	user, err := h.service.Authenticate(r.Context(), webUserAuthRequest.Email, webUserAuthRequest.PlainPassword)
	if err != nil {
		h.handleError(w, err)
		return
	}
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	token, err := generateJWT(user.ID, user.Email, user.IsAdmin)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func generateJWT(userID int64, email string, isAdmin bool) (string, error) {
	role := domain.ClientRole
	if isAdmin {
		role = domain.AdminRole
	}

	claims := jwt.MapClaims{
		"sub":        fmt.Sprintf("%d", userID),
		"authorized": true,
		"email":      email,
		"role":       role,
		"exp":        time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}
