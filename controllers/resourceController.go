package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/ademuanthony/Bas/common"
	"github.com/ademuanthony/Bas/services"
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/Bas/models"
	"time"
	"github.com/ademuanthony/Bas/resources"
)

func CreateResource(w http.ResponseWriter, r *http.Request) {

	currentUser := r.Context().Value("UserInfo").(map[string]interface{})

	var resourceVm resources.CreateResourceResource

	err := json.NewDecoder(r.Body).Decode(&resourceVm);

	if err != nil{
		common.DisplayAppError(w, err, "Invalid request", http.StatusBadRequest)
		return
	}

	resourceService := services.ResourceService{Orm:orm.NewOrm()}
	application := new(models.Application)
	application.Id = resourceVm.ApplicationId

	resource := models.Resource{
		CreatedBy:int64(currentUser["UserId"].(float64)),
		UpdatedBy:int64(currentUser["UserId"].(float64)),
		CreatedDate:time.Now(),
		UpdatedDate:time.Now(),
		Key:resourceVm.Key,
		Application:application,
	}

	resource.CreatedDate = time.Now()
	resource.CreatedBy = int64(currentUser["UserId"].(float64))
	resource.UpdatedDate = time.Now()
	resource.UpdatedBy = int64(currentUser["UserId"].(float64))

	resource, err = resourceService.Create(resource)

	if err != nil{
		common.SendResult(w, resources.ResponseResource{Message:err.Error(), Success:false}, http.StatusNotAcceptable)
		return
	}
	common.SendResult(w, resources.ResponseResource{Data:resource, Message:"Resource created", Success:true}, http.StatusCreated)
}

func GetResources(w http.ResponseWriter, r *http.Request) {
	resourceService := services.ResourceService{Orm:orm.NewOrm()}

	_resources, _ := resourceService.GetAll()
	common.SendResult(w, resources.ResponseResource{Data:_resources, Success:true}, http.StatusOK)
}

func GetResource(w http.ResponseWriter, r *http.Request) {

}


func DeleteResource(w http.ResponseWriter, r *http.Request) {

}

