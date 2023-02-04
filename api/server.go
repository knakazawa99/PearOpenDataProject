package main

import (
	"api/http/route"
	"context"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	e := echo.New()
	handler, _ := route.NewHandler(ctx)

	route.V1(handler, e)
	e.Logger.Fatal(e.Start(":1323"))
}
