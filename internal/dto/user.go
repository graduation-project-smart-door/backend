package dto

import (
	"github.com/go-ozzo/ozzo-validation"
	"smart-door/internal/domain"
)

type CreateUser struct {
	PersonID   string `json:"person_id,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	Patronymic string `json:"patronymic,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Position   string `json:"position"`
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
		Role:       "user",
	}
}
