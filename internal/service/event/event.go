package event

import (
	"context"

	"smart-door/internal/domain"
	"smart-door/pkg/logging"

	"go.uber.org/zap"
)

type Service struct {
	logger logging.Logger
	db     Repository
}

func NewService(logger logging.Logger, db Repository) *Service {
	return &Service{logger: logger, db: db}
}

func (service *Service) CreateEvent(ctx context.Context, event domain.Event) (*domain.Event, error) {
	newEvent, errCreateEvent := service.db.Create(ctx, &event)
	if errCreateEvent != nil {
		service.logger.Error("error adding event to database", zap.Error(errCreateEvent))
	}

	return newEvent, nil
}

func (service *Service) GetLastEventByUser(ctx context.Context, userID int) (*domain.Event, error) {
	return service.db.LastEvent(ctx, userID)
}

func (service *Service) GetAllEvents(ctx context.Context) ([]*domain.Event, error) {
	return service.db.All(ctx)
}
