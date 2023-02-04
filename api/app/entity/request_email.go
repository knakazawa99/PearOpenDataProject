package entity

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type RequestEmail struct {
	Email string `validate:"required,email"`
}

func NewRequestEmail(email string) (RequestEmail, error) {
	validate = validator.New()
	requestEmail := &RequestEmail{
		Email: email,
	}
	err := validate.Struct(requestEmail)
	if err != nil {
		//if _, ok := err.(*validator.InvalidValidationError); ok {
		//	// TODO: https://github.com/go-playground/validator/blob/master/_examples/simple/main.go
		//	//return RequestEmail{}, echo.NewHTTPError(http.StatusUnprocessableEntity, "")
		//}
		//errors.New("Please Correct Email Format")
		return RequestEmail{}, errors.New("please Correct Email Format")
	}
	return *requestEmail, nil
}
