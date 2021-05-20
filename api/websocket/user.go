package websocket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type user struct {
	username string
	roomID   string
	conn     *websocket.Conn
	chat     *chatRoom
}

func (u *user) read() {
	for {
		if _, payload, err := u.conn.ReadMessage(); err != nil {
			log.Printf("Error on read message: %v\n", err)
			break
		} else {
			u.chat.send <- newMessage(u.roomID, u.username, string(payload))
		}
	}

	u.chat.unregister <- u
}

func (u *user) write(message *message) {
	msg, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshalling message: %v\n", err)
		return
	}

	if err := u.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Printf("Error on write message: %v\n", err)
	}
}
