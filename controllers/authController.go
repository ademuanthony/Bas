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
	var dataResource resources.LoginResource
	var token string
	// Decode the incoming login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid login data", http.StatusBadRequest)
		return
	}
	loginModel := dataResource.Data

	userService := services.UserService{Orm: orm.NewOrm()}
	// Authenticate the login user
	if user, err := userService.Login(loginModel.Username, loginModel.Password); err != nil {
		common.DisplayAppError(w, err, "Invalid credentials", http.StatusUnauthorized)
		return
	} else { // if login is successful
		// Generate json web token
		tokenData := common.TokenData{UserId:user.Id, Permissions:[]int64{1,2,3,4,5,6,7,8,9,0,12,32,43,545,665,75,777,1,2,3,
			4,5,6,7,8,9,0,12,32,43,545,665,75,777,1,2,3,4,5,6,7,8,9,0,12,32,43,545,665,75,777,23}} //todo
		token, err = common.GenerateJWT(tokenData)
		if err != nil {
			common.DisplayAppError(w, err, "Error while generating access token", 500)
			return
		}
		user.PasswordHash = ""
		athUser := resources.AuthUserModel{User: user, Token: token}
		common.SendResult(w, resources.ResponseResource{Data:athUser, Success:true}, http.StatusOK)
	}
}
