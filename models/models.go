package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(Application), new(User), new(Role), new(UserRole), new(Resource), new(RoleResource))
}

type (
	Application struct {
		Id          int64
		CreatedDate time.Time
		CreatedBy   int64
		UpdatedDate time.Time
		UpdatedBy   int64
		Name        string
		Description string
	}

	User struct {
		Id           int64
		CreatedDate  time.Time
		CreatedBy    int64
		UpdatedDate  time.Time
		UpdatedBy    int64
		Username     string
		Password     string
		PasswordHash string
		Email        string
		FirstName    string
		LastName     string
	}

	Role struct {
		Id          int64
		CreatedDate time.Time
		CreatedBy   int64
		UpdatedDate time.Time
		UpdatedBy   int64
		Name        string
	}

	UserRole struct {
		Id          int64
		CreatedDate time.Time
		CreatedBy   int64
		UpdatedDate time.Time
		UpdatedBy   int64
		Role        *Role `orm:"rel(fk)"`
		User        *User `orm:"rel(fk)"`
	}

	Resource struct {
		Id            int64
		CreatedDate   time.Time
		CreatedBy     int64
		UpdatedDate   time.Time
		UpdatedBy     int64
		Key           string
		Application   *Application `orm:"rel(fk)"`
	}

	RoleResource struct {
		Id int64
		RoleId int64
		Resource *Resource `orm:"rel(fk)"`

	}
)