package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"api/domain/entity"
	"api/domain/entity/types"
	"api/domain/repository"
	"api/infrastructure/notify"
)

func TestAuthInteractor_RequestEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAuthRepository := repository.NewMockAuth(ctrl)
	mockDownloadPearRepository := repository.NewMockDownloadPear(ctrl)
	mockCacheRepository := repository.NewMockCache(ctrl)
	mockEmailSender := notify.NewMockEmailSender(ctrl)
	mockReturnAuth := entity.Auth{
		Email: entity.Email("test@gmail.com"),
	}
	mockAuthRepository.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(mockReturnAuth, nil)
	mockAuthRepository.EXPECT().SaveAuth(gomock.Any(), gomock.Any()).Return(nil)
	mockEmailSender.EXPECT().Send(gomock.Any()).Return(nil)

	authInteractor := NewAuth(mockAuthRepository, mockDownloadPearRepository, mockCacheRepository, mockEmailSender)
	authEntity := entity.Auth{
		Email: entity.Email("test@gmail.com"),
	}
	db := &gorm.DB{}
	err := authInteractor.RequestEmail(db, authEntity)
	assert.Nil(t, err)
}

func TestAuthInteractor_DownloadWithToken_InvalidToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAuthRepository := repository.NewMockAuth(ctrl)
	mockDownloadPearRepository := repository.NewMockDownloadPear(ctrl)
	mockCacheRepository := repository.NewMockCache(ctrl)
	mockEmailSender := notify.NewMockEmailSender(ctrl)
	mockReturnAuth := entity.Auth{
		Email: entity.Email("test@gmail.com"),
		Token: "hogehoge",
	}
	mockDownloadPear := entity.DownloadPear{
		AuthInfo: mockReturnAuth,
		Version:  "1.0.0",
		FileName: "test.zip",
	}
	mockAuthRepository.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(mockReturnAuth, nil)
	mockDownloadPearRepository.EXPECT().Find(gomock.Any(), gomock.Any()).Return(mockDownloadPear, nil)
	authInteractor := NewAuth(mockAuthRepository, mockDownloadPearRepository, mockCacheRepository, mockEmailSender)

	inputDownloadPear := entity.DownloadPear{
		AuthInfo: entity.Auth{
			Email: entity.Email("test&gmail.com"),
			Token: "hogehoge",
		},
		Version: "1.0.0",
	}
	expectedResult := mockDownloadPear
	db := &gorm.DB{}
	result, err := authInteractor.DownloadWithToken(db, inputDownloadPear)
	assert.Nil(t, err)
	assert.Equal(t, result.Version, expectedResult.Version)
}

func TestAuthInteractor_DownloadWithToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAuthRepository := repository.NewMockAuth(ctrl)
	mockDownloadPearRepository := repository.NewMockDownloadPear(ctrl)
	mockCacheRepository := repository.NewMockCache(ctrl)
	mockEmailSender := notify.NewMockEmailSender(ctrl)
	mockReturnAuth := entity.Auth{
		Email: entity.Email("test@gmail.com"),
		Token: "hoge1",
	}
	mockAuthRepository.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(mockReturnAuth, nil)
	authInteractor := NewAuth(mockAuthRepository, mockDownloadPearRepository, mockCacheRepository, mockEmailSender)

	inputDownloadPear := entity.DownloadPear{
		AuthInfo: entity.Auth{
			Email: entity.Email("test&gmail.com"),
			Token: "hoge2",
		},
		Version: "1.0.0",
	}
	db := &gorm.DB{}
	_, err := authInteractor.DownloadWithToken(db, inputDownloadPear)
	assert.NotNil(t, err)
}

