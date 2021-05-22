package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	httpRoutes "github.com/bagus2x/inamen-go-be/api/http"
	"github.com/bagus2x/inamen-go-be/api/middleware"
	wsRoutes "github.com/bagus2x/inamen-go-be/api/websocket"
	"github.com/bagus2x/inamen-go-be/config"
	"github.com/bagus2x/inamen-go-be/pkg/auth"
	"github.com/bagus2x/inamen-go-be/pkg/tournament"
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
	accessToken := os.Getenv("ACCESS_TOKEN_KEY")

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

	// Router
	r := chi.NewMux()

	// Global Middlewares
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*", "ws://*", "wss://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Repositories
	tourRepo := tournament.NewRepo(db)

	// Services
	authService := auth.NewService(accessToken)
	tourService := tournament.NewService(tourRepo)

	// Middlewares
	authMiddleware := middleware.NewAuth(authService)

	// HTTP End Points
	httpRoutes.Tournament(r, tourService, authMiddleware)

	// Websocker End Points
	wsRoutes.Start(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
