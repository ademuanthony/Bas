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

// Handler /resources [POST]
// Creates a new role in the system
func CreateResource(w http.ResponseWriter, r *http.Request) {

	currentUser := r.Context().Value("UserInfo").(map[string]interface{})

	var resourceVm resources.CreateResourceResource

	err := json.NewDecoder(r.Body).Decode(&resourceVm);

	if err != nil{
		common.DisplayAppError(w, err, "Invalid request", http.StatusBadRequest)
		return
	}

	resourceService := services.AclService{Orm:orm.NewOrm()}
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

	resource, err = resourceService.CreateResource(resource)

	if err != nil{
		common.SendResult(w, resources.ResponseResource{Message:err.Error(), Success:false}, http.StatusNotAcceptable)
		return
	}
	common.SendResult(w, resources.ResponseResource{Data:resource, Message:"Resource created", Success:true}, http.StatusCreated)
}


// Handler /resources/createmany [POST]
// Creates a new role in the system
func CreateResources(w http.ResponseWriter, r *http.Request) {

	currentUser := r.Context().Value("UserInfo").(map[string]interface{})

	var resourcesVm resources.CreateResourcesInputDto

	err := json.NewDecoder(r.Body).Decode(&resourcesVm);

	if err != nil{
		common.DisplayAppError(w, err, "Invalid request", http.StatusBadRequest)
		return
	}

	resourceService := services.AclService{Orm:orm.NewOrm()}
	application := new(models.Application)
	application.Id = resourcesVm.ApplicationId

	var outputs = make([]resources.CreateResourcesOutputDto, len(resourcesVm.Keys))

	for index, key := range resourcesVm.Keys{
		resource := models.Resource{
			CreatedBy:int64(currentUser["UserId"].(float64)),
			UpdatedBy:int64(currentUser["UserId"].(float64)),
			CreatedDate:time.Now(),
			UpdatedDate:time.Now(),
			Key:key,
			Application:application,
		}

		resource.CreatedDate = time.Now()
		resource.CreatedBy = int64(currentUser["UserId"].(float64))
		resource.UpdatedDate = time.Now()
		resource.UpdatedBy = int64(currentUser["UserId"].(float64))

		resource, err = resourceService.CreateResource(resource)

		output := resources.CreateResourcesOutputDto{Key:key}

		if err != nil{
			output.Message = err.Error()
			output.Success = false
		}else{
			output.Success = true
		}
		outputs[index] = output
	}

	common.SendResult(w, resources.ResponseResource{Data:outputs, Message:"Resources created", Success:true}, http.StatusCreated)
}


// Handler /resources/{id}/update [POST]
// Creates a new role in the system
func UpdateResource(w http.ResponseWriter, r *http.Request) {
	// get id from incoming url
	vars := mux.Vars(r)
	idParam := vars["id"]

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil{
		common.DisplayAppError(w, err, "Please send a valid Id", http.StatusBadRequest)
		return
	}

	currentUser := r.Context().Value("UserInfo").(map[string]interface{})

	var resourceVm resources.CreateResourceResource

	err = json.NewDecoder(r.Body).Decode(&resourceVm);

	if err != nil{
		common.DisplayAppError(w, err, "Invalid request", http.StatusBadRequest)
		return
	}

	o := orm.NewOrm()
	resourceService := services.AclService{Orm:o}

	resource, err := resourceService.GetResourceById(id)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid resource ID", http.StatusBadRequest)
		return
	}

	application := new(models.Application)
	application.Id = resourceVm.ApplicationId

	resource.Key = resourceVm.Key
	resource.UpdatedDate = time.Now()
	resource.UpdatedBy = int64(currentUser["UserId"].(float64))

	_, err = o.Update(&resource)

	if err != nil{
		common.SendResult(w, resources.ResponseResource{Message:err.Error(), Success:false}, http.StatusNotAcceptable)
		return
	}
	common.SendResult(w, resources.ResponseResource{Data:resource, Message:"Resource created", Success:true}, http.StatusCreated)
}


// Handler /resources [GET]
// Returns a list of Resources
func GetResources(w http.ResponseWriter, r *http.Request) {
	resourceService := services.AclService{Orm:orm.NewOrm()}

	_resources, _ := resourceService.GetAllResources()
	common.SendResult(w, resources.ResponseResource{Data:_resources, Success:true}, http.StatusOK)
}

