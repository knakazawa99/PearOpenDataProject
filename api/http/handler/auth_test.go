package handler

import (
	"api/app/usecase"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRequestEmail(t *testing.T) {
	requestJson := `{"email":"test@gmail.com"}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	mockAuthUseCase := usecase.NewMockAuth(ctrl)
	h := NewAuth(mockAuthUseCase)

	err := h.RequestEmail(c)
	assert.Nil(t, err)
}

func TestRequestEmailFail(t *testing.T) {
	requestJson := `{"email":"test_gmail_com"}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	mockAuthUseCase := usecase.NewMockAuth(ctrl)
	h := NewAuth(mockAuthUseCase)

	err := h.RequestEmail(c)
	assert.NotNil(t, err)
}