package entity

import (
	"errors"

	"github.com/go-playground/validator/v10"

	"api/domain/entity/types"
)

type Auth struct {
	Email    Email
	Token    string
	Type     types.AuthType
	Password string
}

func NewAuth(email string) (Auth, error) {
	validate := validator.New()
	auth := &Auth{
		Email: Email(email),
	}
	err := validate.Struct(auth)
	if err != nil {
		//if _, ok := err.(*validator.InvalidValidationError); ok {
		//	// TODO: https://github.com/go-playground/validator/blob/master/_examples/simple/main.go
		//	//return RequestEmail{}, echo.NewHTTPError(http.StatusUnprocessableEntity, "")
		//}
		//errors.New("Please Correct Email Format")
		return Auth{}, errors.New("please Correct Email Format")
	}
	return *auth, nil
}
