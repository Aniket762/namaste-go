// Return a DB which will help other objects to interact with the DB

package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	db  *gorm.DB
)

func Connect(){
	// first param for type of DB and following is admin:pwd/tableName, post ? is the required param for running mysql
	d,err := gorm.Open("mysql","admin:12345/simplerest?charset=utf8&parseTime=True&loc=Local") 
	if err != nil{
		panic(err)
	}
	db = d 
}

func GetDB() *gorm.DB{
	return db 
}