package controllers

import (
	"net/http"
	"github.com/ademuanthony/Bas/resources"
	"encoding/json"
	"github.com/ademuanthony/Bas/common"
	"github.com/ademuanthony/Bas/services"
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
	var contentType string
	if emailResource.IsHtml{
		contentType = "text/html"
	}else{
		contentType = "text/plain"
	}
	emailService := services.EmailService{}
	err = emailService.SendEmail(emailResource.From, emailResource.To,
		emailResource.Title, emailResource.Message, contentType)
	if err != nil{
		common.DisplayAppError(w, err, "Email not sent", http.StatusBadRequest)
		return
	}

	common.SendResult(w, resources.ResponseResource{Data:true, Message:"Email Sent", Success:true}, http.StatusOK)
}
