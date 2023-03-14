package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"api/domain/entity"
	"api/http/request"
	"api/usecase"
	"api/utils"
)

type Auth interface {
	RequestEmail(ctx echo.Context) error
	DownloadWithToken(ctx echo.Context) error
}

type auth struct {
	authUseCase usecase.Auth
}

func (a auth) RequestEmail(ctx echo.Context) error {
	req := &request.ReqeustEmail{}
	if err := ctx.Bind(req); err != nil {
		// TODO: error handling
		fmt.Println("Err: ", err)
	}
	requestEmailEntity, err := entity.NewAuth(req.Organization, req.Name, req.Email)
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

func (a auth) DownloadWithToken(ctx echo.Context) error {
	req := &request.TokenWithDownload{}
	req.Token = ctx.QueryParam("token")
	req.Version = ctx.QueryParam("version")
	req.Email = ctx.QueryParam("email")

	requestDownLoadWithTokenEntity, err := entity.NewDownloadPear(req.Email, req.Token, req.Version)
	if err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errorMessage)
	}
	downloadPear, err := a.authUseCase.DownloadWithToken(requestDownLoadWithTokenEntity)
	if err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusBadRequest, errorMessage)
	}
	response := ctx.Response()
	response.Header().Set(echo.HeaderAccessControlExposeHeaders, "Content-Disposition")
	return ctx.Attachment(downloadPear.FileName, utils.GenerateOutPutFileName(downloadPear.FileName, downloadPear.Version))
}

func NewAuth(authUseCase usecase.Auth) Auth {
	return &auth{
		authUseCase: authUseCase,
	}
}
