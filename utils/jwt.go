package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AshiClaims struct {
	Auth
	jwt.RegisteredClaims
}

var AshiSecret = []byte("nakashima_ashi_suki")

func GenToken(auth Auth) (string, error) {
	claim := AshiClaims{
		auth,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Nakashima",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString(AshiSecret)
	return tokenStr, err
}

func ParseToken(tokenStr string) (*AshiClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AshiClaims{}, func(token *jwt.Token) (interface{}, error) {
		return AshiSecret, nil
	})
	if err != nil {
		fmt.Println("token parse err :", err.Error())
	}

	if claim, ok := token.Claims.(*AshiClaims); ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("invalid token")
}

func RefreshToken(tokenStr string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenStr, &AshiClaims{}, func(token *jwt.Token) (interface{}, error) {
		return AshiSecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*AshiClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour))
		return GenToken(claims.Auth)
	}
	return "", errors.New("Cloudn't handle this token")
}
