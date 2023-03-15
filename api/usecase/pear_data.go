package usecase

import (
	"api/domain/presenter"
	"api/domain/repository"
	"api/http/response"
	"api/utils"
)

type Pear interface {
	GetDataVersions() ([]response.PearDataVersionOutput, error)
	GetAdminDataVersions() ([]response.PearAdminDataVersionOutput, error)
}

type pearDataInteractor struct {
	pearRepository       repository.Pear
	pearVersionPresenter presenter.PearVersion
}

func (p pearDataInteractor) GetDataVersions() ([]response.PearDataVersionOutput, error) {
	db, err := utils.ConnectDB()
	dbForClose, err := db.DB()
	defer dbForClose.Close()
	pears, err := p.pearRepository.FindReleasedPears(db)
	if err != nil {
		return []response.PearDataVersionOutput{}, err
	}
	return p.pearVersionPresenter.OutPutPearVersions(pears), nil
}

func (p pearDataInteractor) GetAdminDataVersions() ([]response.PearAdminDataVersionOutput, error) {
	db, err := utils.ConnectDB()
	dbForClose, err := db.DB()
	defer dbForClose.Close()
	pears, err := p.pearRepository.FindPears(db)
	if err != nil {
		return []response.PearAdminDataVersionOutput{}, err
	}
	return p.pearVersionPresenter.OutPutPearAdminVersions(pears), nil
}

func NewPearData(pearRepository repository.Pear, pearVersionPresenter presenter.PearVersion) Pear {
	return pearDataInteractor{
		pearRepository:       pearRepository,
		pearVersionPresenter: pearVersionPresenter,
	}
}
