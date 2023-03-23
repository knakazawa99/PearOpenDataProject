package response

import (
	"time"
)

type PearDataVersionOutput struct {
	Version     string    `json:"version"`
	ReleaseNote string    `json:"release_note"`
	CreatedAt   time.Time `json:"created_at"`
}

type PearAdminDataVersionOutput struct {
	ID             int       `json:"id"`
	FilePath       string    `json:"file_path"`
	Version        string    `json:"version"`
	ReleaseNote    string    `json:"release_note"`
	ReleaseComment string    `json:"release_comment"`
	ReleaseFlag    bool      `json:"release_flag"`
	CreatedAt      time.Time `json:"created_at"`
}
