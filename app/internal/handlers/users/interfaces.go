package users

import (
	"context"
	"smart-door/app/internal/domain"
)

type Service interface {
	GetListUsers(ctx context.Context) ([]domain.User, error)
	CreateUser(ctx context.Context, user domain.User) error
}
