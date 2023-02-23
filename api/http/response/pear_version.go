package response

import (
	"time"
)

type PearDataVersionOutput struct {
	Version     string    `json:"version"`
	ReleaseNote string    `json:"release_note"`
	CreatedAt   time.Time `json:"created_at"`
}
