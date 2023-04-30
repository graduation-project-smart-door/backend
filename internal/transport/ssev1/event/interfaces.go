package event

import (
	"context"

	"smart-door/internal/domain"
)

type EventPolicy interface {
	CreateEvent(ctx context.Context, event domain.Event) (*domain.Event, error)
}
