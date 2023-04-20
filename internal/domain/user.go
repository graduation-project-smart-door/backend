package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         int       `json:"id,omitempty"`
	PersonID   uuid.UUID `json:"person_id,omitempty"`
	Email      string    `json:"email,omitempty"`
	FirstName  string    `json:"first_name,omitempty"`
	Patronymic string    `json:"patronymic,omitempty"`
	LastName   string    `json:"last_name,omitempty"`
	Role       string    `json:"role,omitempty"`
	Phone      string    `json:"phone,omitempty"`
	Password   string    `json:"-"`
	Avatar     string    `json:"avatar,omitempty"`
	Position   string    `json:"position"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
