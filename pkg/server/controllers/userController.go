package userController

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MrKAKTyC/lets-go-chat/client/auth"
	"github.com/MrKAKTyC/lets-go-chat/pkg/services"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService services.UserService
}

// CreateUser converts echo context to params.
func (controller *UserController) CreateUser(ctx echo.Context) error {
	fmt.Println("Creating user")
	var err error
	req := ctx.Request()
	login, password := req.FormValue("userName"), req.FormValue("password")
	user, err := controller.UserService.RegisterUser(auth.CreateUserRequest{Password: password, UserName: login})
	if err != nil {
		sendError(ctx.Response().Writer, err)
		return err
	}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		sendError(ctx.Response().Writer, err)
		return err
	}
	sendJSONResponse(ctx.Response().Writer, jsonUser)
	return err
}

// LoginUser converts echo context to params.
func (controller *UserController) LoginUser(ctx echo.Context) error {
	fmt.Println("Logingin user")
	var err error
	req := ctx.Request()
	login, password := req.FormValue("userName"), req.FormValue("password")
	resp, err := controller.UserService.AuthorizeUser(auth.LoginUserRequest{Password: password, UserName: login})
	if err != nil {
		sendError(ctx.Response().Writer, err)
		return err
	}
	jsonUser, _ := json.Marshal(resp)
	ctx.Response().Header().Set("X-Rate-Limit", "120")
	ctx.Response().Header().Set("X-Expires-After", time.Now().AddDate(0, 0, 1).UTC().String())
	sendJSONResponse(ctx.Response().Writer, jsonUser)
	return err
}

func sendError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}

func sendJSONResponse(w http.ResponseWriter, jsonResponse []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
