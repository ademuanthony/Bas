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
	"github.com/gorilla/mux"
	"strconv"
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

// Handler /applications [POST]
// Returns a list of models.Applications
func GetApplications(w http.ResponseWriter, r *http.Request) {
	applicationService := services.ApplicationService{Orm:orm.NewOrm()}
	applications, err := applicationService.GetAll()
	if err != nil{
		common.DisplayAppError(w, err, err.Error(), http.StatusNoContent)
		return
	}
	common.SendResult(w, resources.ResponseResource{Data:applications, Success:true}, http.StatusOK)
}

// Handler /applications/{id} [GET]
// Returns a single models.Application
func GetApplication(w http.ResponseWriter, r *http.Request) {
	//get id from in ocming request
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid Application Id", http.StatusBadRequest)
	}
	applicationService := services.ApplicationService{Orm:orm.NewOrm()}
	application,err := applicationService.GetById(id)
	if err != nil{
		common.DisplayAppError(w, err, "Application not found", http.StatusNoContent)
	}
	common.SendResult(w, resources.ResponseResource{Data:application, Success:true}, http.StatusOK)
}

// Handler /applications/{id} [DELETE]
// Deletes the application with the specified id
func DeleteApplication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	service := services.ApplicationService{Orm:orm.NewOrm()}

	message, err := service.DeleteApplication(id)

	if err != nil{
		common.DisplayAppError(w, err, err.Error(), http.StatusBadRequest)
	}
	common.SendResult(w, resources.ResponseResource{Message:message, Success:true}, http.StatusOK)
}

