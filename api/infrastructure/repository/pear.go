package repository

import (
	"gorm.io/gorm"

	"api/domain/entity"
	"api/domain/repository"
	"api/infrastructure/gormmodel"
)

type pear struct {
}

func (p pear) FindPears(db *gorm.DB) ([]entity.Pear, error) {
	var gormPears []gormmodel.GormPear
	if err := db.Find(&gormPears).Error; err != nil {
		return []entity.Pear{}, err
	}
	pears := make([]entity.Pear, len(gormPears))
	for i := range gormPears {
		pears[i] = entity.Pear{
			ID:             gormPears[i].ID,
			FilePath:       gormPears[i].FilePath,
			Version:        gormPears[i].Version,
			ReleaseNote:    gormPears[i].ReleaseNote,
			ReleaseComment: gormPears[i].ReleaseComment,
			ReleaseFlag:    gormPears[i].ReleaseFlag,
			CreatedAt:      *gormPears[i].CreatedAt,
		}
	}
	return pears, nil
}

func (p pear) FindReleasedPears(db *gorm.DB) ([]entity.Pear, error) {
	var gormPears []gormmodel.GormPear
	if err := db.Where("release_flag = ?", 1).Find(&gormPears).Error; err != nil {
		return []entity.Pear{}, err
	}
	pears := make([]entity.Pear, len(gormPears))
	for i := range gormPears {
		pears[i] = entity.Pear{
			ID:             gormPears[i].ID,
			FilePath:       gormPears[i].FilePath,
			Version:        gormPears[i].Version,
			ReleaseNote:    gormPears[i].ReleaseNote,
			ReleaseComment: gormPears[i].ReleaseComment,
			ReleaseFlag:    gormPears[i].ReleaseFlag,
			CreatedAt:      *gormPears[i].CreatedAt,
		}
	}
	return pears, nil
}

func (p pear) Update(db *gorm.DB, pear entity.Pear) error {
	gormPear := gormmodel.GormPear{
		ID:             pear.ID,
		ReleaseComment: pear.ReleaseComment,
		ReleaseNote:    pear.ReleaseNote,
		ReleaseFlag:    pear.ReleaseFlag,
	}
	if err := db.Debug().Model(&gormPear).Where("id = ?", pear.ID).
		Updates(map[string]interface{}{"release_comment": pear.ReleaseComment, "release_note": pear.ReleaseNote, "release_flag": pear.ReleaseFlag}).
		Error; err != nil {
		return err
	}
	return nil
}

func (p pear) Create(db *gorm.DB, pear entity.Pear) (entity.Pear, error) {
	gormPear := gormmodel.GormPear{
		ReleaseComment: pear.ReleaseComment,
		ReleaseNote:    pear.ReleaseNote,
		ReleaseFlag:    pear.ReleaseFlag,
		Version:        pear.Version,
		FilePath:       pear.FilePath,
	}
	if err := db.Create(&gormPear).Error; err != nil {
		return entity.Pear{}, err
	}
	var createdGormPear gormmodel.GormPear

	if err := db.First(&createdGormPear, gormPear.ID).Error; err != nil {
		return entity.Pear{}, err
	}

	return entity.Pear{
		ID:             createdGormPear.ID,
		ReleaseNote:    createdGormPear.ReleaseNote,
		ReleaseComment: createdGormPear.ReleaseComment,
		Version:        createdGormPear.Version,
		FilePath:       createdGormPear.FilePath,
		CreatedAt:      *createdGormPear.CreatedAt,
	}, nil

}

func NewPear() repository.Pear {
	return &pear{}
}
