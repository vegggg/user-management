package handler

import (
	"context"

	"github.com/golang/glog"
	"github.com/vegggg/user-management/entity"
	user_managementpb "github.com/vegggg/user-management/proto/user_mgnt/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u UserManagement) CreateUser(ctx context.Context, in *user_managementpb.CreateUserRequest) (*user_managementpb.UserProfile, error) {
	glog.Info("[CreateUser]", "user_id ", in.User.GetPhone())

	// validate token first
	if err := u.otp.IsPhoneValidatedByOTP(ctx, in.GetUser().GetPhone()); err != nil {
		glog.Info("[CreateUser] phone is not validated", "error ", err)
		return nil, status.Error(codes.FailedPrecondition, "otp is not validated")
	}

	_, err := u.user.CreateUser(ctx, (*entity.UserProfile)(in.User))
	if err != nil {
		glog.Info("[CreateUser]", "error ", err)
		return nil, err
	}

	return in.GetUser(), nil
}
