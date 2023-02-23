package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"api/domain/entity"
	"api/domain/presenter"
	"api/domain/repository"
	"api/http/response"
)

func TestPearDataInteractor_GetDataVersions(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPearRepository := repository.NewMockPear(ctrl)
	mockPears := make([]entity.Pear, 2)
	mockPears[0] = entity.Pear{ID: 1, Version: "1.0.0", FilePath: "hoge/data.zip"}
	mockPears[1] = entity.Pear{ID: 2, Version: "1.0.0", FilePath: "hoge/data.zip"}

	mockPearRepository.EXPECT().FindPears(gomock.Any()).Return(mockPears, nil)
	pearVersionPresenter := presenter.NewPearVersion()
	pearDataInteractor := NewPearData(mockPearRepository, pearVersionPresenter)
	pears, err := pearDataInteractor.GetDataVersions()
	assert.IsType(t, []response.PearDataVersionOutput{}, pears)
	assert.Nil(t, err)
}
