package dtos

import (
	"regexp"

	"github.com/dfanso/reddit-clone/internal/models"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9]+(_[a-zA-Z0-9]+)*$`)

type RegisterRequest struct {
	Email    string `json:"email"`    // User's email address
	Username string `json:"username"` // User's chosen username
	Password string `json:"password"` // User's password
}

func (r RegisterRequest) Validate() error {
	return validation.ValidateStruct(&r,
		// Email: required, valid email format, 5-100 characters
		validation.Field(&r.Email, validation.Required, validation.Length(5, 100), is.Email),
		// Username: required, 3-20 characters
		validation.Field(&r.Username, validation.Required, validation.Length(3, 20), validation.Match(usernameRegex)),
		// Password: required, 8-72 characters
		validation.Field(&r.Password, validation.Required, validation.Length(8, 72)),
	)
}

// LoginRequest defines the structure for login input data
type LoginRequest struct {
	Email    string `json:"email"`    // User's email address
	Password string `json:"password"` // User's password
}

// Validate validates the LoginRequest fields
func (r LoginRequest) Validate() error {
	return validation.ValidateStruct(&r,
		// Email: required, valid email format, 5-100 characters
		validation.Field(&r.Email, validation.Required, validation.Length(5, 100), is.Email),
		// Password: required, 8-72 characters
		validation.Field(&r.Password, validation.Required, validation.Length(8, 72)),
	)
}

// LoginResponse defines the structure for login response data
type LoginResponse struct {
	User  *models.User `json:"user"`  // User data
	Token string       `json:"token"` // JWT token
}
