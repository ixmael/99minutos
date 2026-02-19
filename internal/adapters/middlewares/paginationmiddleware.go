package middlewares

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ixmael/99minutos/internal/core/domain"
)

// PaginationMiddleware
func PaginationMiddleware() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pageStr := r.URL.Query().Get("page")
			limitStr := r.URL.Query().Get("limit")

			page, err := strconv.Atoi(pageStr)
			if err != nil || page < 1 {
				page = 1
			}

			limit, err := strconv.Atoi(limitStr)
			if err != nil || limit < 1 {
				limit = domain.DefaultLimit
			}

			ctx := context.WithValue(r.Context(), domain.PageKey, page)
			ctx = context.WithValue(ctx, domain.LimitKey, limit)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
