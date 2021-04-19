package user

import (
	"context"

	"github.com/vegggg/user-management/entity"
)

type UserService interface {
	GetUser(ctx context.Context, phone string) (*entity.UserProfile, error)
	CreateUser(ctx context.Context, user *entity.UserProfile) (*entity.UserProfile, error)
}
