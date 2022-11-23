package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //from the docs of gorm

)

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql","root:pass/simplerest?charset=utf8&parseTime=True&Loc=Local" )
	if err != nil{
		panic(err)
	}
	db =d
}

func GetDB() *gorm.DB{
	return db
}