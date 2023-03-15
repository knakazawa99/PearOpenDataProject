package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"api/domain/entity"
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
	err := authInteractor.RequestEmail(authEntity)
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
	result, err := authInteractor.DownloadWithToken(inputDownloadPear)
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
	_, err := authInteractor.DownloadWithToken(inputDownloadPear)
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

	result, err := authInteractor.AdminSignUp(requestAuth)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
