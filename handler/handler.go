package handler

import (
	otp "github.com/vegggg/user-management/otp"
	user "github.com/vegggg/user-management/user"
)

type UserManagement struct {
	user user.UserService
	otp  otp.OTPService
}

func NewUserManagement(user user.UserService, otp otp.OTPService) *UserManagement {
	u := new(UserManagement)
	u.user = user
	u.otp = otp
	return u
}
