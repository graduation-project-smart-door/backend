package user

import (
	"context"
	"smart-door/internal/domain"
	"smart-door/internal/repository/postgres"
	"time"

	"github.com/Masterminds/squirrel"
)

const (
	scheme      = "public"
	table       = "users"
	tableScheme = scheme + "." + table
)

type Repository struct {
	client       postgres.Client
	queryBuilder squirrel.StatementBuilderType
}

func NewRepository(client postgres.Client) *Repository {
	return &Repository{client: client, queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}
}

func (repository *Repository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	newUser := userModel{}
	newUser.FromDomain(*user)

	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	sql, args, errBuild := repository.queryBuilder.Insert(tableScheme).
		Columns(
			"person_id", "email", "first_name", "patronymic", "last_name",
			"role", "phone", "password", "avatar", "position",
		).
		Values(
			user.PersonID,
			user.Email,
			user.FirstName,
			user.Patronymic,
			user.LastName, user.Role,
			user.Phone, user.Password,
			user.Avatar,
			user.Position).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if errBuild != nil {
		return nil, errBuild
	}

	var userID int
	err := repository.client.QueryRow(sql, args...).Scan(&userID)
	if err != nil {
		return nil, err
	}

	newUser.ID = userID
	return userModelToDomain(newUser), nil
}
