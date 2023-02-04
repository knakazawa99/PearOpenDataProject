package route

import (
	"api/http/handler"
	"context"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Example handler.Example
}

func NewHandler(ctx context.Context) (Handler, error) {
	return Handler{
		Example: handler.NewExample(),
	}, nil
}

func V1(handler Handler, e *echo.Echo) {
	v1 := e.Group("/v1")

	example := v1.Group("/example")
	example.GET("/:id", handler.Example.Get)
	example.POST("/:id", handler.Example.Post)

	return
}
