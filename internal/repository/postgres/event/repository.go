package event

import (
	"context"

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
