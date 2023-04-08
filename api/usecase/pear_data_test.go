package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"api/domain/entity"
	"api/domain/presenter"
	"api/domain/repository"
	"api/http/response"
)

func TestPearDataInteractor_GetDataVersions(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPearRepository := repository.NewMockPear(ctrl)
	mockCacheRepository := repository.NewMockCache(ctrl)
	mockPears := make([]entity.Pear, 2)
	mockPears[0] = entity.Pear{ID: 1, Version: "1.0.0", FilePath: "hoge/data.zip"}
	mockPears[1] = entity.Pear{ID: 2, Version: "1.0.0", FilePath: "hoge/data.zip"}

	mockPearRepository.EXPECT().FindPears(gomock.Any()).Return(mockPears, nil)
	pearVersionPresenter := presenter.NewPearVersion()
	pearDataInteractor := NewPearData(mockPearRepository, mockCacheRepository, pearVersionPresenter)
	db := &gorm.DB{}
	pears, err := pearDataInteractor.GetDataVersions(db)
	assert.IsType(t, []response.PearDataVersionOutput{}, pears)
	assert.Nil(t, err)
}

func TestPearDataInteractor_UpdateAdminData(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPearRepository := repository.NewMockPear(ctrl)
	mockCacheRepository := repository.NewMockCache(ctrl)

	mockPearRepository.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
	mockCacheRepository.EXPECT().Get(gomock.Any()).Return("token", nil)
	pearVersionPresenter := presenter.NewPearVersion()
	pearDataInteractor := NewPearData(mockPearRepository, mockCacheRepository, pearVersionPresenter)

	pearEntity := entity.Pear{}
	authorizationEntity := entity.Authorization{
		JWTKey:   "key",
		JWTToken: "token",
	}
	db := &gorm.DB{}
	err := pearDataInteractor.UpdateAdminData(db, pearEntity, authorizationEntity)
	assert.Nil(t, err)
}

func TestPearDataInteractor_CreateData(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPearRepository := repository.NewMockPear(ctrl)
	mockCacheRepository := repository.NewMockCache(ctrl)
	pearVersionPresenter := presenter.NewPearVersion()
	pearDataInteractor := NewPearData(mockPearRepository, mockCacheRepository, pearVersionPresenter)

	pearEntity := entity.Pear{
		ReleaseNote: "release_note",
	}
	authorizationEntity := entity.Authorization{
		JWTKey:   "key",
		JWTToken: "token",
	}

	mockPearRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(entity.Pear{ReleaseNote: "release_note"}, nil)
	mockCacheRepository.EXPECT().Get(gomock.Any()).Return("token", nil)

	db := &gorm.DB{}
	createdPearEntity, err := pearDataInteractor.CreateData(db, pearEntity, authorizationEntity)
	assert.Nil(t, err)
	assert.Equal(t, createdPearEntity.ReleaseNote, pearEntity.ReleaseNote)
}
