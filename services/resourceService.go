package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/Bas/models"
	"errors"
)

type ResourceService struct {
	Orm orm.Ormer
}

func (this *ResourceService) Create(resource models.Resource) (models.Resource, error) {
	if this.Orm.QueryTable("resource").Filter("key", resource.Key).Exist() {
		return resource, errors.New("Resource with the same key already exists")
	}

	_, err := this.Orm.Insert(&resource)

	return resource, err
}

func (this *ResourceService) GetById(id int64) (models.Resource, error) {
	var resource models.Resource
	err := this.Orm.QueryTable("resource").Filter("id", id).One(&resource)

	return resource, err
}

func (this *ResourceService) GetByName(name struct{}) (models.Resource, error) {
	var resource models.Resource
	err := this.Orm.QueryTable("resource").Filter("name", name).One(&resource)

	return resource, err
}

func (this *ResourceService) GetAll() ([]models.Resource, error) {
	var resources []models.Resource
	_, err := this.Orm.QueryTable("resource").RelatedSel().All(&resources)

	return resources, err
}