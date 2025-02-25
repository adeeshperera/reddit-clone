# Go Backend

## Project Structure

```
.
├── cmd/
│   └── main.go                 # Application entry point
├── config/
│   └── config.go              # Configuration management
├── internal/
│   ├── controllers/           # Request handlers
│   ├── models/                # Data models
│   ├── repositories/          # Data access layer
│   ├── routes/                # Route definitions
│   └── services/             # Business logic
    └── types/                # Types and interfaces

├── pkg/
│   ├── database/             # Database connections
│   ├── middleware/           # Custom middleware
│   └── utils/                # Utility functions
├── .air.toml                 # Air configuration for hot reload
├── .env.development          # Development environment variables
├── .env.production          # Production environment variables
├── go.mod                    # Go module file
└── README.md                # Project documentation
```

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Air (for hot reload)

## Getting Started

1. **Clone the repository**
   ```bash
   git clone https://github.com/dfanso/go-echo-boilerplate.git
   cd go-echo-boilerplate
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   ```
   Edit the `.env` file with your configuration.

4. **Install Air for hot reload**
   ```bash
   go install github.com/air-verse/air@latest
   ```

5. **Run the application**
   
   Development mode with hot reload:
   ```bash
   air
   ```

   Production mode:
   ```bash
   go run cmd/main.go
   ```

## Environment Variables

```env
# Server Configuration
SERVER_PORT=8080

# PostgreSQL Configuration
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=myapp
```

## API Endpoints

### Import Postman Collection

Import the `go-echo-boilerplate.postman_collection.json` file into Postman to test the API endpoints.

## Development

### Adding New Routes

1. Create a new controller in `internal/controllers/`
2. Add corresponding service and repository if needed
3. Register routes in `internal/routes/`

Example:
```go
// internal/routes/routes.go
func RegisterRoutes(e *echo.Echo, controller *controllers.UserController) {
    api := e.Group("/api")
    api.GET("/users", controller.GetAll)
    api.POST("/users", controller.Create)
}
```

### Error Handling

The boilerplate includes a centralized error handling system:

```go
if err != nil {
    return utils.ErrorResponse(ctx, http.StatusBadRequest, "Error message", err)
}
return utils.SuccessResponse(ctx, http.StatusOK, "Success message", data)
```

## Production Deployment

1. Build the application:
   ```bash
   go build -o app cmd/main.go
   ```

2. Set up production environment:
   ```bash
   cp .env.production .env
   ```

3. Run the application:
   ```bash
   ./app
   ```

## License

This project is licensed under the MIT License - see the LICENSE file for details