package handler

import (
	"context"

	"github.com/golang/glog"
	user_managementpb "github.com/vegggg/user-management/proto/user_mgnt/v1"
)

func (u UserManagement) GetUser(ctx context.Context, in *user_managementpb.GetUserRequest) (*user_managementpb.UserProfile, error) {
	glog.Info("[GetUser]", "phone ", in.GetPhone())
	user, err := u.user.GetUser(ctx, in.GetPhone())
	if err != nil {
		glog.Info("[GetUser]", "error ", err)
		return nil, err
	}
	return (*user_managementpb.UserProfile)(user), nil
}
