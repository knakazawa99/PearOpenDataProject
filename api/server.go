package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"api/http/route"
)

func main() {
	ctx := context.Background()
	e := echo.New()
	e.Use(middleware.CORS())

	handler, _ := route.NewHandler(ctx)
	middleware, _ := route.NewMiddleware()
	route.V1(handler, middleware, e)
	e.Logger.Fatal(e.Start(":80"))
}
