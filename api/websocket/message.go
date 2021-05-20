package websocket

import (
	"time"

	"github.com/bagus2x/inamen-go-be/utils"
)

type message struct {
	ID        string `json:"_id"`
	RoomID    string `json:"roomID"`
	Sender    string `json:"sender"`
	Body      string `json:"body"`
	CreatedAt int64  `json:"createdAt"`
}

func newMessage(roomID, sender, body string) *message {
	return &message{
		ID:        utils.GenerateID().Hex(),
		RoomID:    roomID,
		Sender:    sender,
		Body:      body,
		CreatedAt: time.Now().Unix(),
	}
}
