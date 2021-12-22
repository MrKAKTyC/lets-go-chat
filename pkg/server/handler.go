package serv

import (
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
	defer dbConnection.Close()
	userRepository := repository.UserPGS(dbConnection)
	messageRepository := repository.MessagePGS(dbConnection)
	otpService := service.NewOtpService()

	chatRoom := websocket.NewChatRoom(otpService, messageRepository, userRepository)
	userService := service.NewUserService(userRepository, otpService, "/chat/ws.rtm.start?token=")

	go chatRoom.Run()
	server := &controller.User{UserService: userService, ChatRoom: chatRoom}

	serv.RegisterHandlers(router, server)
	handlerChain := middleware.PanicRecoverer(middleware.ErrorLogger(middleware.RequestLogger(router)))

	http.ListenAndServe(":"+config.Server.Port, handlerChain)
}
