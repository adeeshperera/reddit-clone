package routes

import (
	"github.com/dfanso/go-echo-boilerplate/internal/controllers"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes registers all application routes
func RegisterRoutes(e *echo.Echo, userController *controllers.UserController) {
	// API group
	api := e.Group("/api")

	// Register all routes
	registerUserRoutes(api, userController)
}

// registerUserRoutes registers user-related routes
func registerUserRoutes(api *echo.Group, userController *controllers.UserController) {

	users := api.Group("/users")
	{
		users.GET("", userController.GetAll)
		users.GET("/:id", userController.GetByID)
		users.POST("", userController.Create)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Delete)
	}
}
