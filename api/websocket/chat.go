package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

type chatRoom struct {
	rooms      map[string]map[string]*user
	send       chan *message
	register   chan *user
	unregister chan *user
}

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1 << 10,
	ReadBufferSize:  1 << 10,
	CheckOrigin:     func(r *http.Request) bool { return r.Method == http.MethodGet },
}

func (cr *chatRoom) listen() {
	for {
		select {
		case user := <-cr.register:
			cr.add(user)
		case message := <-cr.send:
			cr.broadcast(message)
		case user := <-cr.unregister:
			cr.disconnect(user)
		}
	}
}

func (cr *chatRoom) add(u *user) {
	// create a room if not exist
	if _, ok := cr.rooms[u.roomID]; !ok {
		cr.rooms[u.roomID] = make(map[string]*user)
	}

	cr.rooms[u.roomID][u.username] = u

	body := fmt.Sprintf("%s register the chat\n", u.username)
	cr.broadcast(newMessage(u.roomID, "SERVER", body))
}

func (cr *chatRoom) broadcast(message *message) {
	for _, u := range cr.rooms[message.RoomID] {
		u.write(message)
	}
}

func (cr *chatRoom) disconnect(u *user) {
	if _, ok := cr.rooms[u.roomID]; ok {
		defer cr.rooms[u.roomID][u.username].conn.Close()
		delete(cr.rooms[u.roomID], u.username)

		body := fmt.Sprintf("%s has left the chat\n", u.username)
		cr.broadcast(newMessage(u.roomID, "SERVER", body))
	}
}

func (cr *chatRoom) handleChat(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error on upgrade: %v\n", err)
		return
	}

	roomID := chi.URLParam(r, "roomID")
	username := chi.URLParam(r, "username")

	u := &user{
		username: username,
		roomID:   roomID,
		conn:     conn,
		chat:     cr,
	}

	cr.register <- u

	u.read()
}

func Start(r chi.Router) {
	c := &chatRoom{
		rooms:      make(map[string]map[string]*user),
		send:       make(chan *message),
		register:   make(chan *user),
		unregister: make(chan *user),
	}

	r.Get("/ws/{roomID}/{username}", c.handleChat)

	go c.listen()
}
