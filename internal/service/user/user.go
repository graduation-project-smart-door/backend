package user

import (
	"context"

	"smart-door/internal/domain"
	"smart-door/pkg/logging"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service struct {
	logger logging.Logger
	db     Repository
}

func NewService(logger logging.Logger, db Repository) *Service {
	return &Service{logger: logger, db: db}
}

func (service *Service) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	all, errGetAll := service.db.All(ctx)
	if errGetAll != nil {
		service.logger.Error("error get all users from database", zap.Error(errGetAll))
		return nil, errGetAll
	}

	return all, nil
}

func (service *Service) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	newUser, errCreateUser := service.db.Create(ctx, &user)
	if errCreateUser != nil {
		service.logger.Error("error adding user to database", zap.Error(errCreateUser))
		return nil, errCreateUser
	}

	return newUser, nil
}

func (service *Service) GetUserByPersonID(ctx context.Context, personID uuid.UUID) (*domain.User, error) {
	user, errGetUser := service.db.GetByPersonID(ctx, personID)
	if errGetUser != nil {
		service.logger.Error("error get user by person ID in database",
			zap.Error(errGetUser),
			zap.String("personID", personID.String()),
		)
		return nil, errGetUser
	}

	return user, nil
}
