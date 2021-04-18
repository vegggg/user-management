package handler

import (
	"context"

	"github.com/golang/glog"
	user_managementpb "github.com/vegggg/user-management/proto/user_mgnt/v1"
)

type UserManagement struct{}

func (u UserManagement) GetUser(ctx context.Context, in *user_managementpb.GetUserRequest) (*user_managementpb.UserProfile, error) {
	glog.Info("[GetUser]", "user_id ", in.GetId())

	return nil, nil
}
func (u UserManagement) CreateUser(ctx context.Context, in *user_managementpb.CreateUserRequest) (*user_managementpb.UserProfile, error) {
	glog.Info("[CreateUser]", "user_id ", in.User.GetPhone())

	return nil, nil
}
func NewUserManagement() *UserManagement {
	return new(UserManagement)
}
