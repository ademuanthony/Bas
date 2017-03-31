package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/ademuanthony/Bas/common"
	"github.com/ademuanthony/Bas/data"
	"github.com/ademuanthony/Bas/resources"
)

func AuthRegister(w http.ResponseWriter, r *http.Request)  {
	var userResource resources.UserResource
	err := json.NewDecoder(r.Body).Decode(&userResource)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid user data", http.StatusBadRequest)
		return
	}
	user := userResource.Data
	userRepository := data.UserRepository{}
	id, err := userRepository.CreateUser(user)
	if err != nil{
		common.DisplayAppError(w, err, err.Error(), http.StatusBadRequest)
		return
	}
	user.PasswordHash = ""
	user.Id = id
	common.SendResult(w, resources.UserResource{Data:user}, http.StatusCreated)
}

func AuthLogin(w http.ResponseWriter, r *http.Request){
	var dataResource resources.LoginResource
	var token string
	// Decode the incoming login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid login data", http.StatusBadRequest)
		return
	}
	loginModel := dataResource.Data

	userRepository := data.UserRepository{}
	// Authenticate the login user
	if user, err := userRepository.Login(loginModel.Username, loginModel.Password); err != nil{
		common.DisplayAppError(w, err, "Invalid credentials", http.StatusUnauthorized)
		return
	}else {// if login is successful
		// Generate json web token
		token, err = common.GenerateJWT(user.Username, "member")
		if err != nil{
			common.DisplayAppError(w, err, "Error while generating access token", 500)
			return
		}
		user.PasswordHash = ""
		athUser := resources.AuthUserModel{User:user, Token:token}
		common.SendResult(w, athUser, http.StatusOK)
	}
}
