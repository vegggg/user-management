package handler

import (
	"context"

	user_managementpb "github.com/vegggg/user-management/proto/user_mgnt/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Send OTP on new user
func (u UserManagement) SendOTP(ctx context.Context, in *user_managementpb.SendOTPRequest) (*user_managementpb.SendOTPResponse, error) {
	oldUser, _ := u.user.GetUser(ctx, in.GetPhone())
	if oldUser != nil && oldUser.Phone != "" {
		return nil, status.Error(codes.AlreadyExists, "phone number is taken")
	}

	otp, err := u.otp.GenOTP(ctx, in.GetPhone())
	if err != nil {
		return nil, err
	}
	return &user_managementpb.SendOTPResponse{
		Otp: otp,
	}, nil
}
