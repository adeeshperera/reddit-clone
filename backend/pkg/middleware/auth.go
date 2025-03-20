package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dfanso/reddit-clone/pkg/auth"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware verifies the JWT token and extracts user data into context
func AuthMiddleware(jwtManager *auth.JWTManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header missing", http.StatusUnauthorized)
				return
			}

			// Extract token from "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
				return
			}

			tokenString := parts[1]
			token, err := jwtManager.VerifyToken(tokenString)
			if err != nil || !token.Valid {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			// Store user_id and role in request context
			ctx := context.WithValue(r.Context(), "user_id", claims["user_id"])
			ctx = context.WithValue(ctx, "role", claims["role"])
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
