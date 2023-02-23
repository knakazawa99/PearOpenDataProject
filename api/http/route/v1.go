package route

import (
	"context"

	"github.com/labstack/echo/v4"

	"api/domain/presenter"
	"api/http/handler"
	"api/infrastructure/notify"
	"api/infrastructure/repository"
	"api/usecase"
)

type Handler struct {
	Example handler.Example
	Auth    handler.Auth
	Pear    handler.PearData
}

func NewHandler(ctx context.Context) (Handler, error) {
	authRepository := repository.NewAuth()
	downloadPearRepository := repository.NewDownloadPear()
	pearRepository := repository.NewPear()

	emailSender := notify.NewEmailSender()

	pearVersionPresenter := presenter.NewPearVersion()

	authUseCase := usecase.NewAuth(authRepository, downloadPearRepository, emailSender)
	pearUseCase := usecase.NewPearData(pearRepository, pearVersionPresenter)
	return Handler{
		Example: handler.NewExample(),
		Auth:    handler.NewAuth(authUseCase),
		Pear:    handler.NewPearData(pearUseCase),
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

	pear := v1.Group("/pears")
	pear.GET("/", handler.Pear.GetPearVersions)

	return
}
