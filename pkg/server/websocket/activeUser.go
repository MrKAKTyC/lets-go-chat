package websocket

import (
	"bytes"
	"log"
	"time"

	"github.com/MrKAKTyC/lets-go-chat/pkg/dao"
	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ActiveUser struct {
	realUserID string
	chatRoom   *ChatRoom
	conn       *websocket.Conn
	inbox      chan dao.Message
}

func (au *ActiveUser) ReadMessage() {
	defer func() {
		au.chatRoom.Leave(au)
	}()
	au.conn.SetReadLimit(maxMessageSize)
	au.conn.SetReadDeadline(time.Now().Add(pongWait))
	au.conn.SetPongHandler(func(string) error { au.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, message, err := au.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Print("error: ", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		au.chatRoom.messageChan <- dao.Message{
			Sender:  au.realUserID,
			Content: string(message),
			Date:    time.Now(),
		}
	}
}

func (au *ActiveUser) WriteMessage() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		au.conn.Close()
	}()
	for {
		select {
		case message, ok := <-au.inbox:
			au.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				au.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			writer, err := au.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println(err)
				return
			}
			writer.Write([]byte(message.Content))

			if err := writer.Close(); err != nil {
				log.Println(err)
				return
			}
		case <-ticker.C:
			au.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := au.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
