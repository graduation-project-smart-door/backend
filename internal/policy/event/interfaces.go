package event

import (
	"context"

	"smart-door/internal/domain"
)

type EventService interface {
	CreateEvent(ctx context.Context, event domain.Event) (*domain.Event, error)
	GetAllEvents(ctx context.Context) ([]*domain.Event, error)
}
