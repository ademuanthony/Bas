package controllers

import (
	"encoding/json"
	"github.com/ademuanthony/bas/common"
	"github.com/ademuanthony/bas/resources"
	"github.com/astaxie/beego/orm"
	"net/http"
	"github.com/ademuanthony/bas/services"
	"errors"
	"github.com/ademuanthony/bas/models"
	"github.com/gorilla/mux"
	"strconv"
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
	o := orm.NewOrm()
	userService := services.UserService{Orm: o}
	id, err := userService.CreateUser(user)
	if err != nil {
		common.DisplayAppError(w, err, err.Error(), http.StatusBadRequest)
		return
	}
	user.PasswordHash = ""
	user.Id = id

	aclService := services.AclService{Orm:o}

	aclService.AddUserToRole(id, 4)

	common.SendResult(w, resources.ResponseResource{Data: user, Success:true}, http.StatusCreated)
}

// AuthUpdate changes the details of a user specified by the id
// Handler /auth/{id}/update [POST]
func AuthUpdate(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value("UserInfo")
	common.Log("CurrentUser", currentUser)
	var userResource resources.UserResource
	err := json.NewDecoder(r.Body).Decode(&userResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid userModel data", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid User Id", http.StatusBadRequest)
	}

	userModel := userResource.Data
	userService := services.UserService{Orm: orm.NewOrm()}

	user, err := userService.GetUserById(id)

	if err != nil{
		common.DisplayAppError(w, err, "Invalid user Id", http.StatusBadRequest)
		return
	}

	if userModel.Email != ""{
		user.Email = userModel.Email
	}

	if userModel.Username != ""{
		user.Username = userModel.Username
	}

	if userModel.FirstName != ""{
		user.FirstName = userModel.FirstName
	}

	if userModel.LastName != ""{
		user.LastName = userModel.LastName
	}

	err = userService.UpdateUser(user)
	if err != nil {
		common.DisplayAppError(w, err, err.Error(), http.StatusBadRequest)
		return
	}
	userModel.PasswordHash = ""
	userModel.Id = id
	common.SendResult(w, resources.ResponseResource{Data: userModel, Success: true}, http.StatusCreated)
}


func AuthCreateAccounts(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value("UserInfo")
	common.Log("CurrentUser", currentUser)
	var userResources []resources.UserResource
	err := json.NewDecoder(r.Body).Decode(&userResources)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid user data", http.StatusBadRequest)
		return
	}

	users := make([]models.User, len(userResources))
	o := orm.NewOrm()
	//o.Begin()
	for index, userResource := range userResources{
		user := userResource.Data
		userService := services.UserService{Orm: o}
		id, _ := userService.CreateUser(user)
		/*if err != nil {
			common.DisplayAppError(w, err, err.Error(), http.StatusBadRequest)
			o.Rollback()
			return
		}*/
		user.PasswordHash = ""
		user.Id = id
		users[index] = user
	}

	//o.Commit()
	common.SendResult(w, resources.ResponseResource{Data: users, Success:true}, http.StatusCreated)

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

func ChangePassword(w http.ResponseWriter, r * http.Request) {
	var model resources.ChangePasswordModel
	// Decode the incoming json
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid request data", http.StatusBadRequest)
		return
	}

	userService := services.UserService{Orm: orm.NewOrm()}
	// check to see that the newpassword matches its confirmation
	if model.NewPassword != model.ConfirmPassword{
		common.DisplayAppError(w, errors.New("Wrong password confirmation"), "Wrong password confirmation", http.StatusBadRequest)
		return
	}
	currentUser := r.Context().Value("UserInfo").(map[string]interface{})
	err = userService.ChangePassword(int64(currentUser["UserId"].(float64)), model.OldPassword, model.NewPassword)
	if err != nil{
		common.DisplayAppError(w, err, err.Error(), http.StatusBadRequest)
		return
	}
	common.SendResult(w, resources.ResponseResource{Data:true, Success:true}, http.StatusOK)
}

// ChangePasswordForId changes the password of a user specified by the id
// Handler /auth/{id}/changepassword [POST]
func ChangePasswordForId(w http.ResponseWriter, r *http.Request)  {
	//get id from in coming request
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid User Id", http.StatusBadRequest)
	}

	var model resources.ChangePasswordModel
	// Decode the incoming json
	err = json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid request data", http.StatusBadRequest)
		return
	}

	userService := services.UserService{Orm:orm.NewOrm()}

	err = userService.ChangePasswordForId(id, model.NewPassword)
	if err != nil{
		common.DisplayAppError(w, err, err.Error(), http.StatusBadRequest)
		return
	}
	common.SendResult(w, resources.ResponseResource{Data:true, Success:true}, http.StatusOK)

}