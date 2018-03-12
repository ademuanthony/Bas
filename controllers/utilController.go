package controllers

import (
	"net/http"
	"github.com/ademuanthony/bas/resources"
	"encoding/json"
	"github.com/ademuanthony/bas/common"
	"github.com/ademuanthony/bas/services"
)

// Handler /utils/sendemail [POST]
// SendEmail sends an email for the specified params
func SendEmail(w http.ResponseWriter, r *http.Request) {
	var emailResource resources.EmailResource
	err := json.NewDecoder(r.Body).Decode(&emailResource)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid email data", http.StatusBadRequest)
		return
	}
	emailService := services.EmailService{}
	err = emailService.SendEmail(emailResource.From, emailResource.To,
		emailResource.Title, emailResource.Message, emailResource.ContentType)
	if err != nil{
		common.DisplayAppError(w, err, "Email not sent", http.StatusBadRequest)
		return
	}

	common.SendResult(w, resources.ResponseResource{Data:true, Message:"Email Sent", Success:true}, http.StatusOK)
}

// Handler /utils/sendemail [POST]
// SendMultipleEmail sends an email for the specified params
func SendMultipleEmail(w http.ResponseWriter, r *http.Request) {
	var emailResources []resources.EmailResource
	err := json.NewDecoder(r.Body).Decode(&emailResources)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid email data", http.StatusBadRequest)
		return
	}
	emailService := services.EmailService{}
	go emailService.SendMultipleEmails(emailResources)

	common.SendResult(w, resources.ResponseResource{Data:true, Message:"Email queued on the server", Success:true}, http.StatusOK)
}
