package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/Bas/models"
	"errors"
)

type ApplicationService struct {
	Orm orm.Ormer
}

func (this *ApplicationService) CreateApplication(application models.Application) (models.Application, error) {
	if this.Orm.QueryTable("application").Filter("name", application.Name).Exist() {
		return application, errors.New("Application with the same name already exists")
	}

	_, err := this.Orm.Insert(&application)

	return application, err
}

func (this ApplicationService) GetById(id int64) (models.Application, error) {
	var application models.Application
	err := this.Orm.QueryTable("application").Filter("id", id).One(&application)

	return application, err
}

func (this ApplicationService) GetByName(name struct{}) (models.Application, error) {
	var application models.Application
	err := this.Orm.QueryTable("application").Filter("name", name).One(&application)

	return application, err
}

func (this *ApplicationService) GetAll() ([]models.Application, error) {
	var applications []models.Application
	_, err := this.Orm.QueryTable("application").All(&applications)

	return applications, err
}