package usecase

import (
	"gorm.io/gorm"

	"api/domain/presenter"
	"api/domain/repository"
)

type Pear interface {
	GetDataVersions() ([]presenter.PearDataVersionOutput, error)
}

type pearDataInteractor struct {
	pearRepository       repository.Pear
	pearVersionPresenter presenter.PearVersion
}

func (p pearDataInteractor) GetDataVersions() ([]presenter.PearDataVersionOutput, error) {
	db := &gorm.DB{}
	pears, err := p.pearRepository.FindPears(db)
	if err != nil {
		return []presenter.PearDataVersionOutput{}, err
	}
	return p.pearVersionPresenter.OutPutPearVersions(pears), nil
}

func NewPearData(pearRepository repository.Pear, pearVersionPresenter presenter.PearVersion) Pear {
	return pearDataInteractor{
		pearRepository:       pearRepository,
		pearVersionPresenter: pearVersionPresenter,
	}
}