// Handler /resources/{id} [GET]
// Returns a single Resource by id
func GetResourceById(w http.ResponseWriter, r *http.Request) {
	// get id from incoming url
	vars := mux.Vars(r)
	idParam := vars["id"]

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil{
		common.DisplayAppError(w, err, "Please send a valid Id", http.StatusBadRequest)
		return
	}


	aclService := services.AclService{Orm:orm.NewOrm()}

	resource, err := aclService.GetResourceById(id)
	if err != nil{
		common.DisplayAppError(w, err, "Resource not found", http.StatusNotFound)
		return
	}
	common.SendResult(w, resources.ResponseResource{Data:resource,Success:true}, http.StatusOK)
}

// Handler /resources/{id}/delete [POST]
// Deletes the resource with the specified id
func DeleteResource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil{
		common.DisplayAppError(w, err, "Please send a valid ID", http.StatusBadRequest)
	}
	service := services.AclService{Orm:orm.NewOrm()}
	message, err := service.DeleteResource(id)

	if err != nil{
		common.DisplayAppError(w, err, message, http.StatusNotFound)
		return
	}

	common.SendResult(w, resources.ResponseResource{Message:message, Success:true}, http.StatusOK)
}


// Handler /roles 	[POST]
// creates a new role in the system
func CreateRole(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value("UserInfo").(map[string]interface{})

	var role models.Role

	err := json.NewDecoder(r.Body).Decode(&role);
	role.CreatedBy = int64(currentUser["UserId"].(float64))
	role.CreatedDate = time.Now()
	role.UpdatedDate = time.Now()
	role.UpdatedBy = role.CreatedBy

	if err != nil{
		common.DisplayAppError(w, err, "Invalid request", http.StatusBadRequest)
		return
	}

	resourceService := services.AclService{Orm:orm.NewOrm()}

	role, err = resourceService.CreateRole(role)

	if err != nil{
		common.DisplayAppError(w, err, err.Error(), http.StatusBadRequest)
		return
	}

	common.SendResult(w, resources.ResponseResource{Message:"Role created", Success:true, Data:role}, http.StatusOK)
}

// Handler /roles 	[GET]
// Returns a list of models.Role
func GetRoles(w http.ResponseWriter, r *http.Request) {
	service := services.AclService{Orm:orm.NewOrm()}
	roles := service.GetRoles()
	common.SendResult(w, resources.ResponseResource{Data:roles, Success:true}, http.StatusOK)
}

// Handler /roles/{id} 	[GET]
// Returns a single models.Role
func GetRole(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid ID", http.StatusBadRequest)
		return
	}

	service := services.AclService{Orm:orm.NewOrm()}
	role, err := service.GetRoleById(id)

	if err != nil{
		common.DisplayAppError(w, err, "Role not found", http.StatusNotFound)
		return
	}

	common.SendResult(w, resources.ResponseResource{Data:role, Success:true}, http.StatusOK)
}

// Handler /roles/{id} 	[DELETE]
// Deletes the specified role from the system
func DeleteRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil{
		common.DisplayAppError(w, err, "Invalid Id", http.StatusBadRequest)
		return
	}

	service := services.AclService{Orm:orm.NewOrm()}
	err = service.DeleteRole(id)

	if err != nil{
		common.DisplayAppError(w, err, "Error in deleting role", http.StatusInternalServerError)
		return
	}

	common.SendResult(w, resources.ResponseResource{Message:"Role deleted", Success:true}, http.StatusOK)
}



// Handler roles/{roleId}/resources 	[GET]
// Return all resources in the specified role
func GetResourceInRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["roleId"], 10, 64)

	if err != nil{
		common.DisplayAppError(w, err, "Invalid Id", http.StatusBadRequest)
	}
	service := services.AclService{Orm:orm.NewOrm()}
	_resources := service.GetResourcesInRole(id);
	common.SendResult(w, resources.ResponseResource{Data:_resources, Success:true}, http.StatusOK)

}

// Handler roles/{roleId}/resources/add/{resourceId} 	[POST]
// Add a resource with the specified id to the specified role
func AddResourceToRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleId, err := strconv.ParseInt(vars["roleId"], 10, 64)

	if err != nil{
		common.DisplayAppError(w, err, "Invalid role ID", http.StatusBadRequest)
		return
	}

	resourceId, err := strconv.ParseInt(vars["resourceId"], 10, 64)

	if err != nil{
		common.DisplayAppError(w, err, "Invalid resource ID", http.StatusBadRequest)
		return
	}

	service := services.AclService{Orm:orm.NewOrm()}

	id, err := service.AddResourceToRole(resourceId, roleId)

	if err != nil{
		common.DisplayAppError(w, err, err.Error(), http.StatusInternalServerError)
		return
	}
	common.SendResult(w, resources.ResponseResource{Message:"Resource added to role", Success:true, Data:id}, http.StatusOK)
}

