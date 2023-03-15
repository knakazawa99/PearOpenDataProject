package entity

import (
	"time"
)

type Pear struct {
	ID             int
	Version        string
	FilePath       string
	ReleaseNote    string
	ReleaseComment string
	ReleaseFlag    bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
