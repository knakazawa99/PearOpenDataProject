package repository

import (
	"gorm.io/gorm"

	"api/domain/entity"
	"api/domain/repository"
)

type downloadPear struct {
}

func (d downloadPear) Find(db *gorm.DB, downloadPear entity.DownloadPear) (entity.DownloadPear, error) {
	//TODO implement me
	return entity.DownloadPear{}, nil
}

func NewDownloadPear() repository.DownloadPear {
	return &downloadPear{}
}
