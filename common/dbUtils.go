package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"fmt"
)

var Db *gorm.DB

func createDatabaseConnection() {
	var err error
	//Db, err = gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	Db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		AppConfig.DbUserName, AppConfig.DbPassword, AppConfig.Database))
	/*Db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
		AppConfig.DbUserName, AppConfig.DbPassword, AppConfig.DbHost, AppConfig.Database))*/
	if err != nil{
		log.Fatalf("[CreateDatabaseConnection]: %s\n", err)
	}
}

func closeDatabaseConnection() {
	Db.Close()
}
