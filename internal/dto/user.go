package dto

import (
	"time"

	"smart-door/internal/domain"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type CreateUser struct {
	PersonID   uuid.UUID `json:"person_id,omitempty"`
	FirstName  string    `json:"first_name,omitempty"`
	Patronymic string    `json:"patronymic,omitempty"`
	LastName   string    `json:"last_name,omitempty"`
	Position   string    `json:"position"`
}

type RecognizeUser struct {
	PersonID  uuid.UUID `json:"person_id"`
	Direction string    `json:"direction"`
	EventTime time.Time `json:"event_time"`
}

func (user *CreateUser) Validate() error {
	return validation.ValidateStruct(user,
		validation.Field(&user.PersonID, validation.Required),
		validation.Field(&user.FirstName, validation.Required),
		validation.Field(&user.LastName, validation.Required),
		validation.Field(&user.Position, validation.Required),
	)
}

func (user *CreateUser) ToDomain() domain.User {
	return domain.User{
		PersonID:   user.PersonID,
		FirstName:  user.FirstName,
		Patronymic: user.Patronymic,
		LastName:   user.LastName,
		Position:   user.Position,
		Role:       domain.UserRole,
	}
}

func (user *RecognizeUser) Validate() error {
	return validation.ValidateStruct(user,
		validation.Field(&user.PersonID, validation.Required),
		validation.Field(&user.Direction, validation.In("exit", "enter")),
		validation.Field(&user.EventTime, validation.Required))
}

func (user *RecognizeUser) ToDomain() domain.Event {
	return domain.Event{
		Direction: user.Direction,
		EventTime: user.EventTime,
	}
}
