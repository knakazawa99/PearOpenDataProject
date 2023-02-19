package usecase

import (
	"gorm.io/gorm"

	"api/domain/entity"
	"api/domain/repository"
	"api/infrastructure/notify"
	"api/utils"
)

type Auth interface {
	RequestEmail(auth entity.Auth) error
}

type authInteractor struct {
	authRepository repository.Auth
	emailSender    notify.EmailSender
}

func (a authInteractor) RequestEmail(authRequest entity.Auth) error {
	db := &gorm.DB{}
	auth, err := a.authRepository.FindByEmail(db, authRequest.Email)
	if err != nil {
		return err
	}
	auth.Token = utils.GenerateToken(string(auth.Email))
	if auth.Email != "" {
		auth.Email = authRequest.Email
	}

	if err := a.authRepository.SaveAuth(db, authRequest); err != nil {
		return err
	}
	messageContentWithToken := utils.GenerateMessageContentWithToken(auth.Token)
	if err = a.emailSender.Send(notify.EmailDTO{Email: string(auth.Email), MessageContent: messageContentWithToken}); err != nil {
		return err
	}
	return nil
}

func NewAuth(authRepository repository.Auth, emailSender notify.EmailSender) Auth {
	return authInteractor{
		authRepository: authRepository,
		emailSender:    emailSender,
	}
}
