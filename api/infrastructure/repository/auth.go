package repository

import (
	"gorm.io/gorm"

	"api/domain/entity"
	"api/domain/entity/types"
	"api/domain/repository"
	"api/infrastructure/gormmodel"
)

type auth struct {
}

func (a auth) SaveAuth(db *gorm.DB, auth entity.Auth) error {
	var gormAuthInformation gormmodel.GormAuthInformation
	if err := db.Where("email = ?", auth.Email).Take(&gormAuthInformation).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			gormAuthInformation = gormmodel.GormAuthInformation{
				Email:    string(auth.Email),
				AuthType: string(types.TypeAdmin),
			}
			if err := db.Save(&gormAuthInformation).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}
	if auth.Type == types.TypeAdmin {
		gormToken := gormmodel.GormToken{
			Token:             auth.Token,
			AuthInformationID: gormAuthInformation.ID,
		}
		if err := db.Save(&gormToken).Error; err != nil {
			return err
		}
	} else if auth.Type == types.TypeUser {
		gormAdminInformation := gormmodel.GormAdminInformation{
			AuthInformationID: gormAuthInformation.ID,
			Password:          auth.Password,
		}
		if err := db.Save(&gormAdminInformation).Error; err != nil {
			return err
		}
	} else {
	}
	return nil
}

func (a auth) FindByEmail(db *gorm.DB, email entity.Email) (entity.Auth, error) {
	var gormAuthInformation gormmodel.GormAuthInformation
	var gormToken gormmodel.GormToken
	if err := db.Where("email = ?", email).Take(&gormAuthInformation).Error; err != nil {

	}
	if err := db.Where("auth_information_id = ?", gormAuthInformation.ID).Take(&gormToken).Error; err != nil {

	}

	return entity.Auth{
		Email: entity.Email(gormAuthInformation.Email),
		Token: gormToken.Token,
		Type:  types.AuthType(gormAuthInformation.AuthType),
	}, nil
}

func NewAuth() repository.Auth {
	return &auth{}
}
