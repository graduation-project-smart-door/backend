package user

import (
	"context"

	"smart-door/internal/domain"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	GetByPersonID(ctx context.Context, personID uuid.UUID) (*domain.User, error)
	All(ctx context.Context) ([]*domain.User, error)
}
