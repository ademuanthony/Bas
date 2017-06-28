package controllers

import (
	"bitbucket.org/superfluxteam/pmsserver/common"
	"bitbucket.org/superfluxteam/pmsserver/resources"
	"net/http"
)

func ChangePassword(w http.ResponseWriter, r * http.Request) {
	common.SendResult(w, resources.ResponseResource{Data:true, Success:true}, http.StatusOK)
}