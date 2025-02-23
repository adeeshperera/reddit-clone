# Go Echo Boilerplate

A production-ready boilerplate for building RESTful APIs using Go Echo framework with PostgreSQL and GORM.

## Features

- **Echo Framework**: High performance, minimalist Go web framework
- **PostgreSQL & GORM**: Using GORM as ORM with PostgreSQL
- **UUID Support**: UUID primary keys for better data distribution
- **Clean Architecture**: Follows clean architecture principles with proper separation of concerns
- **Environment Management**: Multiple environment support with proper configuration management
- **Structured Logging**: Custom logging middleware with colored output
- **API Versioning**: Built-in support for API versioning
- **Input Validation**: Request validation using go-playground/validator
- **Error Handling**: Centralized error handling with proper HTTP status codes
- **Hot Reload**: Support for hot reload during development using Air

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
   git clone https://github.com/yourusername/go-echo-boilerplate.git
   cd go-echo-boilerplate
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   ```bash
   cp .env.development .env
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

### User Routes
- `GET /api/users` - Get all users
- `GET /api/users/:id` - Get user by ID
- `POST /api/users` - Create new user
- `PUT /api/users/:id` - Update user
- `DELETE /api/users/:id` - Delete user

### Health Check
- `GET /health` - Service health check

## Request/Response Examples

### Create User
```bash
# Request
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepass123",
    "role": "user",
    "status": "active"
  }'

# Response
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user",
    "status": "active",
    "created_at": "2025-01-04T10:30:00Z",
    "updated_at": "2025-01-04T10:30:00Z"
  }
```

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