package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string
type Status string

const (
	// Roles
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"

	// Statuses
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusBanned   Status = "banned"

	// Password constraints
	MinPasswordLength = 8
	MaxPasswordLength = 72
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string         `json:"name" validate:"required,min=2,max=50" gorm:"not null"`
	Email     string         `json:"email" validate:"required,email" gorm:"uniqueIndex;not null"`
	Password  string         `json:"password,omitempty" validate:"required,min=8,max=72" gorm:"not null"`
	Role      Role           `json:"role" validate:"required,oneof=admin user" gorm:"type:varchar(20);not null;default:'user'"`
	Status    Status         `json:"status" validate:"required,oneof=active inactive banned" gorm:"type:varchar(20);not null;default:'active'"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Create a singleton validator instance
var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (u *User) Validate() error {
	return validate.Struct(u)
}

func (u *User) ValidateUpdate() error {
	if u.Password != "" {
		return validate.StructPartial(u, "Name", "Email", "Password", "Role", "Status")
	}
	return validate.StructPartial(u, "Name", "Email", "Role", "Status")
}

// Password handling methods
func (u *User) HashPassword() error {
	if len(u.Password) == 0 {
		return validation.NewError("validation_error", "password cannot be empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// GORM Hooks
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// Set timestamps
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now

	// Set defaults if empty
	if u.Role == "" {
		u.Role = RoleUser
	}
	if u.Status == "" {
		u.Status = StatusActive
	}

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
