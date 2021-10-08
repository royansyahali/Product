package config

import (
	"erajaya-project/domain"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	dbHost := "localhost"
	dbUsername := "root"
	dbName := "Erajaya"

	mysqlConfig := fmt.Sprintf("%s:@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbHost, dbName)

	db, err := gorm.Open(mysql.Open(mysqlConfig), &gorm.Config{})

	if err != nil {
		return nil, errors.Wrap(err, "failed to open DB connection")
	}
	if !db.Migrator().HasTable(&domain.Product{}) {
		db.Migrator().CreateTable(&domain.Product{})
	}

	return db, nil
}
