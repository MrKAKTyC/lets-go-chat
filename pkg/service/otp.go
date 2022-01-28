package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type OtpService struct {
	otpStorage map[string]time.Time
}

func NewOtpService(storage map[string]time.Time) *OtpService {
	return &OtpService{otpStorage: storage}
}

func (s *OtpService) GenerateOTP() string {
	rand.Seed(time.Now().Unix())
	otpInt := rand.Int31()
	otp := fmt.Sprintf("%010d", otpInt)
	s.otpStorage[otp] = time.Now()
	return otp
}

func (s *OtpService) UseOTP(otpToUse string) error {
	_, ok := s.otpStorage[otpToUse]
	var err error
	if ok {
		delete(s.otpStorage, otpToUse)
	} else {
		err = errors.New("no such otp")
	}
	return err
}
