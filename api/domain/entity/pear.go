package entity

import (
	"time"
)

type Pear struct {
	ID          int
	Version     string
	FilePath    string
	ReleaseNote string
	CreatedAt   time.Time
}
