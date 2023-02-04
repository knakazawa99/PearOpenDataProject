package handler

import (
	"api/http/request"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Example interface {
	Get(ctx echo.Context) error
	Post(ctx echo.Context) error
}

type example struct {
}

func (e example) Get(ctx echo.Context) error {
	req := &request.SampleGetRequest{}
	if err := ctx.Bind(req); err != nil {
		// error handling
	}
	fmt.Println(fmt.Sprintf("%d, %s, %s", req.ID, req.Name, req.Remarks))
	return ctx.String(http.StatusOK, "Get Example")
}

func (e example) Post(ctx echo.Context) error {
	req := &request.SamplePostRequest{}
	if err := ctx.Bind(req); err != nil {
		// error handling
	}
	fmt.Println(fmt.Sprintf("%d, %s, %s", req.ID, req.Name, req.Remarks))
	return ctx.String(http.StatusOK, "Post Example")
}

func NewExample() Example {
	return &example{}
}
