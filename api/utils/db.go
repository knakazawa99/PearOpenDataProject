package utils

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetDBFromContext(ctx echo.Context) *gorm.DB {
	return ctx.Get("db").(*gorm.DB)
}
