package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"api/domain/entity"
	"api/http/request"
	"api/usecase"
)

type PearData interface {
	GetPearVersions(ctx echo.Context) error
	GetAdminPearVersions(ctx echo.Context) error
	UpdateAdminPear(ctx echo.Context) error
	UploadPear(ctx echo.Context) error
}

type pearData struct {
	pearUseCase usecase.Pear
}

func (p pearData) GetPearVersions(ctx echo.Context) error {
	pearVersion, err := p.pearUseCase.GetDataVersions()
	if err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusInternalServerError, errorMessage)
	}
	return ctx.JSON(http.StatusOK, pearVersion)
}

func (p pearData) GetAdminPearVersions(ctx echo.Context) error {
	pearVersion, err := p.pearUseCase.GetAdminDataVersions()
	if err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusInternalServerError, errorMessage)
	}
	return ctx.JSON(http.StatusOK, pearVersion)
}

func (p pearData) UpdateAdminPear(ctx echo.Context) error {
	req := &request.PearUpdate{}
	jwtToken := ctx.Get("jwtToken").(string)
	jwtKey := ctx.Get("jwtKey").(string)
	if err := ctx.Bind(req); err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errorMessage)
	}
	pearEntity := entity.Pear{
		ID:             req.ID,
		ReleaseNote:    req.ReleaseNote,
		ReleaseComment: req.ReleaseComment,
		ReleaseFlag:    req.ReleaseFlag,
	}

	authorizationEntity := entity.Authorization{
		JWTKey:   jwtKey,
		JWTToken: jwtToken,
	}
	if err := p.pearUseCase.UpdateAdminData(pearEntity, authorizationEntity); err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusInternalServerError, errorMessage)
	}
	return ctx.JSON(http.StatusNoContent, "")
}

func (p pearData) UploadPear(ctx echo.Context) error {
	req := &request.PearCreate{}
	jwtToken := ctx.Get("jwtToken").(string)
	jwtKey := ctx.Get("jwtKey").(string)
	if err := ctx.Bind(req); err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errorMessage)
	}
	pearEntity := entity.Pear{
		ReleaseNote:    req.ReleaseNote,
		ReleaseComment: req.ReleaseComment,
		ReleaseFlag:    req.ReleaseFlag,
		Version:        req.Version,
		FilePath:       "/var/pear/data/" + req.Version + ".zip",
	}

	authorizationEntity := entity.Authorization{
		JWTKey:   jwtKey,
		JWTToken: jwtToken,
	}

	createdPearEntity, err := p.pearUseCase.CreateData(pearEntity, authorizationEntity)
	if err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusInternalServerError, errorMessage)
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(pearEntity.FilePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy file contents to destination file
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, createdPearEntity)
}

func NewPearData(pearUseCase usecase.Pear) PearData {
	return &pearData{
		pearUseCase: pearUseCase,
	}
}
