package service

import (
	"errors"
	"testing"

	"github.com/MrKAKTyC/lets-go-chat/pkg/dao"
	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/auth"
	"github.com/MrKAKTyC/lets-go-chat/pkg/service/mocks"
	"github.com/stretchr/testify/mock"
)

var testRegisterRepository = []struct {
	mockArgument      string
	createUserRequest *auth.CreateUserRequest
	err               error
}{
	{
		mockArgument: "UName",
		createUserRequest: &auth.CreateUserRequest{
			UserName: "UName",
			Password: "Password",
		},
		err: nil,
	},
	{
		mockArgument: "UsedName",
		createUserRequest: &auth.CreateUserRequest{
			UserName: "UsedName",
			Password: "Password",
		},
		err: errors.New("fail"),
	},
}

var testInvalidUser = []struct {
	invalidUserRequest auth.CreateUserRequest
}{
	{
		invalidUserRequest: auth.CreateUserRequest{
			UserName: "sh",
			Password: "VeryValidPassword",
		},
	},
	{
		invalidUserRequest: auth.CreateUserRequest{
			UserName: "VeryValidUserName",
			Password: "pwd",
		},
	},
}

func initUserService() (*User, *mocks.UserRepository, *mocks.OTPService) {
	userRepository := &mocks.UserRepository{}
	otpService := &mocks.OTPService{}
	userService := NewUserService(userRepository, otpService, "url/")
	return userService, userRepository, otpService
}

func TestRegister(t *testing.T) {
	userDAO := &dao.User{
		ID:       "id",
		Login:    "UName",
		Password: "Password",
	}
	for _, testCase := range testRegisterRepository {
		userService, userRepository, _ := initUserService()
		userRepository.On("Create", testCase.mockArgument, mock.Anything).Return(userDAO, testCase.err)
		_, err := userService.Register(*testCase.createUserRequest)

		if err != testCase.err {
			t.Errorf("No errors expected: %s", err.Error())
		}
	}

	for _, testCase := range testInvalidUser {
		userService, _, _ := initUserService()
		_, err := userService.Register(testCase.invalidUserRequest)
		if err == nil {
			t.Error("User shouldn't be created with invalid params")
		}
	}

}

var testAuthorize = []struct {
	dao              *dao.User
	loginUserRequest *auth.LoginUserRequest
}{
	{
		dao: &dao.User{},
		loginUserRequest: &auth.LoginUserRequest{
			UserName: "UNameValid",
			Password: "Password",
		},
	},
	{
		dao: nil,
		loginUserRequest: &auth.LoginUserRequest{
			UserName: "UNameInvalid",
			Password: "Password",
		},
	},
}

func TestAuthorize(t *testing.T) {

	for _, testCase := range testAuthorize {
		userService, userRepository, otpService := initUserService()
		otpService.On("GenerateOTP").Return("OTP")

		userRepository.On("Get", testCase.loginUserRequest.UserName, mock.Anything).Return(testCase.dao, nil)
		response, err := userService.Authorize(*testCase.loginUserRequest)
		if err != nil {
			t.Error("No errors expected")
		}
		if response.Url != "url/OTP" {
			t.Errorf("Wrong url %s", response.Url)
		}
	}

}
