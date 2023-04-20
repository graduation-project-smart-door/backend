package event

import (
	"time"

	"smart-door/internal/domain"
)

type eventModel struct {
	ID        int       `db:"id"`
	Direction string    `db:"direction"`
	UserID    int       `db:"user_id"`
	EventTime time.Time `db:"event_time"`
}

func (model *eventModel) FromDomain(event domain.Event) {
	model.ID = event.ID
	model.Direction = event.Direction
	model.UserID = event.UserID
	model.EventTime = event.EventTime
}

func eventModelToDomain(event eventModel) *domain.Event {
	return &domain.Event{
		ID:        event.ID,
		Direction: event.Direction,
		UserID:    event.UserID,
		EventTime: event.EventTime,
	}
}
