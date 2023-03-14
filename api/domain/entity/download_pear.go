package entity

import (
	"errors"
	"regexp"

	"api/domain/entity/types"
)

type DownloadPear struct {
	AuthInfo Auth
	Version  string
	FileName string
}

func NewDownloadPear(email string, token string, version string) (DownloadPear, error) {
	if match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, email); !match {
		return DownloadPear{}, errors.New("please Correct Email Format")
	}

	if token == "" {
		return DownloadPear{}, errors.New("token should not be nil")
	}

	if version == "" {
		return DownloadPear{}, errors.New("version should not be nil")
	}
	downloadPear := &DownloadPear{
		AuthInfo: Auth{
			Email(email),
			token,
			types.TypeAdmin,
			AuthUser{},
			"",
		},
		Version: version,
	}
	return *downloadPear, nil
}
