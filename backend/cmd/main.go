package main

import (
	"log"

	"github.com/dfanso/go-echo-boilerplate/config"
	"github.com/dfanso/go-echo-boilerplate/internal/controllers"
	"github.com/dfanso/go-echo-boilerplate/internal/models"
	"github.com/dfanso/go-echo-boilerplate/internal/repositories"
	"github.com/dfanso/go-echo-boilerplate/internal/routes"
	"github.com/dfanso/go-echo-boilerplate/internal/services"
	"github.com/dfanso/go-echo-boilerplate/pkg/database"

	customMiddleware "github.com/dfanso/go-echo-boilerplate/pkg/middleware"
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

	// Initialize dependencies
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Register routes
	routes.RegisterRoutes(e, userController)

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
