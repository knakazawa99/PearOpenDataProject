package entity

import (
	"errors"
	"regexp"

	"api/domain/entity/types"
)

type Auth struct {
	Email    Email
	Token    string
	Type     types.AuthType
	Password string
}

func NewAuth(email string) (Auth, error) {
	if match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, email); !match {
		return Auth{}, errors.New("please Correct Email Format")
	}
	auth := &Auth{
		Email: Email(email),
	}
	return *auth, nil
}
