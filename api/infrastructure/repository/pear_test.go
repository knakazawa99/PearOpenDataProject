package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"api/infrastructure/gormmodel"
	"api/testutil"
)

func TestPear_FindPears(t *testing.T) {
	db := testutil.DB()
	defer testutil.CloseDB(db)
	testutil.TruncateTables(db, []interface{}{
		&gormmodel.GormPear{},
	})

	pearRepository := NewPear()

	testPearData := make([]gormmodel.GormPear, 2)
	testPearData[0] = gormmodel.GormPear{
		Version:  "1.0.0",
		FilePath: "test.zip",
	}
	testPearData[1] = gormmodel.GormPear{
		Version:  "1.0.1",
		FilePath: "test.zip",
	}
	db.Create(&testPearData)
	pears, err := pearRepository.FindPears(db)
	assert.Equal(t, "1.0.0", pears[0].Version)
	assert.Equal(t, 2, len(pears))
	assert.Nil(t, err)
}
