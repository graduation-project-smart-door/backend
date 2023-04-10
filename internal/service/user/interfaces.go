package user

import (
	"context"
	"smart-door/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
}
