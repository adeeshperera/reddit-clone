package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	dto "github.com/dfanso/reddit-clone/internal/dtos"
	"github.com/dfanso/reddit-clone/internal/models"
	"github.com/dfanso/reddit-clone/internal/repositories"
	"github.com/dfanso/reddit-clone/internal/types"
)

type UserService struct {
	repo *repositories.UserRepository
	types.UserPaginationResult
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) FindOne(ctx context.Context, filter interface{}) (*models.User, error) {
	return s.repo.FindOne(ctx, filter)
}

func (s *UserService) GetAll(ctx context.Context) ([]models.User, error) {
	return s.repo.FindAll(ctx)
}

func (s *UserService) FindPaginated(ctx context.Context, query interface{}, page int, limit int) (*types.UserPaginationResult, error) {
	return s.repo.FindPaginated(ctx, query, page, limit)
}

func (s *UserService) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *UserService) Create(ctx context.Context, user *models.User) (*models.User, error) {
	return s.repo.Create(ctx, user)
}

func (s *UserService) Update(ctx context.Context, user *models.User) error {
	return s.repo.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

// RegisterUser creates a new user with the provided details and returns the user with formatted username
func (s *UserService) RegisterUser(ctx context.Context, req dto.RegisterRequest) (*models.User, error) {
	// Create user model
	user := &models.User{
		Email:    req.Email,
		Handler:  req.Username,
		Name:     req.Username, // Using username as initial name
		Password: req.Password,
		Role:     models.RoleUser,
		Status:   models.StatusUnverified,
		Stage:    models.StageEmailVerification,
	}

	// Check if email already exists
	filter := map[string]any{"email": req.Email}
	existingUser, err := s.FindOne(ctx, filter)
	if err != nil {
		return nil, errors.New("Error checking existing user")
	}
	if existingUser != nil {
		return nil, errors.New("email already in use")
	}

	// Check if username already exists
	filter = map[string]any{"handler": req.Username}
	existingUser, err = s.FindOne(ctx, filter)
	if err != nil {
		return nil, errors.New("Error checking existing username")
	}
	if existingUser != nil {
		return nil, errors.New("Username already taken")
	}

	// Hash the password
	if err := user.HashPassword(); err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Save the user to the database
	createdUser, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	// Format the username with the "u/" prefix for the frontend
	createdUser.Handler = fmt.Sprintf("u/%s", createdUser.Handler)

	return createdUser, nil
}
