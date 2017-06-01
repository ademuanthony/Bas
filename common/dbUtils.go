package common

import (
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func createDatabaseConnection() {
	var err error

	// set default database
	/*err = orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
	AppConfig.DbUserName, AppConfig.DbPassword, AppConfig.Database), 30)*/
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:@/bas?charset=utf8", 30)

	if err != nil {
		panic(err)
		log.Fatalf("[CreateDatabaseConnection]: %s\n", err)
	}
}

func closeDatabaseConnection() {
	//Db.Close()
}
