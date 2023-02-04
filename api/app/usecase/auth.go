package usecase

import "api/app/entity"

type Auth interface {
	RequestEmail(entity.RequestEmail) error
}

type auth struct {
}

func (a auth) RequestEmail(entity.RequestEmail) error {
	//TODO implement me
	panic("implement me")
}

func NewAuth() Auth {
	return auth{}
}
