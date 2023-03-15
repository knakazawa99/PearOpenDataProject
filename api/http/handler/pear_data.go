package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"api/usecase"
)

type PearData interface {
	GetPearVersions(ctx echo.Context) error
	GetAdminPearVersions(ctx echo.Context) error
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

func NewPearData(pearUseCase usecase.Pear) PearData {
	return &pearData{
		pearUseCase: pearUseCase,
	}
}
