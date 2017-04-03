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

func CreateApplication(w http.ResponseWriter, r *http.Request) {
	var application models.Application

	err := json.NewDecoder(r.Body).Decode(&application);

	if err != nil{
		common.DisplayAppError(w, err, "Invalid request", http.StatusBadRequest)
		return
	}

	applicationService := services.ApplicationService{Orm:orm.NewOrm()}
	//userInfo := r.Context().Value("UserInfo")


	application.CreatedDate = time.Now()
	application.CreatedBy = 1 //TODO

	application, err = applicationService.CreateApplication(application)

	if err != nil{
		common.SendResult(w, resources.ResponseResource{Message:err.Error(), Success:false}, http.StatusNotAcceptable)
		return
	}
	common.SendResult(w, resources.ResponseResource{Data:application, Message:"Application created", Success:true}, http.StatusCreated)
}

func GetApplications(w http.ResponseWriter, r *http.Request) {
	applicationService := services.ApplicationService{Orm:orm.NewOrm()}
	applications, err := applicationService.GetAll()
	if err != nil{
		common.DisplayAppError(w, err, err.Error(), http.StatusNoContent)
		return
	}
	common.SendResult(w, resources.ResponseResource{Data:applications, Success:true}, http.StatusOK)
}

func GetApplication(w http.ResponseWriter, r *http.Request) {

}


func DeleteApplication(w http.ResponseWriter, r *http.Request) {

}

