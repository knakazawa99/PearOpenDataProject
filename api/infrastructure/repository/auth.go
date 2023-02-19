package repository

import (
	"gorm.io/gorm"

	"api/domain/entity"
	"api/domain/repository"
)

type auth struct {
}

func (a auth) FindByEmail(db *gorm.DB, email entity.Email) (entity.Auth, error) {
	return entity.Auth{}, nil
}

func NewAuth() repository.Auth {
	return &auth{}
}
