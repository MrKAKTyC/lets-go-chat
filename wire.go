package lets_go_chat

import (
	"github.com/MrKAKTyC/lets-go-chat/pkg/config"
	"github.com/MrKAKTyC/lets-go-chat/pkg/controller"
	"github.com/MrKAKTyC/lets-go-chat/pkg/repository"
	"github.com/MrKAKTyC/lets-go-chat/pkg/server/websocket"
	"github.com/MrKAKTyC/lets-go-chat/pkg/service"

	"github.com/google/wire"
)

func InitializeController(co config.Config) *controller.User {
	wire.Build(
		wire.FieldsOf(new(config.Config), "DB"),
		repository.GetDBConnection,
		repository.NewUserPGS,
		repository.NewMessagePGS,
		repository.NewInMemoryOTP,
		websocket.NewChatRoom,
		service.NewOtp,
		service.NewUser,
		controller.NewUser)

	return &controller.User{}
}
