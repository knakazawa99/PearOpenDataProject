package middleware

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	DB interface {
		Connection(next echo.HandlerFunc) echo.HandlerFunc
	}

	dbImpl struct {
	}
)

func NewDB() DB {
	return &dbImpl{}
}

func (d dbImpl) Connection(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := gormConnection()
		fmt.Println("db in middleware:: ", db)
		if err != nil {

		}
		dbForClose, err := db.DB()

		if err != nil {

		}
		defer dbForClose.Close()

		c.Set("DB", db)

		return next(c)
	}
}

func gormConnection() (*gorm.DB, error) {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("failed: load env value : %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
