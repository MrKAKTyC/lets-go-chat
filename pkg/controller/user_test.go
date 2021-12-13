package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MrKAKTyC/lets-go-chat/pkg/controller/mocks"
	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/auth"
	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/types"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

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
	userController, us, _ := initController()
	c := initContext()

	id, name := "uid", "name"
	LURes := &auth.CreateUserResponse{Id: &id, UserName: &name}
	us.On("Register", mock.Anything).Return(LURes, nil)
	err := userController.CreateUser(c)
	if err != nil {
		t.Error("No error is expected", err)
	}

	userController, us, _ = initController()
	us.On("Register", mock.Anything).Return(nil, errors.New("Service fail"))
	err = userController.CreateUser(c)
	if err == nil {
		t.Error("Error is expected")
	}

}

func TestLoginUser(t *testing.T) {
	userController, us, _ := initController()
	c := initContext()

	LURes := &auth.LoginUserResponse{Url: "url"}
	us.On("Authorize", mock.Anything).Return(LURes, nil)
	err := userController.LoginUser(c)
	if err != nil {
		t.Error("No error is expected", err)
	}

	userController, us, _ = initController()
	us.On("Authorize", mock.Anything).Return(nil, errors.New("Service fail"))
	err = userController.LoginUser(c)
	if err == nil {
		t.Error("Error is expected")
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
	userController, _, cr := initController()
	c := initContext()

	params := types.WsRTMStartParams{}
	cr.On("ServeWs", mock.Anything, mock.Anything).Return(nil)
	err := userController.WsRTMStart(c, params)
	if err != nil {
		t.Error("No error expected", err)
	}

	userController, _, cr = initController()
	cr.On("ServeWs", mock.Anything, mock.Anything).Return(errors.New("Service fail"))
	err = userController.WsRTMStart(c, params)
	if err == nil {
		t.Error("Error is expected")
	}

}
