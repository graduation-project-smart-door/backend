package user

import (
	"context"
	"time"

	"smart-door/internal/domain"
	"smart-door/internal/dto"

	"github.com/google/uuid"
)

type Policy struct {
	userService        UserService
	eventService       EventService
	telegramBotService TelegramBotService
	doorService        DoorService
}

func NewPolicy(
	userService UserService,
	eventService EventService,
	telegramBotService TelegramBotService,
	doorService DoorService) *Policy {
	return &Policy{
		userService:        userService,
		eventService:       eventService,
		telegramBotService: telegramBotService,
		doorService:        doorService,
	}
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
	newEvent, errCreateEvent := policy.eventService.CreateEvent(ctx, event)
	if errCreateEvent != nil {
		return newEvent, errCreateEvent
	}

	_, errOpenDoor := policy.doorService.Open()
	if errOpenDoor != nil {
		return newEvent, errOpenDoor
	}

	lastEvent, errGetLastEvent := policy.eventService.GetLastEventByUser(ctx, user.ID)
	if errGetLastEvent != nil {
		return nil, errOpenDoor
	}

	minimumInterval := time.Now().Add(-12 * time.Hour)
	if lastEvent.EventTime.Before(minimumInterval) {
		_, errSendNotification := policy.telegramBotService.SendNotification(
			dto.NewEventNotification(
				user.FirstName,
				user.LastName,
				event.Direction,
			),
		)

		if errSendNotification != nil {
			return newEvent, errSendNotification
		}
	}

	return newEvent, nil
}
