package event

import (
	"context"

	"smart-door/internal/domain"
)

type Policy interface {
	CreateEvent(ctx context.Context, event domain.Event) (*domain.Event, error)
}
