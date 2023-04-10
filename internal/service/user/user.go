package user

import (
	"context"
	"smart-door/internal/domain"

	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
	db     Repository
}

func NewService(logger *zap.Logger, db Repository) *Service {
	return &Service{logger: logger, db: db}
}

func (service *Service) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	newBench, errCreateBench := service.db.Create(ctx, &user)
	if errCreateBench != nil {
		service.logger.Error("error adding user to database", zap.Error(errCreateBench))
	}

	return newBench, nil
}
