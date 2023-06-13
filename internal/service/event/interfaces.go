package event

import (
	"context"

	"smart-door/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, event *domain.Event) (*domain.Event, error)
	LastEvent(ctx context.Context, userID int) (*domain.Event, error)
	All(ctx context.Context) ([]*domain.Event, error)
}
