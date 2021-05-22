package model

import "github.com/dgrijalva/jwt-go"

type AccessTokenClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}
