package users

import (
	"context"
	"go.uber.org/zap"
	"smart-door/app/internal/domain"
)

type Service struct {
	db  Database
	log *zap.Logger
}

func NewService(db Database, log *zap.Logger) *Service {
	return &Service{db: db, log: log}
}

func (s *Service) GetListUsers(ctx context.Context) ([]domain.User, error) {
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		s.log.Error("error get all users", zap.Error(err))
		return nil, err
	}
	return users, nil
}
