package common

import (
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/astaxie/beedb"
)

var Db *sql.DB
var Orm beedb.Model

func createDatabaseConnection() {
	var err error
	Db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		AppConfig.DbUserName, AppConfig.DbPassword, AppConfig.Database))


	if err != nil{
		log.Fatalf("[CreateDatabaseConnection]: %s\n", err)
	}

	beedb.OnDebug = true

	Orm = beedb.New(Db)
}

func closeDatabaseConnection() {
	Db.Close()
}
