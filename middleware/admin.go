package middleware

import (
	"errors"
)

type Credentials struct {
	Name     string
	Password string
}

var logAndPass = Credentials{Name: "doomcaster", Password: "89908990aSa"}

func (cr Credentials) Login() (string, error) {

	if cr == logAndPass {
		acessToken := getToken()
		return acessToken, nil
	} else {
		return "", errors.New("invalid credentials")
	}
}
