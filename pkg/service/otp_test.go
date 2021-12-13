package service

import (
	"testing"
	"time"
)

func TestGenerateOTP(t *testing.T) {
	storage := make(map[string]time.Time)
	otp := newOtpService(storage)
	otp.GenerateOTP()
	if len(storage) != 1 {
		t.Error("Storage should contain generated OTP")
	}
}

func TestUseOTP(t *testing.T) {
	storage := make(map[string]time.Time)
	otpService := newOtpService(storage)

	otp := otpService.GenerateOTP()
	err := otpService.UseOTP(otp)
	if err != nil {
		t.Error("OTP should be extractable")
	}
}