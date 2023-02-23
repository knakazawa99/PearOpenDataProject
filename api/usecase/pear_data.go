package usecase

import (
	"gorm.io/gorm"

	"api/domain/presenter"
	"api/domain/repository"
	"api/http/response"
)

type Pear interface {
	GetDataVersions() ([]response.PearDataVersionOutput, error)
}

type pearDataInteractor struct {
	pearRepository       repository.Pear
	pearVersionPresenter presenter.PearVersion
}

func (p pearDataInteractor) GetDataVersions() ([]response.PearDataVersionOutput, error) {
	db := &gorm.DB{}
	pears, err := p.pearRepository.FindPears(db)
	if err != nil {
		return []response.PearDataVersionOutput{}, err
	}
	return p.pearVersionPresenter.OutPutPearVersions(pears), nil
}

func NewPearData(pearRepository repository.Pear, pearVersionPresenter presenter.PearVersion) Pear {
	return pearDataInteractor{
		pearRepository:       pearRepository,
		pearVersionPresenter: pearVersionPresenter,
	}
}
