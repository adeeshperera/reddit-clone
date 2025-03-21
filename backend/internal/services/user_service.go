package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"

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

func (s *UserService) Create(ctx context.Context, user *models.User) error {
	return s.repo.Create(ctx, user)
}

func (s *UserService) Update(ctx context.Context, user *models.User) error {
	return s.repo.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

// RegisterUser creates a new user with the provided details and returns the user with formatted username
func (s *UserService) RegisterUser(ctx context.Context, email, username, password string) (*models.User, error) {
	// Create user model
	user := &models.User{
		Email:    email,
		Handler:  username,
		Name:     username, // Using username as initial name
		Password: password,
		Role:     models.RoleUser,
		Status:   models.StatusUnverified,
		Stage:    models.StageEmailVerification,
	}

	// Hash the password
	if err := user.HashPassword(); err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Save the user to the database
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Format the username with the "u/" prefix for the frontend
	user.Handler = fmt.Sprintf("u/%s", user.Handler)

	return user, nil
}
