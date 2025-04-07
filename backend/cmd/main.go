package main

import (
    "fmt"
    "log"

    "github.com/dfanso/reddit-clone/pkg/auth"
    "github.com/google/uuid"
)

func main() {
    // Initialize JWTManager
    jwtManager, err := auth.NewJWTManager()
    if err != nil {
        log.Fatalf("Failed to initialize JWTManager: %v", err)
    }

    // Generate a token
    userID := uuid.New()
    role := "user"
    token, err := jwtManager.GenerateToken(userID, role)
    if err != nil {
        log.Fatalf("Failed to generate token: %v", err)
    }
    fmt.Println("Generated Token:", token)

    // Validate the token
    claims, err := jwtManager.ValidateToken(token)
    if err != nil {
        log.Fatalf("Failed to validate token: %v", err)
    }
    fmt.Printf("Validated Claims: UserID=%s, Role=%s\n", claims.UserID, claims.Role)
}
