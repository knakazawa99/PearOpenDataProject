package repository

import (
	"gorm.io/gorm"

	"api/domain/entity"
	"api/domain/repository"
)

type pear struct {
}

func (p pear) FindPears(db *gorm.DB) ([]entity.Pear, error) {
	//TODO implement me
	return []entity.Pear{}, nil
}

func NewPear() repository.Pear {
	return &pear{}
}
