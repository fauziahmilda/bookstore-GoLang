// ini kedua ditulis
// this file for return var call db, help others file interact w db
package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/simplerest?charset=utf8&parseTime=True&loc=Local") //ini belum beres, agatau username sama kata sandi dari mysql
	//charset=utf8&parseTime=True&loc=Local
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
