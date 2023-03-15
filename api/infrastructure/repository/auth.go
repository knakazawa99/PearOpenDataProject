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
				AuthType: string(types.TypeUser),
			}
			if err := db.Create(&gormAuthInformation).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}
	if gormAuthInformation.AuthType == string(types.TypeUser) {
		var gormToken gormmodel.GormToken
		var gormAuthUser gormmodel.GormAuthUser

		if err := db.Where("auth_information_id = ?", gormAuthInformation.ID).Take(&gormToken).Error; err != nil {
			if err.Error() == gorm.ErrRecordNotFound.Error() {
				gormToken.Token = auth.Token
				gormToken.AuthInformationID = gormAuthInformation.ID
				if err := db.Create(&gormToken).Error; err != nil {
					return err
				}
				gormAuthUser.AuthInformationID = gormAuthInformation.ID
				gormAuthUser.Organization = auth.User.Organization
				gormAuthUser.Name = auth.User.Name
				if err := db.Debug().Create(&gormAuthUser).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			gormToken.Token = auth.Token
			if err := db.Save(&gormToken).Error; err != nil {
				return err
			}
		}
	} else if auth.Type == types.TypeAdmin {
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
		return entity.Auth{}, err
	}
	if err := db.Where("auth_information_id = ?", gormAuthInformation.ID).Take(&gormToken).Error; err != nil {
		return entity.Auth{}, err
	}

	if gormAuthInformation.AuthType == string(types.TypeAdmin) {
		var gormAdminInformation gormmodel.GormAdminInformation
		if err := db.Where("auth_information_id = ?", gormAuthInformation.ID).Take(&gormAdminInformation).Error; err != nil {
			return entity.Auth{}, err
		}
		return entity.Auth{
			Email:    entity.Email(gormAuthInformation.Email),
			Token:    gormToken.Token,
			Type:     types.AuthType(gormAuthInformation.AuthType),
			Password: gormAdminInformation.Password,
		}, nil
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
