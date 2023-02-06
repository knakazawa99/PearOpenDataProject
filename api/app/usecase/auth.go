package usecase

import "api/app/entity"

type Auth interface {
	RequestEmail(entity.RequestEmail) error
}

type authInteractor struct {
}

func (a authInteractor) RequestEmail(entity.RequestEmail) error {
	/*
		1. Emailを取得
		2. Emailがなければ登録
		3. Tokenを生成
		4. Emailに対してTokenを生成
	*/
	panic("implement me")
}

func NewAuth() Auth {
	return authInteractor{}
}
