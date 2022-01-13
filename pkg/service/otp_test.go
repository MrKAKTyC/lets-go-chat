package service

import (
	"testing"
)

var dBuserID = "UUID-ID"

func TestGenerateOTP(t *testing.T) {
	storage := make(map[string]string)
	otp := NewOtpService(storage)
	otp.GenerateOTP(dBuserID)
	if len(storage) != 1 {
		t.Error("Storage should contain generated OTP")
	}
}

func TestUseOTP(t *testing.T) {
	storage := make(map[string]string)
	otpService := NewOtpService(storage)

	otp := otpService.GenerateOTP(dBuserID)
	userID, err := otpService.UseOTP(otp)
	if err != nil {
		t.Error("OTP should be extractable")
	}
	if dBuserID != userID {
		t.Error("ID after extraction should be same as placed", DBuserID, userID)
	}
}
