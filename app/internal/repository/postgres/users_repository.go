package postgres

import (
	"context"
	"github.com/uptrace/bun"
	"smart-door/app/internal/domain"
)

type UsersRepository struct {
	db *bun.DB
}

func NewUsersRepository(db *bun.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (u *UsersRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	usersModel := make([]userModel, 0)
	err := u.db.NewSelect().Model(&usersModel).Scan(ctx)
	if err != nil {
		return nil, err
	}
	users := userModelsToDomain(usersModel)
	return users, nil
}
