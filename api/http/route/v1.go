package route

import (
	"context"

	"github.com/labstack/echo/v4"

	"api/http/handler"
	"api/infrastructure/notify"
	"api/infrastructure/repository"
	"api/usecase"
)

type Handler struct {
	Example handler.Example
	Auth    handler.Auth
}

func NewHandler(ctx context.Context) (Handler, error) {
	authRepository := repository.NewAuth()
	emailSender := notify.NewEmailSender()
	authUseCase := usecase.NewAuth(authRepository, emailSender)
	return Handler{
		Example: handler.NewExample(),
		Auth:    handler.NewAuth(authUseCase),
	}, nil
}

func V1(handler Handler, e *echo.Echo) {

	v1 := e.Group("/v1")

	example := v1.Group("/example")
	example.GET("/:id", handler.Example.Get)
	example.POST("/:id", handler.Example.Post)

	auth := v1.Group("/auth")
	auth.POST("/notify/request", handler.Auth.RequestEmail)
	auth.POST("/download", handler.Auth.DownloadWithToken)

	return
}
