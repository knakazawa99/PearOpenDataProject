package gormmodel

import (
	"time"
)

type GormToken struct {
	ID                int
	AuthInformationID int
	Token             string
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}

func (GormToken) TableName() string {
	return "tokens"
}
