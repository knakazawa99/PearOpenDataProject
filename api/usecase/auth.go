package usecase

import (
	"errors"

	"gorm.io/gorm"

	"api/domain/entity"
	"api/domain/repository"
	"api/infrastructure/notify"
	"api/utils"
)

type Auth interface {
	RequestEmail(auth entity.Auth) error
	DownloadWithToken(inputDownloadPear entity.DownloadPear) (entity.DownloadPear, error)
	AdminSignUp(auth entity.Auth) (string, error)
}

type authInteractor struct {
	authRepository         repository.Auth
	downloadPearRepository repository.DownloadPear
	cacheRepository        repository.Cache
	emailSender            notify.EmailSender
}

func (a authInteractor) RequestEmail(authRequest entity.Auth) error {
	db, err := utils.ConnectDB()
	dbForClose, err := db.DB()
	defer dbForClose.Close()

	auth, err := a.authRepository.FindByEmail(db, authRequest.Email)
	if err != nil {
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			return err
		}
		auth.User = authRequest.User
	}
	auth.Token = utils.GenerateToken(string(auth.Email))
	if auth.Email == "" {
		auth.Email = authRequest.Email
	}

	if err := a.authRepository.SaveAuth(db, auth); err != nil {
		return err
	}
	messageContentWithToken := utils.GenerateMessageContentWithToken(auth.Token)
	if err = a.emailSender.Send(notify.EmailDTO{Email: string(auth.Email), MessageContent: messageContentWithToken}); err != nil {
		return err
	}
	return nil
}

func (a authInteractor) DownloadWithToken(inputDownloadPear entity.DownloadPear) (entity.DownloadPear, error) {
	db, err := utils.ConnectDB()
	dbForClose, err := db.DB()
	defer dbForClose.Close()
	auth, err := a.authRepository.FindByEmail(db, inputDownloadPear.AuthInfo.Email)
	if err != nil {
		return entity.DownloadPear{}, err
	}
	if inputDownloadPear.AuthInfo.Token != auth.Token {
		// TODO: Err をTokenが間違えている時に用いるエラーにする
		return entity.DownloadPear{}, errors.New("InvalidToken")
	}
	downloadPear, err := a.downloadPearRepository.Find(db, inputDownloadPear)
	if err != nil {
		return entity.DownloadPear{}, err
	}
	return downloadPear, nil

}

func (a authInteractor) AdminSignUp(requestAuth entity.Auth) (string, error) {
	db, _ := utils.ConnectDB()
	dbForClose, _ := db.DB()
	defer dbForClose.Close()
	auth, err := a.authRepository.FindByEmail(db, requestAuth.Email)
	if err != nil {
		return "", err
	}

	if err := utils.CheckHashPassword(auth.Password, requestAuth.Password); err != nil {
		return "", err
	}

	token := utils.GenerateJWT(string(requestAuth.Email))
	if err := a.cacheRepository.Set(string(requestAuth.Email), token, 60*60*24); err != nil {
		return "", err
	}

	return token, nil
}

func NewAuth(authRepository repository.Auth, downloadPearRepository repository.DownloadPear, cacheRepository repository.Cache, emailSender notify.EmailSender) Auth {
	return authInteractor{
		authRepository:         authRepository,
		downloadPearRepository: downloadPearRepository,
		cacheRepository:        cacheRepository,
		emailSender:            emailSender,
	}
}
