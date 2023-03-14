package gormmodel

import (
	"time"
)

type GormAuthUser struct {
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
