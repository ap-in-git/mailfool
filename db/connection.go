package db

import (
	"database/sql"
	"fmt"
	"github.com/ap-in-git/mailfool/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func Init(config *config.DbConfig) (*sql.DB, *gorm.DB) {
	dbHost := config.Host
	dbDatabase := config.Database
	dbUsername := config.User
	dbPassword := config.Password
	dbPort := config.Port

	connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True", dbUsername, dbPassword, dbHost, dbPort, dbDatabase)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}

	return sqlDB, db
}
