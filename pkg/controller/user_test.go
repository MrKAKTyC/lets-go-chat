package controller

import (
	"errors"
	"github.com/MrKAKTyC/lets-go-chat/pkg/controller/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/auth"
	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/types"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

var id, name, serviceError = "uid", "name", errors.New("service fail")

var testRegister = []struct {
	userResponse *auth.CreateUserResponse
	err          error
	expected     error
}{
	{&auth.CreateUserResponse{Id: &id, UserName: &name}, nil, nil}, //Happy pass
	{nil, serviceError, serviceError},                              //Service fail
}

var testLogin = []struct {
	userResponse *auth.LoginUserResponse
	err          error
	expected     error
}{
	{&auth.LoginUserResponse{Url: "url"}, nil, nil}, //Happy pass
	{nil, serviceError, serviceError},               //Service fail
}

var testWebSocket = []struct {
	err      error
	expected error
}{
	{nil, nil},                   //Happy pass
	{serviceError, serviceError}, //Service fail
}

func initController() (*User, *mocks.UserService, *mocks.ChatRoom) {
	us := &mocks.UserService{}
	cr := &mocks.ChatRoom{}
	userController := &User{
		UserService: us,
		ChatRoom:    cr,
	}
	return userController, us, cr
}

func initContext() echo.Context {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec)
}

func TestCreateUser(t *testing.T) {
	c := initContext()

	for _, testCase := range testRegister {
		userController, us, _ := initController()
		us.On("Register", mock.Anything).Return(testCase.userResponse, testCase.err)
		err := userController.CreateUser(c)
		if err != testCase.expected {
			t.Errorf("Unexpected result. For %s was expected %s, but get %s", testCase.err, testCase.expected, err)
		}
	}

}

func TestLoginUser(t *testing.T) {
	c := initContext()

	for _, testCase := range testLogin {
		userController, us, _ := initController()
		us.On("Authorize", mock.Anything).Return(testCase.userResponse, testCase.err)
		err := userController.LoginUser(c)
		if err != testCase.expected {
			t.Errorf("Unexpected result. For %s was expected %s, but get %s", testCase.err, testCase.expected, err)
		}
	}
}

func TestGetActiveUsers(t *testing.T) {
	userController, _, cr := initController()
	c := initContext()

	cr.On("GetActiveUsers").Return(10)
	err := userController.GetActiveUsers(c)
	if err != nil {
		t.Error("No error is expected", err)
	}
}

func TestWsRTMStart(t *testing.T) {
	c := initContext()

	params := types.WsRTMStartParams{}
	for _, testCase := range testWebSocket {
		userController, _, cr := initController()
		cr.On("ServeWs", mock.Anything, mock.Anything).Return(testCase.err)
		err := userController.WsRTMStart(c, params)
		if err != testCase.expected {
			t.Errorf("Unexpected result. For %s was expected %s, but get %s", testCase.err, testCase.expected, err)
		}
	}

}
