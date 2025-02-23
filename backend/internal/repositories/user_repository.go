package repositories

import (
	"context"
	"errors"
	"math"

	"github.com/dfanso/go-echo-boilerplate/internal/models"
	"github.com/dfanso/go-echo-boilerplate/internal/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
	types.PaginationResult
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindOne(ctx context.Context, query interface{}) (*models.User, error) {
	var user models.User
	result := r.db.WithContext(ctx).Where(query).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	result := r.db.WithContext(ctx).Find(&users)
	for i := range users {
		users[i].Password = ""
	}
	return users, result.Error
}

func (r *UserRepository) FindPaginated(ctx context.Context, query interface{}, page int, limit int) (*types.PaginationResult, error) {

	// Validate page (minimum 1)
	if page < 1 {
		page = 1
	}
	// Validate limit (default to 10, max 100)
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	// Get total count of users matching the query
	var total int64
	db := r.db.WithContext(ctx).Model(&models.User{})
	if query != nil {
		db = db.Where(query) // Apply filter if provided
	}
	err := db.Count(&total).Error
	if err != nil {
		return nil, err
	}

	// Fetch paginated users
	var users []models.User
	offset := (page - 1) * limit // Calculate offset
	db = r.db.WithContext(ctx).Offset(offset).Limit(limit)
	if query != nil {
		db = db.Where(query) // Apply filter to the main query
	}
	err = db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	for i := range users {
		users[i].Password = ""
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	// Return the result
	return &types.PaginationResult{
		Data:       users,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User
	result := r.db.WithContext(ctx).First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, "id = ?", id).Error
}
