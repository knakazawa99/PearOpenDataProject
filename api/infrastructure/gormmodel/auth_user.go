package gormmodel

import (
	"time"

	"gorm.io/gorm"
)

type GormAuthUser struct {
	gorm.Model
	ID                int
	AuthInformationID int
	Organization      string
	Name              string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (GormAuthUser) TableName() string {
	return "auth_users"
}
