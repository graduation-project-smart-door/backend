package user

import (
	"context"

	"smart-door/internal/domain"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
	GetUserByPersonID(ctx context.Context, personID uuid.UUID) (*domain.User, error)
}

type EventService interface {
	CreateEvent(ctx context.Context, event domain.Event) (*domain.Event, error)
}
