package user

import (
	"context"

	"smart-door/internal/domain"

	"github.com/google/uuid"
)

type Policy interface {
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
	CreateEvent(ctx context.Context, event domain.Event, personID uuid.UUID) (*domain.Event, error)
}

type EventBroker interface {
	ToMessage(message domain.Event) error
}
