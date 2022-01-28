package serv

import (
	"log"
	"net/http"
	"time"

	"github.com/MrKAKTyC/lets-go-chat/pkg/config"
	"github.com/MrKAKTyC/lets-go-chat/pkg/controller"
	serv "github.com/MrKAKTyC/lets-go-chat/pkg/generated"
	"github.com/MrKAKTyC/lets-go-chat/pkg/repository"
	"github.com/MrKAKTyC/lets-go-chat/pkg/server/middleware"
	"github.com/MrKAKTyC/lets-go-chat/pkg/server/websocket"
	"github.com/MrKAKTyC/lets-go-chat/pkg/service"

	"github.com/labstack/echo/v4"
)

func Serve(config config.Config) {
	router := echo.New()
	userRepository := repository.UserPGS(config.DB.URL)
	otpService := service.NewOtpService(make(map[string]time.Time))
	chatRoom := websocket.NewChatRoom(otpService)
	userService := service.NewUserService(userRepository, otpService, "/chat/ws.rtm.start?token=")

	go chatRoom.Run()
	userController := controller.NewUser(*userService, *chatRoom)

	router.Use(middleware.PanicRecoverer)
	router.Use(middleware.ErrorLogger)
	router.Use(middleware.RequestLogger)
	serv.RegisterHandlers(router, userController)

	if err := http.ListenAndServe(":"+config.Server.Port, router); err != nil {
		log.Println(err)
	}
}
