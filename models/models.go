package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(User))
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
		Role        Role
		RoleId      int
		User        User
		UserId      int
	}

	Resource struct {
		Id            int64
		CreatedDate   time.Time
		CreatedBy     int64
		UpdatedDate   time.Time
		UpdatedBy     int64
		Key           string
		Application   Application
		ApplicationId int
	}
)
