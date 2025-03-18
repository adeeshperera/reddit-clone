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
type Stage string

const (
	// Roles
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"

	// Statuses
	StatusVerified   Status = "verified"
	StatusUnverified Status = "unverified"
	StatusBanned     Status = "banned"

	// Stages
	StageEmailVerification Stage = "email_verification"
	StageEmailVerified     Stage = "email_verified"
	StageGoogleSSO         Stage = "google_sso"
	StageCompleted         Stage = "completed"

	// Password constraints
	MinPasswordLength = 8
	MaxPasswordLength = 72
)

type User struct {
	ID           uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Handler      string         `json:"handler" validate:"required,min=3,max=20,matches=^[a-zA-Z0-9]+(_[a-zA-Z0-9]+)*$" gorm:"uniqueIndex;not null"`
	Name         string         `json:"name" validate:"required,min=2,max=50" gorm:"not null"`
	Email        string         `json:"email" validate:"required,email" gorm:"uniqueIndex;not null"`
	Password     string         `json:"password,omitempty" validate:"required,min=8,max=72" gorm:"not null"`
	Role         Role           `json:"role" validate:"required,oneof=admin user" gorm:"type:varchar(20);not null;default:'user'"`
	Status       Status         `json:"status" validate:"required,oneof=verified unverified banned" gorm:"type:varchar(20);not null;default:'unverified'"`
	Stage        Stage          `json:"stage" validate:"required,oneof=email_verification email_verified google_sso completed" gorm:"type:varchar(20);not null;default:'email_verification'"`
	Avatar       string         `json:"avatar" gorm:"type:varchar(255)"`
	Banner       string         `json:"banner" gorm:"type:varchar(255)"`
	Description  string         `json:"description" gorm:"type:text"`
	PostKarma    int            `json:"postKarma" gorm:"default:0"`
	CommentKarma int            `json:"commentKarma" gorm:"default:0"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
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
		return validate.StructPartial(u, "Name", "Email", "Password", "Role", "Status", "Handler", "Stage")
	}
	return validate.StructPartial(u, "Name", "Email", "Role", "Status", "Handler", "Stage")
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
		u.Status = StatusUnverified
	}

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
