package users

import (
	"context"
	"smart-door/app/internal/domain"
)

type Database interface {
	GetAllUsers(ctx context.Context) ([]domain.User, error)
}
