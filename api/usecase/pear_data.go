package usecase

import (
	"errors"

	"gorm.io/gorm"

	"api/domain/entity"
	"api/domain/presenter"
	"api/domain/repository"
	"api/http/response"
)

type Pear interface {
	GetDataVersions(db *gorm.DB) ([]response.PearDataVersionOutput, error)
	GetAdminDataVersions(db *gorm.DB) ([]response.PearAdminDataVersionOutput, error)
	UpdateAdminData(db *gorm.DB, pearEntity entity.Pear, authorizationEntity entity.Authorization) error
	CreateData(db *gorm.DB, pearEntity entity.Pear, authorizationEntity entity.Authorization) (entity.Pear, error)
}

type pearDataInteractor struct {
	pearRepository       repository.Pear
	cacheRepository      repository.Cache
	pearVersionPresenter presenter.PearVersion
}

func (p pearDataInteractor) GetDataVersions(db *gorm.DB) ([]response.PearDataVersionOutput, error) {
	pears, err := p.pearRepository.FindReleasedPears(db)
	if err != nil {
		return []response.PearDataVersionOutput{}, err
	}
	return p.pearVersionPresenter.OutPutPearVersions(pears), nil
}

func (p pearDataInteractor) GetAdminDataVersions(db *gorm.DB) ([]response.PearAdminDataVersionOutput, error) {
	pears, err := p.pearRepository.FindPears(db)
	if err != nil {
		return []response.PearAdminDataVersionOutput{}, err
	}
	return p.pearVersionPresenter.OutPutPearAdminVersions(pears), nil
}

func (p pearDataInteractor) UpdateAdminData(db *gorm.DB, pearEntity entity.Pear, authorizationEntity entity.Authorization) error {
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

func (p pearDataInteractor) CreateData(db *gorm.DB, pearEntity entity.Pear, authorizationEntity entity.Authorization) (entity.Pear, error) {
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
