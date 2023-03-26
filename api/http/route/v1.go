package route

import (
	"context"

	"github.com/labstack/echo/v4"

	"api/domain/presenter"
	"api/http/handler"
	"api/http/middleware"
	"api/infrastructure/notify"
	"api/infrastructure/repository"
	"api/usecase"
)

type Handler struct {
	Example handler.Example
	Auth    handler.Auth
	Pear    handler.PearData
}

type Middleware struct {
	Auth middleware.Auth
}

func NewHandler(ctx context.Context) (Handler, error) {
	authRepository := repository.NewAuth()
	downloadPearRepository := repository.NewDownloadPear()
	pearRepository := repository.NewPear()
	cacheRepository := repository.NewCache()

	emailSender := notify.NewEmailSender()

	pearVersionPresenter := presenter.NewPearVersion()

	authUseCase := usecase.NewAuth(authRepository, downloadPearRepository, cacheRepository, emailSender)
	pearUseCase := usecase.NewPearData(pearRepository, cacheRepository, pearVersionPresenter)
	return Handler{
		Example: handler.NewExample(),
		Auth:    handler.NewAuth(authUseCase),
		Pear:    handler.NewPearData(pearUseCase),
	}, nil
}

func NewMiddleware() (Middleware, error) {
	return Middleware{
		Auth: middleware.NewAuth(),
	}, nil
}

func V1(handler Handler, m Middleware, e *echo.Echo) {

	v1 := e.Group("/v1")

	example := v1.Group("/example")
	example.GET("/:id", handler.Example.Get)
	example.POST("/:id", handler.Example.Post)

	auth := v1.Group("/auth")
	auth.POST("/notify/request", handler.Auth.RequestEmail)
	auth.GET("/download", handler.Auth.DownloadWithToken)

	pear := v1.Group("/pears")
	pear.GET("/", handler.Pear.GetPearVersions)

	admin := v1.Group("/admin")
	admin.POST("/signup", handler.Auth.AdminSignup)

	admin.GET("/versions", handler.Pear.GetAdminPearVersions)
	admin.PUT("/versions/:id", handler.Pear.UpdateAdminPear, m.Auth.Auth)
	admin.POST("/versions/", handler.Pear.UploadPear, m.Auth.Auth)

	return
}
