package routes

import (
	"github.com/dfanso/reddit-clone/internal/controllers"
	"github.com/dfanso/reddit-clone/pkg/auth"
	"github.com/dfanso/reddit-clone/pkg/middleware"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes registers all application routes
func RegisterRoutes(e *echo.Echo, userController *controllers.UserController, authController *controllers.AuthController, jwtManager *auth.JWTManager) {
	// API group
	api := e.Group("/api/v1")

	// Register public routes (no authentication required)
	PublicRoutes(api, authController)

	// Register private routes (authentication required)
	PrivateRoutes(api, authController, userController, jwtManager)
}

// PublicRoutes registers routes that don't require authentication
func PublicRoutes(api *echo.Group, authController *controllers.AuthController) {
	// Auth routes - public
	auth := api.Group("/auth")
	auth.POST("/register", authController.Register)
	auth.POST("/login", authController.Login)

	// Add other public routes here as needed
	// For example: auth.POST("/forgot-password", authController.ForgotPassword)
	// auth.POST("/reset-password", authController.ResetPassword)
}

// PrivateRoutes registers routes that require authentication
func PrivateRoutes(api *echo.Group, authController *controllers.AuthController, userController *controllers.UserController, jwtManager *auth.JWTManager) {
	// Create protected group with JWT middleware
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(jwtManager))

	//auth routes - protected
	// auth := protected.Group("/auth")
	// {
	// 	auth.GET("/profile", authController.GetProfile)
	// }

	// User routes - protected
	users := protected.Group("/users")
	{
		users.GET("", userController.GetAll)
		users.GET("/paginated", userController.GetPaginated)
		users.GET("/:id", userController.GetByID)
		users.POST("", userController.Create)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Delete)
	}

	// Add other protected routes here as needed
	// For example: protected.GET("/profile", userController.GetProfile)
	// protected.PUT("/profile", userController.UpdateProfile)
}
