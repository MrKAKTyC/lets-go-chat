package service

import (
	"errors"
	"testing"

	"github.com/MrKAKTyC/lets-go-chat/pkg/dao"
	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/auth"
	"github.com/MrKAKTyC/lets-go-chat/pkg/service/mocks"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	userRepository := mocks.UserRepository{}
	OtpService := mocks.OTPService{}
	createUserRequest := auth.CreateUserRequest{
		UserName: "UName",
		Password: "Password",
	}
	userService := NewUserService(&userRepository, &OtpService, "abc")
	uDao := &dao.User{
		ID:       "id",
		Login:    "UName",
		Password: "Password",
	}
	userRepository.On("Create", "UName", mock.Anything).Return(uDao, nil)

	createUserResponse, err := userService.Register(createUserRequest)

	if *createUserResponse.UserName != createUserRequest.UserName {
		t.Error("Created user should have same name")
	}
	if err != nil {
		t.Errorf("No errors expected: %s", err.Error())
	}

	createUserRequest = auth.CreateUserRequest{
		UserName: "sh",
		Password: "Pass",
	}
	_, err = userService.Register(createUserRequest)
	if err == nil {
		t.Errorf("Expected error for short fields")
	}

	createUserRequestUsed := auth.CreateUserRequest{
		UserName: "UsedName",
		Password: "Password",
	}
	userRepository.On("Create", "UsedName", mock.Anything).Return(nil, errors.New("Fail"))
	_, err = userService.Register(createUserRequestUsed)
	if err == nil {
		t.Error("Error is expected")
	}
}

func TestAuthorize(t *testing.T) {
	userRepository := mocks.UserRepository{}
	OtpService := mocks.OTPService{}
	userService := NewUserService(&userRepository, &OtpService, "url/")
	loginRequestValid := auth.LoginUserRequest{
		UserName: "UNameValid",
		Password: "Password",
	}
	loginRequestInvalid := auth.LoginUserRequest{
		UserName: "UNameInvalid",
		Password: "Password",
	}

	userDao := &dao.User{}
	userRepository.On("Get", "UNameValid", mock.Anything).Return(userDao, nil)
	OtpService.On("GenerateOTP").Return("OTP")
	response, err := userService.Authorize(loginRequestValid)
	if err != nil {
		t.Error("No errors expected")
	}
	if response.Url != "url/OTP" {
		t.Errorf("Wrong url %s", response.Url)
	}

	userRepository.On("Get", "UNameInvalid", mock.Anything).Return(nil, errors.New(""))
	_, err = userService.Authorize(loginRequestInvalid)
	if err == nil {
		t.Error("Error is expected")
	}

}
