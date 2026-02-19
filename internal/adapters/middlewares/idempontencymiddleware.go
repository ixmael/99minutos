package middlewares

import (
	"context"
	"net/http"

	"github.com/ixmael/99minutos/internal/core/domain"
)

// PaginationMiddleware
func IdempontencyMiddleware() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			idempontencyValue := r.Header.Get("Idempotency-Key")

			if idempontencyValue == "" {
				http.Error(w, "The idempotency key is invalid", http.StatusBadRequest)
				return
			}

			ctx := context.WithValue(r.Context(), domain.IdempotencyKey, idempontencyValue)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
