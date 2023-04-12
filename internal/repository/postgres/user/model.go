package user

import (
	"database/sql"
	"time"

	"smart-door/internal/domain"
)

type userModel struct {
	ID         int            `db:"id"`
	PersonID   string         `db:"person_id"`
	Email      sql.NullString `db:"email"`
	FirstName  string         `db:"first_name"`
	Patronymic sql.NullString `db:"patronymic"`
	LastName   string         `db:"last_name"`
	Role       string         `db:"role"`
	Phone      sql.NullString `db:"phone"`
	Password   sql.NullString `db:"password"`
	Avatar     sql.NullString `db:"avatar"`
	Position   string         `db:"position"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (model *userModel) FromDomain(user domain.User) {
	// Обязательные поля
	model.ID = user.ID
	model.PersonID = user.PersonID
	model.FirstName = user.FirstName
	model.LastName = user.LastName
	model.Role = user.Role
	model.Position = user.Position
	model.CreatedAt = user.CreatedAt
	model.UpdatedAt = user.UpdatedAt

	// Необязательные
	model.Phone = model.stringToNullString(user.Phone)
	model.Password = model.stringToNullString(user.Password)
	model.Avatar = model.stringToNullString(user.Avatar)
	model.Patronymic = model.stringToNullString(user.Patronymic)
	model.Email = model.stringToNullString(user.Email)
}

func (model *userModel) stringToNullString(value string) sql.NullString {
	if value != "" {
		return sql.NullString{String: value, Valid: true}
	}

	return sql.NullString{}
}

func userModelToDomain(user userModel) *domain.User {
	domainUser := &domain.User{
		ID:         user.ID,
		PersonID:   user.PersonID,
		Email:      "",
		FirstName:  user.FirstName,
		Patronymic: "",
		LastName:   user.LastName,
		Role:       user.Role,
		Phone:      "",
		Password:   "",
		Avatar:     "",
		Position:   user.Position,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}

	if user.Email.Valid {
		domainUser.Email = user.Email.String
	}

	if user.Patronymic.Valid {
		domainUser.Patronymic = user.Patronymic.String
	}

	if user.Phone.Valid {
		domainUser.Phone = user.Phone.String
	}

	if user.Password.Valid {
		domainUser.Password = user.Password.String
	}

	if user.Avatar.Valid {
		domainUser.Avatar = user.Avatar.String
	}

	return domainUser
}
