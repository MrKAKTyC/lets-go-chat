package service

import (
	"fmt"
	"github.com/MrKAKTyC/lets-go-chat/pkg/repository"
	"math/rand"
	"time"
)

type OTPService interface {
	GenerateOTP(userId string) string
	UseOTP(otpToUse string) (string, error)
}

type OtpService struct {
	otpStorage *repository.OtpStorage
}

func NewOtp(storage *repository.OtpStorage) *OTPService {
	var otp OTPService
	otp = &OtpService{otpStorage: storage}
	return &otp
}

func (s *OtpService) GenerateOTP(userId string) string {
	rand.Seed(time.Now().Unix())
	otpInt := rand.Int31()
	otp := fmt.Sprintf("%010d", otpInt)
	(*s.otpStorage).Put(otp, userId)
	return otp
}

func (s *OtpService) UseOTP(otpToUse string) (string, error) {
	return (*s.otpStorage).Get(otpToUse)
}