func TestAuthInteractor_AdminSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAuthRepository := repository.NewMockAuth(ctrl)
	mockDownloadPearRepository := repository.NewMockDownloadPear(ctrl)
	mockCacheRepository := repository.NewMockCache(ctrl)
	mockEmailSender := notify.NewMockEmailSender(ctrl)
	mockReturnAuth := entity.Auth{
		Email:    entity.Email("test@gmail.com"),
		Token:    "hoge1",
		Password: "$2a$10$uOAkYkVZAxaU/nbbG34gQuTj7WX7LZ2AS/off8DfQpEA4gRz.esoC",
	}
	mockAuthRepository.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(mockReturnAuth, nil)
	mockCacheRepository.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	authInteractor := NewAuth(mockAuthRepository, mockDownloadPearRepository, mockCacheRepository, mockEmailSender)

	requestAuth := entity.Auth{
		Email:    entity.Email("test@gmail.com"),
		Token:    "hoge1",
		Password: "hogehoge",
	}

	db := &gorm.DB{}
	result, err := authInteractor.AdminSignUp(db, requestAuth)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestAuthInteractor_SaveAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAuthRepository := repository.NewMockAuth(ctrl)
	mockDownloadPearRepository := repository.NewMockDownloadPear(ctrl)
	mockCacheRepository := repository.NewMockCache(ctrl)
	mockEmailSender := notify.NewMockEmailSender(ctrl)
	mockAuthEntity := entity.Auth{
		Email:    entity.Email("test@gmail.com"),
		Token:    "hoge1",
		Password: "hogehoge",
		Type:     types.TypeAdmin,
	}
	mockAuthRepository.EXPECT().SaveAuth(gomock.Any(), gomock.Any()).Return(nil)
	mockAuthRepository.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(mockAuthEntity, nil)
	authInteractor := NewAuth(mockAuthRepository, mockDownloadPearRepository, mockCacheRepository, mockEmailSender)
	requestAuthEntity := entity.Auth{
		Email:    entity.Email("test@gmail.com"),
		Token:    "hoge1",
		Password: "hogehoge",
		Type:     types.TypeAdmin,
	}
	authorizationEntity := entity.Authorization{
		JWTKey:   "key",
		JWTToken: "token",
	}

	mockCacheRepository.EXPECT().Get(gomock.Any()).Return("token", nil)
	mockEmailSender.EXPECT().Send(gomock.Any()).Return(nil)
	db := &gorm.DB{}
	result, err := authInteractor.SaveAdmin(db, requestAuthEntity, authorizationEntity)
	assert.Nil(t, err)
	assert.Equal(t, result.Password, mockAuthEntity.Password)
}

func TestAuthInteractor_GetAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAuthRepository := repository.NewMockAuth(ctrl)
	mockDownloadPearRepository := repository.NewMockDownloadPear(ctrl)
	mockCacheRepository := repository.NewMockCache(ctrl)
	mockEmailSender := notify.NewMockEmailSender(ctrl)
	mockAuthEntities := make([]entity.Auth, 2)
	mockAuthEntities[0] = entity.Auth{
		Email:    entity.Email("test@gmail.com"),
		Token:    "hoge1",
		Password: "hogehoge",
		Type:     types.TypeAdmin,
	}
	mockAuthEntities[1] = entity.Auth{
		Email:    entity.Email("test2@gmail.com"),
		Token:    "hoge1",
		Password: "hogehoge",
		Type:     types.TypeAdmin,
	}

	mockAuthRepository.EXPECT().FindByType(gomock.Any(), gomock.Any()).Return(mockAuthEntities, nil)
	mockCacheRepository.EXPECT().Get(gomock.Any()).Return("token", nil)
	authInteractor := NewAuth(mockAuthRepository, mockDownloadPearRepository, mockCacheRepository, mockEmailSender)
	authorizationEntity := entity.Authorization{
		JWTKey:   "key",
		JWTToken: "token",
	}
	db := &gorm.DB{}
	result, err := authInteractor.GetAdmin(db, authorizationEntity)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
}

func TestAuthInteractor_DeleteAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAuthRepository := repository.NewMockAuth(ctrl)
	mockDownloadPearRepository := repository.NewMockDownloadPear(ctrl)
	mockCacheRepository := repository.NewMockCache(ctrl)
	mockEmailSender := notify.NewMockEmailSender(ctrl)

	mockAuthRepository.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
	mockCacheRepository.EXPECT().Get(gomock.Any()).Return("token", nil)
	authInteractor := NewAuth(mockAuthRepository, mockDownloadPearRepository, mockCacheRepository, mockEmailSender)

	authEntity := entity.Auth{
		Email:    entity.Email("test@gmail.com"),
		Token:    "hoge1",
		Password: "hogehoge",
		Type:     types.TypeAdmin,
	}
	authorizationEntity := entity.Authorization{
		JWTKey:   "key",
		JWTToken: "token",
	}

	db := &gorm.DB{}
	err := authInteractor.DeleteAdmin(db, authEntity, authorizationEntity)
	assert.Nil(t, err)
}
