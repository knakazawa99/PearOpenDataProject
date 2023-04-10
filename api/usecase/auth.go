package usecase

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"api/domain/entity"
	"api/domain/entity/types"
	"api/domain/repository"
	"api/infrastructure/notify"
	"api/utils"
)

type Auth interface {
	RequestEmail(db *gorm.DB, auth entity.Auth) error
	DownloadWithToken(db *gorm.DB, inputDownloadPear entity.DownloadPear) (entity.DownloadPear, error)
	AdminSignUp(db *gorm.DB, auth entity.Auth) (string, error)
	SaveAdmin(db *gorm.DB, auth entity.Auth, authorizationEntity entity.Authorization) (entity.Auth, error)
	GetAdmin(db *gorm.DB, authorizationEntity entity.Authorization) ([]entity.Auth, error)
	DeleteAdmin(db *gorm.DB, auth entity.Auth, authorizationEntity entity.Authorization) error
}

type authInteractor struct {
	authRepository         repository.Auth
	downloadPearRepository repository.DownloadPear
	cacheRepository        repository.Cache
	emailSender            notify.EmailSender
}

func (a authInteractor) RequestEmail(db *gorm.DB, authRequest entity.Auth) error {
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

func (a authInteractor) DownloadWithToken(db *gorm.DB, inputDownloadPear entity.DownloadPear) (entity.DownloadPear, error) {
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

func (a authInteractor) AdminSignUp(db *gorm.DB, requestAuth entity.Auth) (string, error) {
	auth, err := a.authRepository.FindByEmail(db, requestAuth.Email)
	if err != nil {
		return "", err
	}

	if err := utils.CheckHashPassword(auth.Password, requestAuth.Password); err != nil {
		return "", err
	}

	token := utils.GenerateJWT(string(requestAuth.Email))
	if err := a.cacheRepository.Set(string(requestAuth.Email), token, time.Hour*24); err != nil {
		return "", err
	}

	return token, nil
}

func (a authInteractor) SaveAdmin(db *gorm.DB, auth entity.Auth, authorizationEntity entity.Authorization) (entity.Auth, error) {
	jwtToken, err := a.cacheRepository.Get(authorizationEntity.JWTKey)
	if err != nil {
		return entity.Auth{}, err
	}
	if jwtToken != authorizationEntity.JWTToken {
		return entity.Auth{}, errors.New("incorrect jwt token")
	}

	auth.Password, err = utils.PasswordEncrypt(auth.Password)
	if err != nil {
		return entity.Auth{}, err
	}

	if err := a.authRepository.SaveAuth(db, auth); err != nil {
		return entity.Auth{}, err
	}
	resultAuth, err := a.authRepository.FindByEmail(db, auth.Email)
	if err != nil {
		return entity.Auth{}, err
	}
	messageContent := "管理者のアカウントを追加しました。\n\nパスワードは管理者に聞いてください。"
	if err = a.emailSender.Send(notify.EmailDTO{Email: string(auth.Email), MessageContent: messageContent}); err != nil {
		return entity.Auth{}, err
	}

	return resultAuth, nil
}

func (a authInteractor) GetAdmin(db *gorm.DB, authorizationEntity entity.Authorization) ([]entity.Auth, error) {
	jwtToken, err := a.cacheRepository.Get(authorizationEntity.JWTKey)
	if err != nil {
		return []entity.Auth{}, err
	}
	if jwtToken != authorizationEntity.JWTToken {
		return []entity.Auth{}, errors.New("incorrect jwt token")
	}

	authEntities, err := a.authRepository.FindByType(db, types.TypeAdmin)
	if err != nil {
		return []entity.Auth{}, err
	}
	return authEntities, nil
}

func (a authInteractor) DeleteAdmin(db *gorm.DB, auth entity.Auth, authorizationEntity entity.Authorization) error {
	jwtToken, err := a.cacheRepository.Get(authorizationEntity.JWTKey)
	if err != nil {
		return err
	}
	if jwtToken != authorizationEntity.JWTToken {
		return errors.New("incorrect jwt token")
	}

	auth.Type = types.TypeAdmin
	if err := a.authRepository.Delete(db, auth); err != nil {
		return err
	}

	return nil
}

func NewAuth(authRepository repository.Auth, downloadPearRepository repository.DownloadPear, cacheRepository repository.Cache, emailSender notify.EmailSender) Auth {
	return authInteractor{
		authRepository:         authRepository,
		downloadPearRepository: downloadPearRepository,
		cacheRepository:        cacheRepository,
		emailSender:            emailSender,
	}
}
