package route

import (
	"api/http/handler"
	"context"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Example handler.Example
	Auth    handler.Auth
}

func NewHandler(ctx context.Context) (Handler, error) {
	return Handler{
		Example: handler.NewExample(),
		Auth:    handler.NewAuth(),
	}, nil
}

func V1(handler Handler, e *echo.Echo) {

	v1 := e.Group("/v1")

	example := v1.Group("/example")
	example.GET("/:id", handler.Example.Get)
	example.POST("/:id", handler.Example.Post)

	auth := v1.Group("/auth")
	auth.POST("/email/request", handler.Auth.RequestEmail)

	return
}
