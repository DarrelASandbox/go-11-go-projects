package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// "mysql", "DBUsername:DBpassword@tcp(127.0.0.1:3306)/book_management_api?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", "REPLACEME:REPLACEME(127.0.0.1:3306)/book_management_api?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
