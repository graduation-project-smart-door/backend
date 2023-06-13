package event

import (
	"context"
	"database/sql"
	"errors"

	"smart-door/internal/apperror"
	"smart-door/internal/domain"
	"smart-door/internal/repository/postgres"

	"github.com/Masterminds/squirrel"
)

const (
	scheme      = "public"
	table       = "events"
	tableScheme = scheme + "." + table
)

type Repository struct {
	client       postgres.Client
	queryBuilder squirrel.StatementBuilderType
}

func NewRepository(client postgres.Client) *Repository {
	return &Repository{client: client, queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}
}

func (repository *Repository) Create(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	newEvent := eventModel{}
	newEvent.FromDomain(*event)

	sql, args, errBuild := repository.queryBuilder.Insert(tableScheme).
		Columns(
			"event_time", "direction", "user_id",
		).
		Values(
			newEvent.EventTime,
			newEvent.Direction,
			newEvent.UserID).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if errBuild != nil {
		return nil, errBuild
	}

	var eventID int
	err := repository.client.QueryRow(sql, args...).Scan(&eventID)
	if err != nil {
		return nil, err
	}

	newEvent.ID = eventID
	return eventModelToDomain(newEvent), nil
}

func (repository *Repository) LastEvent(ctx context.Context, userID int) (*domain.Event, error) {
	query, args, errBuild := repository.queryBuilder.
		Select("id", "direction", "user_id", "event_time").
		Where("user_id = ?", userID).
		OrderBy("event_time DESC").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if errBuild != nil {
		return nil, errBuild
	}

	var model domain.Event
	err := repository.client.QueryRow(query, args...).Scan(
		&model.ID,
		&model.Direction,
		&model.UserID,
		&model.EventTime,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.ErrNotFound
		}
		return nil, err
	}

	return &model, nil
}
