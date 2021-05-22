package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/bagus2x/inamen-go-be/pkg/model"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

type service struct {
	accessTokenKey string
}

type Service interface {
	ParseAccessToken(accessToken string) (*model.AccessTokenClaims, error)
	CreateAccessToken(userID, username string) (string, error)
}

func NewService(accessTokenKey string) Service {
	return &service{
		accessTokenKey: accessTokenKey,
	}
}

func (s service) ParseAccessToken(accessToken string) (*model.AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &model.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.accessTokenKey), nil
	})

	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return nil, model.ErrTokenExpired
		}
		log.Error(err)
		return nil, model.ErrInvalidAccessToken
	}

	if claims, ok := token.Claims.(*model.AccessTokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, model.ErrInvalidAccessToken
}

func (s service) CreateAccessToken(userID, username string) (string, error) {
	claims := model.AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  "inamen.vercel.app",
			Subject:   userID,
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Username: username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.accessTokenKey))
}
