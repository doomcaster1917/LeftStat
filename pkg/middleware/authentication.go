package middleware

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthorizationHeader struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    uint32 `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

var secretKey = []byte("sDf3223_2*''sdsf3")

func getToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 720).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func check_JWT(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return errors.New(fmt.Sprintf("%v", err))
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok {
		return nil
	} else {
		return errors.New(fmt.Sprintf("%v", err))
	}
}

func Auth(authHeader string) error {

	err := check_JWT(authHeader)
	if err != nil {
		return errors.New(fmt.Sprintf("%v", err))
	} else {
		return nil
	}
}
