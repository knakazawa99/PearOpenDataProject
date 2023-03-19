package usecase

import (
	"api/domain/entity"
	"api/domain/presenter"
	"api/domain/repository"
	"api/http/response"
	"api/utils"
)

type Pear interface {
	GetDataVersions() ([]response.PearDataVersionOutput, error)
	GetAdminDataVersions() ([]response.PearAdminDataVersionOutput, error)
	UpdateAdminData(pearEntity entity.Pear) error
}

type pearDataInteractor struct {
	pearRepository       repository.Pear
	pearVersionPresenter presenter.PearVersion
}

func (p pearDataInteractor) GetDataVersions() ([]response.PearDataVersionOutput, error) {
	db, _ := utils.ConnectDB()
	dbForClose, _ := db.DB()
	defer dbForClose.Close()
	pears, err := p.pearRepository.FindReleasedPears(db)
	if err != nil {
		return []response.PearDataVersionOutput{}, err
	}
	return p.pearVersionPresenter.OutPutPearVersions(pears), nil
}

func (p pearDataInteractor) GetAdminDataVersions() ([]response.PearAdminDataVersionOutput, error) {
	db, _ := utils.ConnectDB()
	dbForClose, _ := db.DB()
	defer dbForClose.Close()
	pears, err := p.pearRepository.FindPears(db)
	if err != nil {
		return []response.PearAdminDataVersionOutput{}, err
	}
	return p.pearVersionPresenter.OutPutPearAdminVersions(pears), nil
}

func (p pearDataInteractor) UpdateAdminData(pearEntity entity.Pear) error {
	db, _ := utils.ConnectDB()
	dbForClose, _ := db.DB()
	defer dbForClose.Close()

	if err := p.pearRepository.Update(db, pearEntity); err != nil {
		return err
	}
	return nil
}

func NewPearData(pearRepository repository.Pear, pearVersionPresenter presenter.PearVersion) Pear {
	return pearDataInteractor{
		pearRepository:       pearRepository,
		pearVersionPresenter: pearVersionPresenter,
	}
}
