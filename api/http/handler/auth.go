package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"api/domain/entity"
	"api/http/request"
	"api/usecase"
)

type Auth interface {
	RequestEmail(ctx echo.Context) error
}

type auth struct {
	authUseCase usecase.Auth
}

func (a auth) RequestEmail(ctx echo.Context) error {
	req := &request.ReqeustEmail{}
	if err := ctx.Bind(req); err != nil {
		// TODO: error handling
	}
	requestEmailEntity, err := entity.NewAuth(req.Email)
	if err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errorMessage)
	}
	err = a.authUseCase.RequestEmail(requestEmailEntity)
	if err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusBadRequest, errorMessage)
	}
	return ctx.String(http.StatusOK, "Post Example")
}

func NewAuth(authUseCase usecase.Auth) Auth {
	return &auth{
		authUseCase: authUseCase,
	}
}
