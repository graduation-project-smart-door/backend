package user

import (
	"context"
	"smart-door/internal/domain"
)

type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
}
