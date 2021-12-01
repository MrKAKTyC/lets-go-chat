// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package server

import (
	"fmt"
	"net/http"

	. "github.com/MrKAKTyC/lets-go-chat/pkg/generated/types"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Endpoint to start real time chat
	// (GET /chat/ws.rtm.start)
	WsRTMStart(ctx echo.Context, params WsRTMStartParams) error
	// Register (create) user
	// (POST /user)
	CreateUser(ctx echo.Context) error
	// Number of active users in a chat
	// (GET /user/active)
	GetActiveUsers(ctx echo.Context) error
	// Logs user into the system
	// (POST /user/login)
	LoginUser(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// WsRTMStart converts echo context to params.
func (w *ServerInterfaceWrapper) WsRTMStart(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params WsRTMStartParams
	// ------------- Required query parameter "token" -------------

	err = runtime.BindQueryParameter("form", true, true, "token", ctx.QueryParams(), &params.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter token: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.WsRTMStart(ctx, params)
	return err
}

// CreateUser converts echo context to params.
func (w *ServerInterfaceWrapper) CreateUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateUser(ctx)
	return err
}

// GetActiveUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetActiveUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetActiveUsers(ctx)
	return err
}

// LoginUser converts echo context to params.
func (w *ServerInterfaceWrapper) LoginUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.LoginUser(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/chat/ws.rtm.start", wrapper.WsRTMStart)
	router.POST(baseURL+"/user", wrapper.CreateUser)
	router.GET(baseURL+"/user/active", wrapper.GetActiveUsers)
	router.POST(baseURL+"/user/login", wrapper.LoginUser)

}
