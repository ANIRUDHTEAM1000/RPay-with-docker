package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DNS = "root:Team@1000@tcp(db:3306)/RPay"

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
