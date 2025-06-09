package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dfanso/reddit-clone/pkg/auth"
	"github.com/labstack/echo/v4"
)

// AuthMiddleware verifies the JWT token and extracts user data into context
func AuthMiddleware(jwtManager *auth.JWTManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header missing")
			}

			// Extract token from "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
			}

			tokenString := parts[1]
			claims, err := jwtManager.ValidateToken(tokenString)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
			}

			//TODO: Check if user is exists in database

			// Store user_id and role in request context
			ctx := context.WithValue(c.Request().Context(), "user_id", claims.UserID)
			ctx = context.WithValue(ctx, "role", claims.Role)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
