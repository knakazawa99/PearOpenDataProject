package testutil

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"api/config"
)

type (
	ErrorResponse struct {
		Message string `json:"message,omitempty"`
	}
)

func Config() *config.Config {
	conf := &config.Config{}
	_ = envconfig.Process("PEAR_", conf)
	return conf
}

func DB() *gorm.DB {
	return openTestConnection(Config())
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	_ = sqlDB.Close()
}

func openTestConnection(c *config.Config) (db *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName) + "&loc=Asia%2FTokyo"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		panic(err)
	}
	return
}

func TruncateTables(db *gorm.DB, gormModels []interface{}) {
	stmt := &gorm.Statement{DB: db}

	for i := range gormModels {
		_ = stmt.Parse(gormModels[i])
		db.Exec("set foreign_key_checks = 0;")
		q := fmt.Sprintf("TRUNCATE TABLE %s", stmt.Schema.Table)
		db.Exec(q)
	}
}
