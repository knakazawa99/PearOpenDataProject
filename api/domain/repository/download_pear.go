package repository

import (
	"gorm.io/gorm"

	"api/domain/entity"
)

type DownloadPear interface {
	Find(db *gorm.DB, downloadPear entity.DownloadPear) (entity.DownloadPear, error)
}
