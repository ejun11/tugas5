package connectiondb

import (
	"fmt"
	"tugas5/project/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectGormDB() (*gorm.DB, error) {
	var (
		DB_NAME = "toko-online"
		DB_USER = "root"
		DB_PASS = ""
		DB_HOST = "localhost"
		DB_PORT = "3306"
	)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	dbs, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = dbs.Ping()
	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(
		model.User{},
		model.Category{},
		model.Product{},
		model.Cart{},
		model.Order{},
		model.OrderItem{},
		model.Payment{},
		model.Delivery{},
	)

	return db, nil
}
