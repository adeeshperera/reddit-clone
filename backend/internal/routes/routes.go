package routes

import (
	"github.com/dfanso/reddit-clone/internal/controllers"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes registers all application routes
func RegisterRoutes(e *echo.Echo, userController *controllers.UserController, authController *controllers.AuthController) {
	// API group
	api := e.Group("/api")

	// Register all routes
	registerUserRoutes(api, userController)
	registerAuthRoutes(api, authController)
}

// registerAuthRoutes to map /api/auth/register to the Register method
func registerAuthRoutes(api *echo.Group, authController *controllers.AuthController) {
	auth := api.Group("/auth")
	auth.POST("/register", authController.Register)
}

// registerUserRoutes registers user-related routes
func registerUserRoutes(api *echo.Group, userController *controllers.UserController) {

	users := api.Group("/users")
	{
		users.GET("", userController.GetAll)
		users.GET("/paginated", userController.GetPaginated)
		users.GET("/:id", userController.GetByID)
		users.POST("", userController.Create)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Delete)
	}
}
