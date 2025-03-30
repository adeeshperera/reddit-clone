package services

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	dto "github.com/dfanso/reddit-clone/internal/dtos"
	"github.com/dfanso/reddit-clone/internal/models"
)

type AuthService struct {
	userService *UserService
}

func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{
		userService: userService,
	}
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (*models.User, error) {
	// Find user by email using UserService
	user, err := s.userService.FindOne(ctx, map[string]any{"email": req.Email})
	if err != nil {
		return nil, err // Return error if database fails
	}
	if user == nil {
		return nil, errors.New("user not found") // No user with this email
	}

	// Check if the provided password matches the stored hashed password
	if err := checkPassword(user.Password, req.Password); err != nil {
		return nil, errors.New("invalid password") // Password doesn't match
	}

	// Login successful, return the user
	return user, nil
}

// checkPassword compares a hashed password with a plaintext password
func checkPassword(hashedPassword, plainPassword string) error {
	// Convert the hashed password string to bytes and compare with the plaintext password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return err // Returns bcrypt.ErrMismatchedHashAndPassword if they don't match
	}
	return nil
}
