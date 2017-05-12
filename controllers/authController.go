package controllers

import (
	"encoding/json"
	"github.com/ademuanthony/Bas/common"
	"github.com/ademuanthony/Bas/resources"
	"github.com/astaxie/beego/orm"
	"net/http"
	"github.com/ademuanthony/Bas/services"
)

func AuthRegister(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value("UserInfo")
	common.Log("CurrentUser", currentUser)
	var userResource resources.UserResource
	err := json.NewDecoder(r.Body).Decode(&userResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid user data", http.StatusBadRequest)
		return
	}
	user := userResource.Data
	userService := services.UserService{Orm: orm.NewOrm()}
	id, err := userService.CreateUser(user)
	if err != nil {
		common.DisplayAppError(w, err, err.Error(), http.StatusBadRequest)
		return
	}
	user.PasswordHash = ""
	user.Id = id
	common.SendResult(w, resources.ResponseResource{Data: user, Success:true}, http.StatusCreated)

}

func AuthLogin(w http.ResponseWriter, r *http.Request) {
	var loginModel resources.LoginModel
	var token string
	// Decode the incoming login json
	err := json.NewDecoder(r.Body).Decode(&loginModel)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid login data", http.StatusBadRequest)
		return
	}

	userService := services.UserService{Orm: orm.NewOrm()}
	// Authenticate the login user
	if user, err := userService.Login(loginModel.Username, loginModel.Password); err != nil {
		common.DisplayAppError(w, err, "Invalid credentials", http.StatusUnauthorized)
		return
	} else {
		// if login is successful
		service := services.AclService{Orm:orm.NewOrm()}

		resourceIds := service.GetResourcesForUser(user.Id)
		// Generate json web token
		tokenData := common.TokenData{UserId:user.Id, Permissions:resourceIds}
		token, err = common.GenerateJWT(tokenData)
		if err != nil {
			common.DisplayAppError(w, err, "Error while generating access token", 500)
			return
		}
		user.PasswordHash = ""
		/*
		athUser := resources.AuthUserModel{User: user, Token: token}
		*/
		common.SendResult(w, resources.ResponseResource{Data:token, Success:true}, http.StatusOK)
	}
}

func ChangePassword() {

}