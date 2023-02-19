package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"api/domain/entity"
	"api/usecase"
)

func TestRequestEmail(t *testing.T) {
	requestJson := `{"notify":"test@gmail.com"}`
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
	requestJson := `{"notify":"test_gmail_com"}`
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

func TestDownloadWithToken(t *testing.T) {
	requestJson := `{"email":"test@gmail_com", "token":"hogehoge", "version":"1.0.0"},`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	mockAuthUseCase := usecase.NewMockAuth(ctrl)
	mockAuthUseCase.EXPECT().DownloadWithToken(gomock.Any()).Return(entity.DownloadPear{
		AuthInfo: entity.Auth{},
		FileName: "test.zip",
		Version:  "1.0.0",
	}, nil)
	h := NewAuth(mockAuthUseCase)

	err := h.DownloadWithToken(c)
	assert.NotNil(t, err)
}
