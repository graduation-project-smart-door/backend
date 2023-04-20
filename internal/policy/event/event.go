package event

import (
	"context"

	"smart-door/internal/domain"
)

type Policy struct {
	eventService EventService
}

func NewPolicy(eventService EventService) *Policy {
	return &Policy{eventService: eventService}
}

func (policy *Policy) CreateEvent(ctx context.Context, event domain.Event) (*domain.Event, error) {
	return policy.eventService.CreateEvent(ctx, event)
}
