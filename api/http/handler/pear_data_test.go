package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"api/http/response"
	"api/usecase"
)

func TestPearData_GetPearVersions(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/pears", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	mockPearUseCase := usecase.NewMockPear(ctrl)

	mockPearDataVersionOutputs := make([]response.PearDataVersionOutput, 2)
	mockPearDataVersionOutputs[0] = response.PearDataVersionOutput{
		Version:     "1.0.0",
		ReleaseNote: "release",
		CreatedAt:   time.Now(),
	}
	mockPearDataVersionOutputs[1] = response.PearDataVersionOutput{
		Version:     "1.0.1",
		ReleaseNote: "release",
		CreatedAt:   time.Now(),
	}
	mockPearUseCase.EXPECT().GetDataVersions().Return(mockPearDataVersionOutputs, nil)
	handler := NewPearData(mockPearUseCase)

	err := handler.GetPearVersions(c)
	assert.Nil(t, err)
}
