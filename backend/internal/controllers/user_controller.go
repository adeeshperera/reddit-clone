package controllers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/dfanso/reddit-clone/internal/models"
	"github.com/dfanso/reddit-clone/internal/services"
	"github.com/dfanso/reddit-clone/pkg/utils"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (c *UserController) GetAll(ctx echo.Context) error {
	users, err := c.service.GetAll(ctx.Request().Context())
	if err != nil {
		return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get users", err)
	}
	return utils.SuccessResponse(ctx, http.StatusOK, "Users retrieved successfully", users)
}

func (c *UserController) GetPaginated(ctx echo.Context) error {
	fmt.Println("GetPaginated called")

	// Extract page and limit from query parameters
	pageStr := ctx.QueryParam("page")
	limitStr := ctx.QueryParam("limit")

	page, _ := strconv.Atoi(pageStr) // this like parsInt
	limit, _ := strconv.Atoi(limitStr)

	// Call the service method
	result, err := c.service.FindPaginated(ctx.Request().Context(), nil, page, limit)
	if err != nil {
		return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get paginated users", err)
	}

	return utils.SuccessResponse(ctx, http.StatusOK, "Paginated users retrieved successfully", result)
}

func (c *UserController) GetByID(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "User ID is required", nil)
	}

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID format", err)
	}

	user, err := c.service.GetByID(ctx.Request().Context(), parsedID)
	if err != nil {
		return utils.ErrorResponse(ctx, http.StatusNotFound, "User not found", err)
	}

	return utils.SuccessResponse(ctx, http.StatusOK, "User retrieved successfully", user)
}

func (c *UserController) Create(ctx echo.Context) error {
	var user models.User
	body, _ := io.ReadAll(ctx.Request().Body)
	fmt.Printf("Raw Request Body: %s\n", string(body))
	ctx.Request().Body = io.NopCloser(bytes.NewBuffer(body))

	if err := ctx.Bind(&user); err != nil {
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
	}

	if user.Role == "" {
		user.Role = models.RoleUser
	}
	if user.Status == "" {
		user.Status = models.StatusUnverified
	}

	// Validate all fields
	if err := user.Validate(); err != nil {
		if e, ok := err.(validation.Errors); ok {
			return utils.ErrorResponse(ctx, http.StatusBadRequest, "Validation failed", e)
		}
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user data", err)
	}

	// Hash password
	if err := user.HashPassword(); err != nil {
		return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to process password", err)
	}

	filter := map[string]interface{}{"email": user.Email}
	existingUser, err := c.service.FindOne(ctx.Request().Context(), filter)
	if err != nil {
		return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Error checking for existing user", err)
	}
	if existingUser != nil {
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "User with this email already exists", nil)
	}

	createdUser, err := c.service.Create(ctx.Request().Context(), &user)
	if err != nil {
		return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create user", err)
	}

	return utils.SuccessResponse(ctx, http.StatusCreated, "User created successfully", createdUser)
}

func (c *UserController) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "User ID is required", nil)
	}

	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
	}

	user.ID = uuid.MustParse(id)
	user.UpdatedAt = time.Now()

	// Validate fields
	if err := user.ValidateUpdate(); err != nil {
		if e, ok := err.(validation.Errors); ok {
			return utils.ErrorResponse(ctx, http.StatusBadRequest, "Validation failed", e)
		}
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user data", err)
	}

	// Hash password if provided
	if user.Password != "" {
		if err := user.HashPassword(); err != nil {
			return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to process password", err)
		}
	}

	if err := c.service.Update(ctx.Request().Context(), &user); err != nil {
		return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update user", err)
	}

	return utils.SuccessResponse(ctx, http.StatusOK, "User updated successfully", user)
}

func (c *UserController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "User ID is required", nil)
	}

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid ID format", err)
	}

	if err := c.service.Delete(ctx.Request().Context(), parsedID); err != nil {
		return utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete user", err)
	}

	return utils.SuccessResponse(ctx, http.StatusOK, "User deleted successfully", nil)
}
