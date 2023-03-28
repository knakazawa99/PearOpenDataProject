package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"api/domain/entity"
	"api/domain/entity/types"
	"api/http/request"
	"api/http/response"
	"api/usecase"
	"api/utils"
)

type Auth interface {
	RequestEmail(ctx echo.Context) error
	DownloadWithToken(ctx echo.Context) error
	AdminSignup(ctx echo.Context) error
	RegisterAdmin(ctx echo.Context) error
}

type auth struct {
	authUseCase usecase.Auth
}

func (a auth) RequestEmail(ctx echo.Context) error {
	req := &request.ReqeustEmail{}
	if err := ctx.Bind(req); err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errorMessage)
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

func (a auth) AdminSignup(ctx echo.Context) error {
	req := &request.AdminAuth{}
	if err := ctx.Bind(req); err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errorMessage)
	}
	requestAuthEntity, err := entity.NewAdminAuth(req.Email, req.Password)
	if err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errorMessage)
	}
	jwtToken, err := a.authUseCase.AdminSignUp(requestAuthEntity)
	if err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusBadRequest, errorMessage)
	}

	return ctx.JSON(http.StatusOK, response.AdminAuth{
		JWTToken: jwtToken,
	})
}

func (a auth) RegisterAdmin(ctx echo.Context) error {
	req := request.AdminAuth{}
	if err := ctx.Bind(req); err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errorMessage)
	}
	authEntity := entity.Auth{
		Email:    entity.Email(req.Email),
		Password: req.Password,
		Type:     types.TypeAdmin,
	}

	jwtToken := ctx.Get("jwtToken").(string)
	jwtKey := ctx.Get("jwtKey").(string)
	authorizationEntity := entity.Authorization{
		JWTKey:   jwtKey,
		JWTToken: jwtToken,
	}

	auth, err := a.authUseCase.SaveAdmin(authEntity, authorizationEntity)
	if err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusBadRequest, errorMessage)
	}

	return ctx.JSON(http.StatusCreated, auth)

}

func NewAuth(authUseCase usecase.Auth) Auth {
	return &auth{
		authUseCase: authUseCase,
	}
}
