package http

import (
	"encoding/json"
	"net/http"

	"github.com/bagus2x/inamen-go-be/api/middleware"
	"github.com/bagus2x/inamen-go-be/pkg/model"
	"github.com/bagus2x/inamen-go-be/pkg/tournament"
	"github.com/go-chi/chi/v5"
)

func Tournament(r chi.Router, service tournament.Service, auth middleware.Authentication) {
	r.Route("/api/v1/tournaments", func(r chi.Router) {
		r.Get("/", auth.Auth(getTournamentsByHostID(service)))
		r.Post("/", auth.Auth(createTournament(service)))
	})
}

func createTournament(service tournament.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		hostID := ctx.Value("userID").(string)
		username := ctx.Value("username").(string)

		var req model.CreateTourRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			failure(w, err.Error(), status(err))
			return
		}

		defer r.Body.Close()

		res, err := service.Create(hostID, username, &req)
		if err != nil {
			failure(w, err.Error(), status(err))
			return
		}

		success(w, res, "", 201)
	}
}

func getTournamentsByHostID(service tournament.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		hostID := ctx.Value("userID").(string)

		res, err := service.FetchToursByHostID(hostID)
		if err != nil {
			failure(w, err.Error(), status(err))
			return
		}

		success(w, res, "", 200)
	}
}