// Handler roles/roleId/resources/add	[POST]
// Add a list of resources to the specified role
func AddResourcesToRole(w http.ResponseWriter, r *http.Request) {
	//currentUser := r.Context().Value("UserInfo").(map[string]interface{})

	var dto resources.AddResourcesToRoleInputDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil{
		common.DisplayAppError(w, err, err.Error(), http.StatusInternalServerError)
	}
	service := services.AclService{Orm:orm.NewOrm()}
	err = service.AddResourcesToRole(dto.ResourceIds, dto.RoleId)
	if err != nil{
		common.DisplayAppError(w, err, err.Error(), http.StatusInternalServerError)
	}
	common.SendResult(w, resources.ResponseResource{Message:"Resources Added", StatusCode:http.StatusCreated, Success:true}, http.StatusCreated)
}

// Handler /roles/{roleId}/resources/remove/{resourceId}	[GET]
// Removes the specified resourceId from the specified roleId
func RemoveResourceFromRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleId, err := strconv.ParseInt(vars["roleId"], 10, 64)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid role ID", http.StatusBadRequest)
	}

	resourceId, err := strconv.ParseInt(vars["resourceId"], 10, 64)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid Resource ID", http.StatusBadRequest)
		return
	}

	service := services.AclService{Orm:orm.NewOrm()}
	err = service.RemoveResourceFromRole(resourceId, roleId)

	if err != nil{
		common.DisplayAppError(w, err, "Error in removing resource from role", http.StatusInternalServerError)
		return
	}

	common.SendResult(w, resources.ResponseResource{Message:"Resource removed from role", Success:true}, http.StatusOK)
}


// Handler roles/{roleId}/users		[POST]
// Returns a []models.User
func GetUsersInRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleId, err := strconv.ParseInt(vars["roleId"], 10, 64)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid Role ID", http.StatusBadRequest)
		return
	}
	service := services.AclService{Orm:orm.NewOrm()}
	users := service.GetUsersInRole(roleId)
	common.SendResult(w, resources.ResponseResource{Data:users, Success:true}, http.StatusOK)
}


// Handler /roles/{roleId}/users/{userId} 	[POST]
// Add a resource with the specified id to the specified role
func AddUserToRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleId, err := strconv.ParseInt(vars["roleId"], 10, 64)

	if err != nil{
		common.DisplayAppError(w, err, "Invalid role ID", http.StatusBadRequest)
		return
	}

	userId, err := strconv.ParseInt(vars["userId"], 10, 64)

	if err != nil{
		common.DisplayAppError(w, err, "Invalid user ID", http.StatusBadRequest)
		return
	}

	service := services.AclService{Orm:orm.NewOrm()}

	err = service.AddUserToRole(userId, roleId)

	if err != nil{
		common.DisplayAppError(w, err, err.Error(), http.StatusInternalServerError)
		return
	}
	common.SendResult(w, resources.ResponseResource{Message:"User added to role", Success:true}, http.StatusOK)
}


// Handler /roles/{roleId}/users/{userId} 	[DELETE]
// Removes the specified resourceId from the specified roleId
func RemoveUserFromRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleId, err := strconv.ParseInt(vars["roleId"], 10, 64)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid role ID", http.StatusBadRequest)
	}

	userId, err := strconv.ParseInt(vars["userId"], 10, 64)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid User ID", http.StatusBadRequest)
		return
	}

	service := services.AclService{Orm:orm.NewOrm()}
	err = service.RemoveUserFromRole(userId, roleId)

	if err != nil{
		common.DisplayAppError(w, err, "Error in removing user from role", http.StatusInternalServerError)
		return
	}

	common.SendResult(w, resources.ResponseResource{Message:"User removed from role", Success:true}, http.StatusOK)
}

//Handler /users/{userId}/roles		[GET]
// Returns a []models.Resource
func GetRolesForUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.ParseInt(vars["userId"], 10, 64)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid User Id", http.StatusBadRequest)
		return
	}

	service := services.AclService{Orm:orm.NewOrm()}

	roles := service.GetRolesForUser(userId)

	common.SendResult(w, resources.ResponseResource{Data:roles, Success:true}, http.StatusOK)
}

//Handler /users/{userId}/resources	[GET]
// Returns a []models.Resource
func GetResourceForUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.ParseInt(vars["userId"], 10, 64)
	if err != nil{
		common.DisplayAppError(w, err, "Invalid User Id", http.StatusBadRequest)
		return
	}
	service := services.AclService{Orm:orm.NewOrm()}

	resourceIds := service.GetResourcesForUser(userId)

	common.SendResult(w, resources.ResponseResource{Data:resourceIds, Success:true}, http.StatusOK)

}