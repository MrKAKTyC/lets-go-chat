package serv

import (
	"database/sql"
	log2 "github.com/labstack/gommon/log"
	"log"
	"net/http"

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
	dbConnection, err := repository.GetDBConnection(config.DB.URL)
	if err != nil {
		return
	}
	defer func(dbConnection *sql.DB) {
		err := dbConnection.Close()
		if err != nil {
			log2.Print(err)
		}
	}(dbConnection)
	userRepository := repository.UserPGS(dbConnection)
	messageRepository := repository.MessagePGS(dbConnection)
	otpService := service.NewOtp(make(map[string]string))

	chatRoom := websocket.NewChatRoom(otpService, messageRepository, userRepository)
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
