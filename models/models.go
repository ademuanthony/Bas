package models

import "github.com/jinzhu/gorm"

type (

	Application struct {
		gorm.Model
		Name string
		Description string
	}

	Auth struct {
		gorm.Model
		Username string
		Password string
	}

	Role struct {
		gorm.Model
		Name string
	}

	UserRole struct {
		gorm.Model
		Role Role
		RoleId int
		Auth Auth
		AuthID int
	}

	Resource struct {
		gorm.Model
		Key string
		Application Application
		ApplicationID int
	}




)