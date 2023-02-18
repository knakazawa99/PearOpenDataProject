package repository

import (
	"gorm.io/gorm"

	"api/domain/entity"
)

type Auth interface {
	FindByEmail(db *gorm.DB, auth entity.Email) (entity.Auth, error)
	SaveAuth(db *gorm.DB, auth entity.Auth) error
}
