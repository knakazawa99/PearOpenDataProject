package gormmodel

import (
	"time"
)

type GormAdminInformation struct {
	ID                int
	AuthInformationID int
	Password          string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (GormAdminInformation) TableName() string {
	return "admin_information"
}
