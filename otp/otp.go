package otp

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
)

var (
	otpKey         string = "otp-"
	isValidatedKey string = "val-"
)

type OTPService interface {
	GenOTP(ctx context.Context, userID string) (string, error)
	ValidateOTP(ctx context.Context, otp string, userID string) error
	IsPhoneValidatedByOTP(ctx context.Context, phone string) error
}

type RedisOTP struct {
	exp    time.Duration
	otpLen int
	cli    *redis.Client
}

func (r RedisOTP) GenOTP(ctx context.Context, userID string) (string, error) {
	otp := cast.ToString(genOtp(r.otpLen))
	_, err := r.cli.Set(context.Background(), otpKey+userID, otp, r.exp).Result()
	if err != nil {
		return "", err
	}
	return otp, err
}

func (r RedisOTP) ValidateOTP(ctx context.Context, otp string, userID string) error {
	rOtp, err := r.cli.Get(ctx, otpKey+userID).Result()
	if err != nil {
		return err
	}

	if rOtp != otp {
		return errors.New("wrong OTP")
	}

	// delete otp when verified
	r.cli.Del(ctx, otpKey+userID)

	_, err = r.cli.Set(ctx, isValidatedKey+userID, userID, r.exp).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r RedisOTP) IsPhoneValidatedByOTP(ctx context.Context, userID string) error {
	v, err := r.cli.Get(ctx, isValidatedKey+userID).Result()
	if err != nil {
		return err
	}

	if v != userID {
		return errors.New("phone is not validated by OTP")
	}

	return nil
}

func NewRedisOTP(r *redis.Client, otpLen int, exp time.Duration) OTPService {
	s := &RedisOTP{
		exp:    exp,
		otpLen: int(math.Pow10(otpLen)),
		cli:    r,
	}
	return s
}
