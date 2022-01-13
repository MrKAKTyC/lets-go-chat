package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type OtpService struct {
	otpStorage map[string]string
}

func NewOtp(storage map[string]string) *OtpService {
	return &OtpService{otpStorage: storage}
}

func (s *OtpService) GenerateOTP(userId string) string {
	rand.Seed(time.Now().Unix())
	otpInt := rand.Int31()
	otp := fmt.Sprintf("%010d", otpInt)
	s.otpStorage[otp] = userId
	return otp
}

func (s *OtpService) UseOTP(otpToUse string) (string, error) {
	userId, ok := s.otpStorage[otpToUse]
	if ok {
		delete(s.otpStorage, otpToUse)
		return userId, nil
	}
	return "", errors.New("no such otp")

}
