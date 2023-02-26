package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"api/domain/entity"
	"api/infrastructure/gormmodel"
	"api/testutil"
)

func TestDownloadPear_Find(t *testing.T) {
	db := testutil.DB()
	defer testutil.CloseDB(db)

	testutil.TruncateTables(db, []interface{}{
		&gormmodel.GormPear{},
	})

	gormPear := gormmodel.GormPear{
		FilePath: "test.zip",
		Version:  "1.0.0",
	}
	db.Create(&gormPear)

	requestDownloadPearEntity := entity.DownloadPear{
		AuthInfo: entity.Auth{},
		Version:  "1.0.0",
	}
	downloadPearRepository := NewDownloadPear()
	resultDownloadPear, err := downloadPearRepository.Find(db, requestDownloadPearEntity)
	assert.Nil(t, err)
	assert.Equal(t, "1.0.0", resultDownloadPear.Version)
}
