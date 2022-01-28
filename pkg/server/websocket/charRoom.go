package websocket

import (
	"log"

	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/types"
	"github.com/MrKAKTyC/lets-go-chat/pkg/service"
	"github.com/labstack/echo/v4"
)

type ChatRoom struct {
	clients    map[*ActiveUser]bool
	broadcast  chan []byte
	otpService service.OtpService
}

func NewChatRoom(otpService service.OtpService) *ChatRoom {
	return &ChatRoom{
		broadcast:  make(chan []byte),
		clients:    make(map[*ActiveUser]bool),
		otpService: otpService,
	}
}

func (cr *ChatRoom) Run() {
	for {
		message := <-cr.broadcast
		for client := range cr.clients {
			client.send <- message
		}
	}
}

func (cr *ChatRoom) GetActiveUsers() int {
	return len(cr.clients)
}

func (cr *ChatRoom) Join(activeUser *ActiveUser) error {
	cr.clients[activeUser] = true
	return nil
}
func ServeWs(cr *ChatRoom, ctx echo.Context, params types.WsRTMStartParams) error {
	err := cr.otpService.UseOTP(params.Token)
	if err != nil {
		return err
	}
	conn, err := upgrader.Upgrade(ctx.Response().Writer, ctx.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	client := &ActiveUser{chatRoom: cr, conn: conn, send: make(chan []byte, 256)}
	err = cr.Join(client)
	if err != nil {
		log.Println(err)
		return err
	}
	// new goroutines.
	go client.WriteMessage()
	go client.ReadMessage()
	return nil
}
