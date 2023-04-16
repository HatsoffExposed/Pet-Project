package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	_ "github.com/go-sql-driver/mysql"
	// "database/sql"
)

var (
	db *gorm.DB
)

func Connect() *gorm.DB {
	d, err := gorm.Open("mysql", "Hatsoff:HISgra@2@tcp(127.0.0.1:3306)/hello?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db = d
	return db
}

// func GetDB() *gorm.DB {
// 	return db
// }
