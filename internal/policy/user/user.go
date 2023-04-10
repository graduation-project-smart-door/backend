package user

import (
	"context"
	"smart-door/internal/domain"
)

type Policy struct {
	userService UserService
}

func NewPolicy(userService UserService) *Policy {
	return &Policy{userService: userService}
}

func (policy *Policy) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	return policy.userService.CreateUser(ctx, user)
}
