package websocket

import (
	"github.com/MrKAKTyC/lets-go-chat/pkg/repository"
	"log"
	"time"

	"github.com/MrKAKTyC/lets-go-chat/pkg/dao"
	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/types"
	"github.com/MrKAKTyC/lets-go-chat/pkg/service"
	"github.com/labstack/echo/v4"
)

type Room interface {
	Run()
	GetActiveUsers() int
	Join(activeUser *ActiveUser)
	Leave(activeUser *ActiveUser)
	ServeWs(ctx echo.Context, params types.WsRTMStartParams) error
}

type ChatRoom struct {
	clients     map[*ActiveUser]bool
	messageChan chan dao.Message
	otpService  *service.OTPService
	messageRepo *repository.MessageRepository
	userRepo    *repository.UserRepository
}

func NewChatRoom(otpService *service.OTPService, messageRepo *repository.MessageRepository, userRepo *repository.UserRepository) *Room {
	var room Room
	room = &ChatRoom{
		clients:     make(map[*ActiveUser]bool),
		messageChan: make(chan dao.Message),
		otpService:  otpService,
		messageRepo: messageRepo,
		userRepo:    userRepo,
	}
	return &room
}

func (cr *ChatRoom) Run() {
	for {
		message := <-cr.messageChan
		(*cr.messageRepo).Create(message)
		for client := range cr.clients {
			client.inbox <- message
		}
	}
}

func (cr *ChatRoom) GetActiveUsers() int {
	return len(cr.clients)
}

func (cr *ChatRoom) Join(activeUser *ActiveUser) {
	log.Println("User ", activeUser.realUserID, " is joining")
	cr.clients[activeUser] = true
	userLastOnlineTime, err := (*cr.userRepo).GetLastOnline(activeUser.realUserID)
	if err != nil {
		log.Println("Cant get last online for user", activeUser.realUserID)
		log.Println(err)
	}
	log.Println(activeUser.realUserID, "last online: ", userLastOnlineTime)
	missedMessages, err := (*cr.messageRepo).GetAfter(*userLastOnlineTime)
	log.Println(activeUser.realUserID, "missed", len(missedMessages), "message(s)")
	if err != nil {
		log.Println(err)
	}
	for _, message := range missedMessages {
		activeUser.inbox <- *message
	}
}

func (cr *ChatRoom) Leave(activeUser *ActiveUser) {
	log.Println("User ", activeUser.realUserID, " is leaving")
	delete(cr.clients, activeUser)
	(*cr.userRepo).UpdateLastOnline(activeUser.realUserID, time.Now())
	activeUser.conn.Close()
}
func (cr *ChatRoom) ServeWs(ctx echo.Context, params types.WsRTMStartParams) error {
	log.Println("Obtaining OTP for: ", params.Token)
	userID, err := (*cr.otpService).UseOTP(params.Token)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Upgrading connection for: ", userID)
	conn, err := upgrader.Upgrade(ctx.Response().Writer, ctx.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	client := &ActiveUser{
		realUserID: userID,
		chatRoom:   cr,
		conn:       conn,
		inbox:      make(chan dao.Message, 256),
	}
	cr.Join(client)

	// new goroutines.
	go client.WriteMessage()
	go client.ReadMessage()
	return nil
}
