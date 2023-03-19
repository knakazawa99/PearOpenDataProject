package repository

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"api/domain/entity"
	"api/infrastructure/gormmodel"
	"api/testutil"
)

func TestPear_FindReleasedPears(t *testing.T) {
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
	pears, err := pearRepository.FindReleasedPears(db)
	assert.Equal(t, "1.0.0", pears[0].Version)
	assert.Equal(t, 2, len(pears))
	assert.Nil(t, err)
}

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

func TestPear_Update(t *testing.T) {
	db := testutil.DB()
	defer testutil.CloseDB(db)
	testutil.TruncateTables(db, []interface{}{
		&gormmodel.GormPear{},
	})

	pearRepository := NewPear()

	testPearData := gormmodel.GormPear{
		Version:        "1.0.0",
		FilePath:       "test.zip",
		ReleaseComment: "release_comment",
		ReleaseNote:    "release_note",
		ReleaseFlag:    true,
	}
	db.Create(&testPearData)
	var beforePearData gormmodel.GormPear
	db.Where("version = ?", "1.0.0").Take(&beforePearData)
	fmt.Println(beforePearData)

	updatePearData := entity.Pear{
		ID:             1,
		Version:        "1.0.0",
		FilePath:       "test.zip",
		ReleaseComment: "updated_comment",
		ReleaseNote:    "updated_note",
		ReleaseFlag:    false,
	}
	err := pearRepository.Update(db, updatePearData)

	var afterPearData gormmodel.GormPear
	db.Where("version = ?", "1.0.0").Take(&afterPearData)
	assert.Nil(t, err)
	assert.Equal(t, updatePearData.ReleaseComment, afterPearData.ReleaseComment)
	assert.Equal(t, updatePearData.ReleaseNote, afterPearData.ReleaseNote)
	assert.Equal(t, updatePearData.ReleaseFlag, afterPearData.ReleaseFlag)
}
