package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"api/domain/entity"
	"api/http/request"
	"api/usecase"
)

type PearData interface {
	GetPearVersions(ctx echo.Context) error
	GetAdminPearVersions(ctx echo.Context) error
	UpdateAdminPear(ctx echo.Context) error
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
	if err := p.pearUseCase.UpdateAdminData(pearEntity); err != nil {
		errorMessage := fmt.Sprintf("error: %s", err)
		ctx.Logger().Error(errorMessage)
		return echo.NewHTTPError(http.StatusInternalServerError, errorMessage)
	}
	return ctx.JSON(http.StatusNoContent, "")
}

func NewPearData(pearUseCase usecase.Pear) PearData {
	return &pearData{
		pearUseCase: pearUseCase,
	}
}
