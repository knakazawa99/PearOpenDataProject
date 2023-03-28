package entity

import (
	"errors"
	"regexp"
	"time"

	"api/domain/entity/types"
)

type Auth struct {
	ID        int
	Email     Email
	Token     string
	Type      types.AuthType
	User      AuthUser
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAuth(organization string, name string, email string) (Auth, error) {
	if organization == "" {
		return Auth{}, errors.New("organization should not be nil")
	}
	if name == "" {
		return Auth{}, errors.New("name should not be nil")
	}
	if match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, email); !match {
		return Auth{}, errors.New("please Correct Email Format")
	}
	auth := &Auth{
		Email: Email(email),
		User: AuthUser{
			Organization: organization,
			Name:         name,
		},
	}
	return *auth, nil
}

func NewAdminAuth(email string, password string) (Auth, error) {
	if email == "" {
		return Auth{}, errors.New("email should not be nil")
	}
	if password == "" {
		return Auth{}, errors.New("password should not be nil")
	}
	auth := &Auth{
		Email:    Email(email),
		Password: password,
	}
	return *auth, nil
}
