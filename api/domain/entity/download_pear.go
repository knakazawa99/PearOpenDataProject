package entity

import (
	"errors"

	"github.com/go-playground/validator/v10"

	"api/domain/entity/types"
)

type DownloadPear struct {
	AuthInfo Auth
	Version  string
	FileName string
}

func NewDownloadPear(email string, token string, version string) (DownloadPear, error) {
	validate := validator.New()
	downloadPear := &DownloadPear{
		AuthInfo: Auth{
			Email(email),
			token,
			types.TypeAdmin,
			"",
		},
		Version: version,
	}
	err := validate.Struct(downloadPear)
	if err != nil {
		//if _, ok := err.(*validator.InvalidValidationError); ok {
		//	// TODO: https://github.com/go-playground/validator/blob/master/_examples/simple/main.go
		//	//return RequestEmail{}, echo.NewHTTPError(http.StatusUnprocessableEntity, "")
		//}
		//errors.New("Please Correct Email Format")
		return DownloadPear{}, errors.New("please Correct Email Format")
	}
	return *downloadPear, nil
}
