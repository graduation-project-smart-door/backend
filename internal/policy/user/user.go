package user

import (
	"context"

	"smart-door/internal/domain"

	"github.com/google/uuid"
)

type Policy struct {
	userService  UserService
	eventService EventService
}

func NewPolicy(userService UserService, eventService EventService) *Policy {
	return &Policy{userService: userService, eventService: eventService}
}

func (policy *Policy) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	return policy.userService.CreateUser(ctx, user)
}

func (policy *Policy) CreateEvent(ctx context.Context, event domain.Event, personID uuid.UUID) (*domain.Event, error) {
	user, errGetUser := policy.userService.GetUserByPersonID(ctx, personID)
	if errGetUser != nil {
		return nil, errGetUser
	}

	event.UserID = user.ID
	return policy.eventService.CreateEvent(ctx, event)
}
