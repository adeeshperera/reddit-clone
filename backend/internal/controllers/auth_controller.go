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
}

func NewAuthController(userService *services.UserService) *AuthController {
	return &AuthController{
		userService: userService,
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
