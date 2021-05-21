package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bagus2x/inamen-go-be/api/websocket"
	"github.com/bagus2x/inamen-go-be/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	// Environment Variables
	mongoUri := os.Getenv("MONGO_URI")
	mongoName := os.Getenv("MONGO_DATABASE")
	mongoMinPool := os.Getenv("MONGO_MIN_POOL")
	mongoMaxPool := os.Getenv("MONGO_MAX_POOL")
	mongoMaxIdle := os.Getenv("MONGO_MAX_IDLE_SECOND")

	minPool, err := strconv.Atoi(mongoMinPool)
	if err != nil {
		log.Fatal(err)
	}

	maxPool, err := strconv.Atoi(mongoMaxPool)
	if err != nil {
		log.Fatal(err)
	}

	maxIdle, err := strconv.Atoi(mongoMaxIdle)
	if err != nil {
		log.Fatal(err)
	}

	if minPool < 0 || maxPool < 0 || maxIdle < 0 {
		log.Fatal("Maxpool or minpool or maxIdle must be positive")
	}

	// Database Connection
	db, err := config.MongoConnection(mongoName, mongoUri, uint64(minPool), uint64(maxPool), uint64(maxIdle))
	if err != nil {
		log.Fatal(err)
	}

	_ = db

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
