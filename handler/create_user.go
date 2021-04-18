package handler

import (
	"context"

	"github.com/golang/glog"
	"github.com/vegggg/user-management/entity"
	user_managementpb "github.com/vegggg/user-management/proto/user_mgnt/v1"
)

func (u UserManagement) CreateUser(ctx context.Context, in *user_managementpb.CreateUserRequest) (*user_managementpb.UserProfile, error) {
	glog.Info("[CreateUser]", "user_id ", in.User.GetPhone())

	// validate token first

	_, err := u.sqlrepo.CreateUser(ctx, (*entity.UserProfile)(in.User))
	if err != nil {
		glog.Info("[CreateUser]", "error ", err)
		return nil, err
	}

	return in.GetUser(), nil
}
