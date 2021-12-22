package service

import (
	"testing"
)

var DBuserID = "UUID-ID"

func TestGenerateOTP(t *testing.T) {
	storage := make(map[string]string)
	otp := newOtpService(storage)
	otp.GenerateOTP(DBuserID)
	if len(storage) != 1 {
		t.Error("Storage should contain generated OTP")
	}
}

func TestUseOTP(t *testing.T) {
	storage := make(map[string]string)
	otpService := newOtpService(storage)

	otp := otpService.GenerateOTP(DBuserID)
	userID, err := otpService.UseOTP(otp)
	if err != nil {
		t.Error("OTP should be extractable")
	}
	if DBuserID != userID {
		t.Error("ID after extraction should be same as placed", DBuserID, userID)
	}
}
