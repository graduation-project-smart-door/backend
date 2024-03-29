package user

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"smart-door/internal/apperror"
	"smart-door/internal/domain"
	"smart-door/internal/repository/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
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

func (repository *Repository) All(ctx context.Context) ([]*domain.User, error) {
	sql, args, errBuild := repository.queryBuilder.Select("id", "person_id", "email", "first_name", "patronymic", "last_name",
		"role", "phone", "password", "avatar", "position").From(tableScheme).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if errBuild != nil {
		return nil, errBuild
	}

	rows, err := repository.client.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var list []*domain.User
	for rows.Next() {
		user := domain.User{}
		if err = rows.Scan(
			&user.ID,
			&user.PersonID,
			&user.Email,
			&user.FirstName,
			&user.Patronymic,
			&user.LastName,
			&user.Role,
			&user.Position,
			&user.Password,
			&user.Avatar,
			&user.Position,
		); err != nil {
			return nil, err
		}

		list = append(list, &user)
	}

	return list, nil
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
			newUser.PersonID,
			newUser.Email,
			newUser.FirstName,
			newUser.Patronymic,
			newUser.LastName, newUser.Role,
			newUser.Phone, newUser.Password,
			newUser.Avatar,
			newUser.Position).
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

func (repository *Repository) GetByPersonID(ctx context.Context, personID uuid.UUID) (*domain.User, error) {
	query, args, errToSQL := repository.queryBuilder.Select("id").
		Columns("person_id", "email", "first_name", "patronymic", "last_name",
			"role", "phone", "password", "avatar", "position").
		From(tableScheme).
		Where(squirrel.Eq{"person_id": personID}).ToSql()

	if errToSQL != nil {
		return nil, errToSQL
	}

	var user userModel
	err := repository.client.QueryRow(query, args...).Scan(
		&user.ID,
		&user.PersonID,
		&user.Email,
		&user.FirstName,
		&user.Patronymic,
		&user.LastName,
		&user.Role,
		&user.Phone,
		&user.Password,
		&user.Avatar,
		&user.Position,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.ErrNotFound
		}
		return nil, err
	}

	return userModelToDomain(user), nil
}
