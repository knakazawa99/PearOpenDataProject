package gormmodel

import (
	"time"
)

type GormAuthInformation struct {
	ID        int
	AuthType  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (GormAuthInformation) TableName() string {
	return "auth_information"
}
