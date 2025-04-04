package controllers

import (
	"net/http"

	dto "github.com/dfanso/reddit-clone/internal/dtos"
	"github.com/dfanso/reddit-clone/internal/services"
	"github.com/dfanso/reddit-clone/pkg/utils"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	userService *services.UserService
	authService *services.AuthService
}

func NewAuthController(userService *services.UserService, authService *services.AuthService) *AuthController {
	return &AuthController{
		userService: userService,
		authService: authService,
	}
}

func (c *AuthController) Register(ctx echo.Context) error {

	// Bind request body to RegisterRequest DTO
	var req dto.RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
	}

	// Validate the DTO
	if err := req.Validate(); err != nil {
		if e, ok := err.(validation.Errors); ok {
			return utils.ErrorResponse(ctx, http.StatusBadRequest, "Validation failed", e)
		}
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid registration data", err)
	}

	// Register the user via the service layer
	user, err := c.userService.RegisterUser(ctx.Request().Context(), req)
	if err != nil {
		return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to register user", err)
	}

	// Return success response
	return utils.SuccessResponse(ctx, http.StatusCreated, "User registered successfully", user)
}

func (c *AuthController) Login(ctx echo.Context) error {
	// Bind request body to LoginRequest DTO
	var req dto.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
	}

	// Validate the DTO
	if err := req.Validate(); err != nil {
		if e, ok := err.(validation.Errors); ok {
			return utils.ErrorResponse(ctx, http.StatusBadRequest, "Validation failed", e)
		}
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid login data", err)
	}

	// Authenticate the user via the auth service
	user, err := c.authService.Login(ctx.Request().Context(), req)
	if err != nil {
		return utils.ErrorResponse(ctx, http.StatusUnauthorized, "Login failed", err)
	}

	// Return success response
	return utils.SuccessResponse(ctx, http.StatusOK, "Login successful", user)
}
