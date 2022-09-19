package postgres

import (
	"context"
	"github.com/oklog/ulid/v2"
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

func (u *UsersRepository) CreateUser(ctx context.Context, user domain.User, password string) error {
	model := userModel{}
	model.FromDomain(user)
	model.ID = ulid.Make().String()
	model.EncryptedPassword = password
	_, err := u.db.NewInsert().Model(&model).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
