package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var (
	db *gorm.DB
)

func Connect() {
	dbInfo := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DB_NAME"))
	d, err := gorm.Open("mysql", dbInfo)
	if err != nil {
		panic(err)
	}

	db = d
	fmt.Println("Connected to MySQL Database!")
}

func GetDB() *gorm.DB {
	return db
}
