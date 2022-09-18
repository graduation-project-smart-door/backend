package users

import (
	"context"
	"smart-door/app/internal/domain"
)

type Database interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
}
