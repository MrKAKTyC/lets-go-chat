package websocket

import (
	"bytes"
	"log"
	"time"

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
	chatRoom *ChatRoom
	conn     *websocket.Conn
	send     chan []byte
}

func (au *ActiveUser) ReadMessage() {
	defer func() {
		delete(au.chatRoom.clients, au)
		au.conn.Close()
	}()
	au.conn.SetReadLimit(maxMessageSize)
	au.conn.SetReadDeadline(time.Now().Add(pongWait))
	au.conn.SetPongHandler(func(string) error { au.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := au.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		au.chatRoom.broadcast <- message
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
		case message, ok := <-au.send:
			au.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				au.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := au.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			n := len(au.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-au.send)
			}

			if err := w.Close(); err != nil {
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
