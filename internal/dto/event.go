package dto

import (
	"time"

	"smart-door/internal/domain"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateEvent struct {
	UserID    int       `json:"user_id"`
	Direction string    `json:"direction"`
	EventType string    `json:"event_type"`
	EventTime time.Time `json:"event_time"`
}

func (event *CreateEvent) Validate() error {
	return validation.ValidateStruct(event,
		validation.Field(&event.UserID, validation.Required),
		validation.Field(&event.Direction, validation.Required),
		validation.Field(&event.EventTime, validation.Required),
		validation.Field(&event.EventType, validation.Required))
}

func (event *CreateEvent) ToDomain() domain.Event {
	return domain.Event{
		EventType: event.EventType,
		UserID:    event.UserID,
		Direction: event.Direction,
		EventTime: event.EventTime,
	}
}
