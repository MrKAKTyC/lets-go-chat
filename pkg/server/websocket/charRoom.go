package websocket

import (
	"log"
	"time"

	"github.com/MrKAKTyC/lets-go-chat/pkg/dao"
	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/types"
	"github.com/MrKAKTyC/lets-go-chat/pkg/service"
	"github.com/labstack/echo/v4"
)

type messageRepository interface {
	GetAfter(after time.Time) ([]*dao.Message, error)
	Create(message dao.Message) error
}

type userRepository interface {
	Get(login, password string) (*dao.User, error)
	Create(login, password string) (*dao.User, error)
	GetLastOnline(userID string) (*time.Time, error)
	UpdateLastOnline(userID string, logoutDate time.Time) error
}

type ChatRoom struct {
	clients     map[*ActiveUser]bool
	messageChan chan dao.Message
	otpService  *service.OtpService
	messageRepo messageRepository
	userRepo    userRepository
}

func NewChatRoom(otpService *service.OtpService, messageRepo messageRepository, userRepo userRepository) *ChatRoom {
	return &ChatRoom{
		clients:     make(map[*ActiveUser]bool),
		messageChan: make(chan dao.Message),
		otpService:  otpService,
		messageRepo: messageRepo,
		userRepo:    userRepo,
	}
}

func (cr *ChatRoom) Run() {
	for {
		message := <-cr.messageChan
		cr.messageRepo.Create(message)
		for client := range cr.clients {
			client.inbox <- message
		}
	}
}

func (cr *ChatRoom) GetActiveUsers() int {
	return len(cr.clients)
}

func (cr *ChatRoom) Join(activeUser *ActiveUser) {
	log.Printf("User %s is joining", activeUser.realUserID)
	cr.clients[activeUser] = true
	userLastOnlineTime, err := cr.userRepo.GetLastOnline(activeUser.realUserID)
	if err != nil {
		log.Printf("Cant get last online for user %s\n%s", activeUser.realUserID, err)
	}
	log.Printf("%s last online: %s", activeUser.realUserID, userLastOnlineTime)
	missedMessages, err := cr.messageRepo.GetAfter(*userLastOnlineTime)
	log.Printf("%s missed %d message(s)", activeUser.realUserID, len(missedMessages))
	if err != nil {
		log.Println(err)
	}
	for _, message := range missedMessages {
		activeUser.inbox <- *message
	}
}

func (cr *ChatRoom) Leave(activeUser *ActiveUser) {
	log.Printf("User %s is leaving", activeUser.realUserID)
	delete(cr.clients, activeUser)
	cr.userRepo.UpdateLastOnline(activeUser.realUserID, time.Now())
	activeUser.conn.Close()
}
func (cr *ChatRoom) ServeWs(ctx echo.Context, params types.WsRTMStartParams) error {
	log.Printf("Obtaining OTP for: %s", params.Token)
	userID, err := cr.otpService.UseOTP(params.Token)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Upgrading connection for: %s", userID)
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
