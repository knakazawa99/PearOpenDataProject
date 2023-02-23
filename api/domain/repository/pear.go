package repository

import (
	"gorm.io/gorm"

	"api/domain/entity"
)

type Pear interface {
	FindPears(db *gorm.DB) ([]entity.Pear, error)
}
