package repository

import (
	"errors"

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
				AuthType: string(auth.Type),
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
				if err := db.Create(&gormAuthUser).Error; err != nil {
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
	} else if gormAuthInformation.AuthType == string(types.TypeAdmin) {
		var gormAdminInformation gormmodel.GormAdminInformation
		if err := db.Where("auth_information_id = ?", gormAuthInformation.ID).Take(&gormAdminInformation).Error; err != nil {
			if err.Error() == gorm.ErrRecordNotFound.Error() {
				gormAdminInformation.Password = auth.Password
				gormAdminInformation.AuthInformationID = gormAuthInformation.ID
				if err := db.Create(&gormAdminInformation).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			gormAdminInformation.Password = auth.Password
			if err := db.Save(&gormAdminInformation).Error; err != nil {
				return err
			}
		}
	} else {
		return errors.New("auth type should be admin or user")
	}
	return nil
}

func (a auth) FindByEmail(db *gorm.DB, email entity.Email) (entity.Auth, error) {
	var gormAuthInformation gormmodel.GormAuthInformation

	if err := db.Where("email = ?", email).Take(&gormAuthInformation).Error; err != nil {
		return entity.Auth{}, err
	}

	if gormAuthInformation.AuthType == string(types.TypeAdmin) {
		var gormAdminInformation gormmodel.GormAdminInformation
		if err := db.Where("auth_information_id = ?", gormAuthInformation.ID).Take(&gormAdminInformation).Error; err != nil {
			return entity.Auth{}, err
		}
		return entity.Auth{
			ID:        gormAuthInformation.ID,
			Email:     entity.Email(gormAuthInformation.Email),
			Type:      types.AuthType(gormAuthInformation.AuthType),
			Password:  gormAdminInformation.Password,
			CreatedAt: gormAuthInformation.CreatedAt,
			UpdatedAt: gormAuthInformation.UpdatedAt,
		}, nil
	}

	var gormToken gormmodel.GormToken
	if err := db.Where("auth_information_id = ?", gormAuthInformation.ID).Take(&gormToken).Error; err != nil {
		return entity.Auth{}, err
	}

	return entity.Auth{
		Email: entity.Email(gormAuthInformation.Email),
		Token: gormToken.Token,
		Type:  types.AuthType(gormAuthInformation.AuthType),
	}, nil
}

func (a auth) FindByType(db *gorm.DB, authType types.AuthType) ([]entity.Auth, error) {
	var gormAuthInformation []gormmodel.GormAuthInformation
	if err := db.Where("auth_type = ?", authType).Find(&gormAuthInformation).Error; err != nil {
		return []entity.Auth{}, err
	}
	authEntities := make([]entity.Auth, len(gormAuthInformation))
	for i, gormAuth := range gormAuthInformation {
		authEntities[i] = entity.Auth{
			ID:        gormAuth.ID,
			Email:     entity.Email(gormAuth.Email),
			CreatedAt: gormAuth.CreatedAt,
			UpdatedAt: gormAuth.UpdatedAt,
		}
	}

	return authEntities, nil
}

func (a auth) Delete(db *gorm.DB, auth entity.Auth) error {
	if auth.Type == types.TypeAdmin {
		var gormAdminInformation gormmodel.GormAdminInformation
		if err := db.Where("auth_information_id = ?", auth.ID).Take(&gormAdminInformation).Error; err != nil {
			return err
		}
		if err := db.Delete(&gormmodel.GormAdminInformation{}, gormAdminInformation.ID).Error; err != nil {
			return err
		}
		if err := db.Delete(&gormmodel.GormAuthInformation{}, auth.ID).Error; err != nil {
			return err
		}
		return nil
	}

	return errors.New("this func has not been implemented")
}

func NewAuth() repository.Auth {
	return &auth{}
}
