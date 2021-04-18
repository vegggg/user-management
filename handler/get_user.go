package handler

import (
	"context"

	"github.com/golang/glog"
	user_managementpb "github.com/vegggg/user-management/proto/user_mgnt/v1"
)

func (u UserManagement) GetUser(ctx context.Context, in *user_managementpb.GetUserRequest) (*user_managementpb.UserProfile, error) {
	glog.Info("[GetUser]", "user_id ", in.GetId())
	user, err := u.sqlrepo.GetUser(ctx, in.GetPhoneNumber())
	if err != nil {
		glog.Info("[GetUser]", "error ", err)
		return nil, err
	}
	return (*user_managementpb.UserProfile)(user), nil
}
