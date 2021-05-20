package main

import (
	"net/http"

	"github.com/bagus2x/inamen-go-be/api/websocket"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {

	r := chi.NewMux()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*", "ws://*", "wss://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello"))
	})

	websocket.Start(r)

	http.ListenAndServe(":8080", r)
}
