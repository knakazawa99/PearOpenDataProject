package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"api/domain/entity"
	"api/domain/entity/types"
	"api/infrastructure/gormmodel"
	"api/testutil"
	"api/utils"
)

func TestAuth_SaveAuthUser(t *testing.T) {
	db := testutil.DB()
	defer testutil.CloseDB(db)

	testutil.TruncateTables(db, []interface{}{
		&gormmodel.GormAuthInformation{},
		&gormmodel.GormToken{},
	})

	authRepository := NewAuth()
	authEntity := entity.Auth{
		Email: entity.Email("test@gmail.com"),
		Type:  types.TypeUser,
		Token: "test",
	}
	err := authRepository.SaveAuth(db, authEntity)
	var resultGormAuthInformation gormmodel.GormAuthInformation
	var resultGormToken gormmodel.GormToken
	db.Where("id = ?", 1).Take(&resultGormAuthInformation)
	db.Where("id = ?", 1).Take(&resultGormToken)

	assert.Nil(t, err)
	assert.Equal(t, string(authEntity.Email), resultGormAuthInformation.Email)
	assert.Equal(t, authEntity.Token, resultGormToken.Token)
}

func TestAuth_SaveAuthAdmin(t *testing.T) {
	db := testutil.DB()
	defer testutil.CloseDB(db)

	testutil.TruncateTables(db, []interface{}{
		&gormmodel.GormAuthInformation{},
		&gormmodel.GormAdminInformation{},
	})

	authRepository := NewAuth()
	authEntity := entity.Auth{
		Email:    entity.Email("noir.alpca@gmail.com"),
		Type:     types.TypeAdmin,
		Password: "hogehoge",
	}
	authEntity.Password, _ = utils.PasswordEncrypt(authEntity.Password)
	err := authRepository.SaveAuth(db, authEntity)
	var resultGormAuthInformation gormmodel.GormAuthInformation
	var resultGormAdminInformation gormmodel.GormAdminInformation
	db.Where("id = ?", 1).Take(&resultGormAuthInformation)
	db.Where("auth_information_id = ?", 1).Take(&resultGormAdminInformation)

	assert.Nil(t, err)
	assert.Equal(t, string(authEntity.Email), resultGormAuthInformation.Email)
	assert.Equal(t, authEntity.Password, resultGormAdminInformation.Password)
}

func TestAuth_FindByEmail(t *testing.T) {
	db := testutil.DB()
	defer testutil.CloseDB(db)

	testutil.TruncateTables(db, []interface{}{
		&gormmodel.GormAuthInformation{},
		&gormmodel.GormToken{},
	})

	authRepository := NewAuth()
	gormAuthInformation := gormmodel.GormAuthInformation{
		ID:       1,
		Email:    "test@gmail.com",
		AuthType: "user",
	}
	gormToken := gormmodel.GormToken{
		ID:                1,
		AuthInformationID: 1,
		Token:             "token_test",
	}
	db.Create(&gormAuthInformation)
	db.Create(&gormToken)

	authEntity, _ := authRepository.FindByEmail(db, entity.Email("test@gmail.com"))
	assert.Equal(t, authEntity.Email, entity.Email("test@gmail.com"))
}

func TestAuth_FindByType(t *testing.T) {
	db := testutil.DB()
	defer testutil.CloseDB(db)

	testutil.TruncateTables(db, []interface{}{
		&gormmodel.GormAuthInformation{},
	})

	gormAuthInformation := gormmodel.GormAuthInformation{
		ID:       1,
		Email:    "test@gmail.com",
		AuthType: "admin",
	}
	gormAuthInformation2 := gormmodel.GormAuthInformation{
		ID:       2,
		Email:    "test@gmail.com",
		AuthType: "user",
	}
	db.Create(&gormAuthInformation)
	db.Create(&gormAuthInformation2)
	authRepository := NewAuth()
	authEntities, err := authRepository.FindByType(db, types.TypeAdmin)
	assert.Nil(t, err)
	assert.Equal(t, gormAuthInformation.ID, authEntities[0].ID)
	assert.Equal(t, 1, len(authEntities))
}

func TestAuth_Delete(t *testing.T) {
	db := testutil.DB()
	defer testutil.CloseDB(db)

	testutil.TruncateTables(db, []interface{}{
		&gormmodel.GormAuthInformation{},
		&gormmodel.GormAdminInformation{},
	})

	gormAuthInformation := gormmodel.GormAuthInformation{
		ID:       1,
		Email:    "test@gmail.com",
		AuthType: "admin",
	}
	gormAdminInformation := gormmodel.GormAdminInformation{
		ID:                1,
		AuthInformationID: 1,
		Password:          "hogehoge",
	}
	db.Create(&gormAuthInformation)
	db.Create(&gormAdminInformation)

	authEntity := entity.Auth{
		ID:       1,
		Email:    entity.Email("test@gmail.com"),
		Type:     types.TypeAdmin,
		Password: "hogehoge",
	}

	authRepository := NewAuth()
	err := authRepository.Delete(db, authEntity)
	assert.Nil(t, err)
	var deletedGormAuthInformation gormmodel.GormAuthInformation
	err = db.First(&deletedGormAuthInformation, 1).Error
	assert.Equal(t, gorm.ErrRecordNotFound.Error(), err.Error())
}
