package serv

import (
	"net/http"

	"github.com/MrKAKTyC/lets-go-chat/pkg/config"
	"github.com/MrKAKTyC/lets-go-chat/pkg/controller"
	serv "github.com/MrKAKTyC/lets-go-chat/pkg/generated"
	"github.com/MrKAKTyC/lets-go-chat/pkg/repository"
	"github.com/MrKAKTyC/lets-go-chat/pkg/service"

	"github.com/labstack/echo/v4"
)

func Serve(config config.Config) {
	router := echo.New()
	server := &controller.User{Service: service.New(repository.UserPGS(config.DB.URL), "url")}

	serv.RegisterHandlers(router, server)

	http.ListenAndServe(":"+config.Server.Port, router)
}
