package database

import (
	"fmt"
	"go_social/internal/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type DBConn struct {
	DBName string
	DBHost string
	DBPort string
	DBUser string
	DBPass string
}

func DB() *gorm.DB {
	common.EnvLoad()
	dbVar := &DBConn{
		DBName: os.Getenv("DB_NAME"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbVar.DBHost, dbVar.DBPort, dbVar.DBUser, dbVar.DBPass, dbVar.DBName)
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	common.Err(db.Error, "Can't connect to database")

	return db
}
