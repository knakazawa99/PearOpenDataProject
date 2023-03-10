package repository

import (
	"gorm.io/gorm"

	"api/domain/entity"
	"api/domain/repository"
	"api/infrastructure/gormmodel"
)

type pear struct {
}

func (p pear) FindPears(db *gorm.DB) ([]entity.Pear, error) {
	var gormPears []gormmodel.GormPear
	if err := db.Find(&gormPears).Error; err != nil {
		return []entity.Pear{}, err
	}
	pears := make([]entity.Pear, len(gormPears))
	for i := range gormPears {
		pears[i] = entity.Pear{
			FilePath:    gormPears[i].FilePath,
			Version:     gormPears[i].Version,
			ReleaseNote: gormPears[i].ReleaseNote,
			CreatedAt:   *gormPears[i].CreatedAt,
		}
	}
	return pears, nil
}

func NewPear() repository.Pear {
	return &pear{}
}
