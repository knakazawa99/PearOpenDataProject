package usecase

import (
	"errors"

	"api/domain/entity"
	"api/domain/presenter"
	"api/domain/repository"
	"api/http/response"
	"api/utils"
)

type Pear interface {
	GetDataVersions() ([]response.PearDataVersionOutput, error)
	GetAdminDataVersions() ([]response.PearAdminDataVersionOutput, error)
	UpdateAdminData(pearEntity entity.Pear, authorizationEntity entity.Authorization) error
	CreateData(pearEntity entity.Pear, authorizationEntity entity.Authorization) (entity.Pear, error)
}

type pearDataInteractor struct {
	pearRepository       repository.Pear
	cacheRepository      repository.Cache
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

func (p pearDataInteractor) UpdateAdminData(pearEntity entity.Pear, authorizationEntity entity.Authorization) error {
	db, _ := utils.ConnectDB()
	dbForClose, _ := db.DB()
	defer dbForClose.Close()

	jwtToken, err := p.cacheRepository.Get(authorizationEntity.JWTKey)
	if err != nil {
		return err
	}
	if jwtToken != authorizationEntity.JWTToken {
		return errors.New("incorrect jwt token")
	}

	if err := p.pearRepository.Update(db, pearEntity); err != nil {
		return err
	}
	return nil
}

func (p pearDataInteractor) CreateData(pearEntity entity.Pear, authorizationEntity entity.Authorization) (entity.Pear, error) {
	db, _ := utils.ConnectDB()
	dbForClose, _ := db.DB()
	defer dbForClose.Close()

	jwtToken, err := p.cacheRepository.Get(authorizationEntity.JWTKey)
	if err != nil {
		return entity.Pear{}, err
	}
	if jwtToken != authorizationEntity.JWTToken {
		return entity.Pear{}, errors.New("incorrect jwt token")
	}

	pear, err := p.pearRepository.Create(db, pearEntity)
	if err != nil {
		return entity.Pear{}, err
	}

	return pear, nil
}

func NewPearData(pearRepository repository.Pear, cacheRepository repository.Cache, pearVersionPresenter presenter.PearVersion) Pear {
	return pearDataInteractor{
		pearRepository:       pearRepository,
		cacheRepository:      cacheRepository,
		pearVersionPresenter: pearVersionPresenter,
	}
}
