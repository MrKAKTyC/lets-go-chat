package controller

import (
	"encoding/json"
	"github.com/MrKAKTyC/lets-go-chat/pkg/service"
	"log"
	"net/http"
	"time"

	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/auth"
	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/types"
	"github.com/MrKAKTyC/lets-go-chat/pkg/server/websocket"
	"github.com/labstack/echo/v4"
)

type User struct {
	UserService *service.UserService
	ChatRoom    *websocket.Room
}

func NewUser(userService *service.UserService, chatRoom *websocket.Room) *User {
	return &User{
		UserService: userService,
		ChatRoom:    chatRoom,
	}
}

// Create converts echo context to params.
func (c *User) CreateUser(ctx echo.Context) error {
	req := ctx.Request()
	login, password := req.FormValue("userName"), req.FormValue("password")
	user, err := (*c.UserService).Register(auth.CreateUserRequest{Password: password, UserName: login})
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
	req := ctx.Request()
	login, password := req.FormValue("userName"), req.FormValue("password")
	resp, err := (*c.UserService).Authorize(auth.LoginUserRequest{Password: password, UserName: login})
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
	activeUsers, err := json.Marshal((*c.ChatRoom).GetActiveUsers())
	sendJSONResponse(ctx.Response().Writer, activeUsers)
	return err
}

func (c *User) WsRTMStart(ctx echo.Context, params types.WsRTMStartParams) error {
	err := (*c.ChatRoom).ServeWs(ctx, params)
	if err != nil {
		sendError(ctx.Response().Writer, err)
	}
	return err
}

func sendError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	_, writeErr := w.Write([]byte(err.Error()))
	if err != nil {
		log.Println(writeErr)
	}
}

func sendJSONResponse(w http.ResponseWriter, jsonResponse []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
	}
}
