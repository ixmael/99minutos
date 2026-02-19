package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ixmael/99minutos/internal/core/domain"
)

const (
	secretKey = "tu_clave_secreta"
)

// AuthMiddleware
func AuthMiddleware() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(secretKey), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				ctx := context.WithValue(r.Context(), domain.RoleKey, claims["role"])
				ctx = context.WithValue(ctx, domain.EmailKey, claims["email"])
				ctx = context.WithValue(ctx, domain.UserIDKey, claims["sub"])

				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				http.Error(w, "Invalid claims", http.StatusUnauthorized)
			}
		})
	}
}
