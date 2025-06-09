package main

import (
	"log"

	"github.com/dfanso/reddit-clone/config"
	"github.com/dfanso/reddit-clone/internal/controllers"
	"github.com/dfanso/reddit-clone/internal/models"
	"github.com/dfanso/reddit-clone/internal/repositories"
	"github.com/dfanso/reddit-clone/internal/routes"
	"github.com/dfanso/reddit-clone/internal/services"
	"github.com/dfanso/reddit-clone/pkg/auth"
	"github.com/dfanso/reddit-clone/pkg/database"

	customMiddleware "github.com/dfanso/reddit-clone/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize PostgreSQL
	db, err := database.NewPostgresClient(
		cfg.Postgres.Host,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.Port,
	)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	// Enable UUID extension if it doesn't exist
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	// Auto Migrate the schema with GORM
	err = db.AutoMigrate(
		&models.User{}, // Add other models here as needed
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize Echo
	e := echo.New()
	e.HideBanner = false // Show the Echo banner
	e.HidePort = false   // Show the port number

	// Middleware
	e.Use(customMiddleware.NewCustomLogger().Middleware())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize JWT manager
	jwtManager, err := auth.NewJWTManager()
	if err != nil {
		log.Fatalf("Failed to initialize JWT manager: %v", err)
	}

	// Initialize dependencies
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userService)
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(userService, authService, jwtManager)

	// Register routes
	routes.RegisterRoutes(e, userController, authController, jwtManager, userService)

	// health check route
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status": "OK",
		})
	})

	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}
