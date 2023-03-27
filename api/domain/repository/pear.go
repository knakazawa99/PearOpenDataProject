package repository

import (
	"gorm.io/gorm"

	"api/domain/entity"
)

type Pear interface {
	FindPears(db *gorm.DB) ([]entity.Pear, error)
	FindReleasedPears(db *gorm.DB) ([]entity.Pear, error)
	Update(db *gorm.DB, pear entity.Pear) error
	Create(db *gorm.DB, pear entity.Pear) (entity.Pear, error)
}
