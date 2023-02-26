package gormmodel

import (
	"time"

	"gorm.io/gorm"
)

type GormAdminInformation struct {
	gorm.Model
	ID                int
	AuthInformationID int
	Password          string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (GormAdminInformation) TableName() string {
	return "admin_information"
}
