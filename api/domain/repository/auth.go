package repository

import (
	"gorm.io/gorm"

	"api/domain/entity"
	"api/domain/entity/types"
)

type Auth interface {
	FindByEmail(db *gorm.DB, auth entity.Email) (entity.Auth, error)
	FindByType(db *gorm.DB, authType types.AuthType) ([]entity.Auth, error)
	SaveAuth(db *gorm.DB, auth entity.Auth) error
	Delete(db *gorm.DB, auth entity.Auth) error
}
