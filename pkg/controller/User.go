package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/auth"
	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/types"
	"github.com/MrKAKTyC/lets-go-chat/pkg/server/websocket"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	Register(user auth.CreateUserRequest) (*auth.CreateUserResponse, error)
	Authorize(user auth.LoginUserRequest) (*auth.LoginUserResponse, error)
}

type ChatRoom interface {
	Run()
	GetActiveUsers() int
	Join(activeUser *websocket.ActiveUser) error
	ServeWs(ctx echo.Context, params types.WsRTMStartParams) error
}

type User struct {
	UserService UserService
	ChatRoom    ChatRoom
}

// Create converts echo context to params.
func (c *User) CreateUser(ctx echo.Context) error {
	req := ctx.Request()
	login, password := req.FormValue("userName"), req.FormValue("password")
	user, err := c.UserService.Register(auth.CreateUserRequest{Password: password, UserName: login})
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

// Login converts echo context to params.
func (c *User) LoginUser(ctx echo.Context) error {
	var err error
	req := ctx.Request()
	login, password := req.FormValue("userName"), req.FormValue("password")
	resp, err := c.UserService.Authorize(auth.LoginUserRequest{Password: password, UserName: login})
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

func (c *User) GetActiveUsers(ctx echo.Context) error {
	activeUsers, _ := json.Marshal(c.ChatRoom.GetActiveUsers())
	sendJSONResponse(ctx.Response().Writer, activeUsers)
	return nil
}

func (c *User) WsRTMStart(ctx echo.Context, params types.WsRTMStartParams) error {
	err := c.ChatRoom.ServeWs(ctx, params)
	if err != nil {
		sendError(ctx.Response().Writer, err)
	}
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
