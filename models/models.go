package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/ademuanthony/Bas/common"
)

func init() {
	orm.RegisterModel(new(User))

	// set default database
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		common.AppConfig.DbUserName, common.AppConfig.DbPassword, common.AppConfig.Database), 30)
}

type (

	Application struct {
		Id int64 `beedb:"PK"`
		CreateDate time.Time
		CreatedBy time.Time
		UpdatedDate time.Time
		UpdatedBy time.Time
		Name string
		Description string
	}

	User struct {
		Id int64 `beedb:"PK" sql:"id"`
		CreateDate time.Time `sql:"created_date"`
		CreatedBy time.Time `sql:"updated_by"`
		UpdatedDate time.Time `sql:"updated_date"`
		UpdatedBy time.Time `sql:"updated_by"`
		Username string `sql:"username" tname:"auths"`
		Password string `sql:"password"`
		PasswordHash string `sql:"password_hash"`
		Email string `sql:"email"`
		FistName string `sql:"first_name"`
		LastName string `sql:"last_name"`
	}

	Role struct {
		Id int64 `beedb:"PK"`
		CreateDate time.Time
		CreatedBy time.Time
		UpdatedDate time.Time
		UpdatedBy time.Time
		Name string
	}

	UserRole struct {
		Id int64 `beedb:"PK"`
		CreateDate time.Time
		CreatedBy time.Time
		UpdatedDate time.Time
		UpdatedBy time.Time
		Role   Role
		RoleId int
		User   User
		UserId int
	}

	Resource struct {
		Id int64 `beedb:"PK"`
		CreateDate time.Time
		CreatedBy time.Time
		UpdatedDate time.Time
		UpdatedBy time.Time
		Key string
		Application Application
		ApplicationId int
	}




)