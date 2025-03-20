package controllers

import (
	"net/http"

	"github.com/dfanso/reddit-clone/dtos"
	"github.com/dfanso/reddit-clone/internal/models"
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

	// Check if email already exists
	filter := map[string]any{"email": req.Email}
	existingUser, err := c.userService.FindOne(ctx.Request().Context(), filter)
	if err != nil {
		return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Error checking existing user", err)
	}
	if existingUser != nil {
		return utils.ErrorResponse(ctx, http.StatusConflict, "Email already in use", nil)
	}

	// Check if username already exists
	filter = map[string]any{"handler": req.Username}
	existingUser, err = c.userService.FindOne(ctx.Request().Context(), filter)
	if err != nil {
		return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Error checking existing username", err)
	}
	if existingUser != nil {
		return utils.ErrorResponse(ctx, http.StatusConflict, "Username already taken", nil)
	}

	// Create new user
	user := &models.User{
		Email:   req.Email,
		Handler: req.Username,
		Name:    req.Username, // Use username as name for simplicity
		Role:    models.RoleUser,
		Status:  models.StatusUnverified,
		Stage:   models.StageEmailVerification,
	}
	user.Password = req.Password

	// Hash the password
	if err := user.HashPassword(); err != nil {
		return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to process password", err)
	}

	// Save the user
	if err := c.userService.Create(ctx.Request().Context(), user); err != nil {
		return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create user", err)
	}

	// Return success response
	return utils.SuccessResponse(ctx, http.StatusCreated, "User registered successfully", user)
}
