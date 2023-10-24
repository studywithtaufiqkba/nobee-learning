package database

import (
	"day_6/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectGORMMySQL(cfg config.DB) (db *gorm.DB, err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/day_6?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	err = db.AutoMigrate()
	if err != nil {
		return nil, err
	}
	return db, nil
}
