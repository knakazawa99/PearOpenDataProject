package handler

import (
	"api/app/entity"
	"api/http/request"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Auth interface {
	RequestEmail(ctx echo.Context) error
}

type auth struct {
}

func (a auth) RequestEmail(ctx echo.Context) error {
	req := &request.ReqeustEmail{}
	if err := ctx.Bind(req); err != nil {
		// error handling
	}
	requestEmailEntity, err := entity.NewRequestEmail(req.Email)
	if err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errorMessage)
	}
	fmt.Println(requestEmailEntity)
	return ctx.String(http.StatusOK, "Post Example")
}

func NewAuth() Auth {
	return &auth{}
}
