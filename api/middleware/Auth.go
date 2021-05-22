package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/bagus2x/inamen-go-be/pkg/auth"
)

type authentication struct {
	auth auth.Service
}

type Authentication interface {
	Auth(next http.Handler) http.HandlerFunc
}

func NewAuth(auth auth.Service) Authentication {
	return &authentication{
		auth: auth,
	}
}

func (a *authentication) Auth(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		splittedHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(splittedHeader) < 2 {
			failure(w, "Authorization header must be filled", 401)
			return
		}

		claims, err := a.auth.ParseAccessToken(splittedHeader[1])
		if err != nil {
			failure(w, err.Error(), 400)
			return
		}

		userIDCtx := context.WithValue(r.Context(), "userID", claims.Subject)
		usernameCtx := context.WithValue(userIDCtx, "username", claims.Username)

		next.ServeHTTP(w, r.WithContext(usernameCtx))
	})
}
