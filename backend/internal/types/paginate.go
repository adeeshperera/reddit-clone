package types

import "github.com/dfanso/go-echo-boilerplate/internal/models"

type UserPaginationResult struct {
	Users      []models.User
	Total      int64
	Page       int
	Limit      int
	TotalPages int
}
