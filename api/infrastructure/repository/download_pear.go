package repository

import (
	"gorm.io/gorm"

	"api/domain/entity"
	"api/domain/repository"
	"api/infrastructure/gormmodel"
)

type downloadPear struct {
}

func (d downloadPear) Find(db *gorm.DB, downloadPear entity.DownloadPear) (entity.DownloadPear, error) {
	var gormPear gormmodel.GormPear
	if err := db.Where("version = ?", downloadPear.Version).Take(&gormPear).Error; err != nil {
		return entity.DownloadPear{}, err
	}

	return entity.DownloadPear{
		AuthInfo: downloadPear.AuthInfo,
		FileName: gormPear.FilePath,
		Version:  gormPear.Version,
	}, nil
}

func NewDownloadPear() repository.DownloadPear {
	return &downloadPear{}
}
