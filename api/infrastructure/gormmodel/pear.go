package gormmodel

import (
	"time"
)

type GormPear struct {
	ID             int
	Version        string
	ReleaseNote    string
	ReleaseComment string
	ReleaseFlag    bool
	FilePath       string
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

func (GormPear) TableName() string {
	return "pears"
}
