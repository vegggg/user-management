package handler

import (
	"context"

	user_managementpb "github.com/vegggg/user-management/proto/user_mgnt/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u UserManagement) VerifyOTP(ctx context.Context, in *user_managementpb.VerifyOTPRequest) (*user_managementpb.VerifyOTPResponse, error) {
	if in.GetOtp() == "123456" {
		return &user_managementpb.VerifyOTPResponse{
			IsVerified: true,
		}, nil
	}
	return nil, status.Error(codes.Aborted, "wrong OTP")
}
