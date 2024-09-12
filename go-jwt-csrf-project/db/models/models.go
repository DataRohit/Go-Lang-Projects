package models

import (
	"time"

	"github.com/datarohit/go-jwt-csrf-project/randomStrings"
	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	Username     string
	PasswordHash string
	Role         string
}

type TokenClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
	Csrf string `json:"csrf"`
}

const (
	RoleAdmin             = "admin"
	RoleUser              = "user"
	RefreshTokenValidTime = time.Hour * 72
	AuthTokenValidTime    = time.Minute * 15
)

func GenerateCSRFSecret() (string, error) {
	return randomStrings.GenerateRandomString(32)
}
