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
	mockEmailSender := notify.NewMockEmailSender(ctrl)
	mockReturnAuth := entity.Auth{
		Email: entity.Email("test@gmail.com"),
	}
	mockAuthRepository.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(mockReturnAuth, nil)
	mockAuthRepository.EXPECT().SaveAuth(gomock.Any(), gomock.Any()).Return(nil)
	mockEmailSender.EXPECT().Send(gomock.Any()).Return(nil)

	authInteractor := NewAuth(mockAuthRepository, mockEmailSender)
	authEntity := entity.Auth{
		Email: entity.Email("test@gmail.com"),
	}
	err := authInteractor.RequestEmail(authEntity)
	assert.Nil(t, err)
}
