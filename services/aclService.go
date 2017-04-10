package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/Bas/models"
	"errors"
	"fmt"
	"strconv"
)

type AclService struct {
	Orm orm.Ormer
}

func (this *AclService) CreateResource(resource models.Resource) (models.Resource, error) {
	if this.Orm.QueryTable("resource").Filter("key", resource.Key).Exist() {
		return resource, errors.New("Resource with the same key already exists")
	}

	_, err := this.Orm.Insert(&resource)

	return resource, err
}

func (this *AclService) GetResourceById(id int64) (models.Resource, error) {
	var resource models.Resource
	err := this.Orm.QueryTable("resource").Filter("id", id).One(&resource)

	return resource, err
}

func (this *AclService) DeleteResourc(id int64) (string, error) {
	resource, err := this.GetResourceById(id)
	if err != nil{
		return "Resource not found", err
	}
	_, err = this.Orm.Delete(resource)
	if err != nil{
		return "Error in delete resource", err
	}
	return "Resource Deleted", nil
}

func (this *AclService) GetResourceByName(name struct{}) (models.Resource, error) {
	var resource models.Resource
	err := this.Orm.QueryTable("resource").Filter("name", name).One(&resource)

	return resource, err
}

func (this *AclService) GetAllResources() ([]models.Resource, error) {
	var resources []models.Resource
	_, err := this.Orm.QueryTable("resource").RelatedSel().All(&resources)

	return resources, err
}

// Create a role in the db
func (this AclService) CreateRole(role models.Role) (models.Role, error) {
	if this.Orm.QueryTable("role").Filter("name", role.Name).Exist(){
		return role, errors.New("A role with the same name already exists")
	}
	_, err := this.Orm.Insert(&role)
	return role, err
}

func (this AclService) GetRoles() []models.Role {
	var roles []models.Role
	this.Orm.QueryTable("role").All(&roles)
	return roles
}

func (this AclService) GetRoleByName(name string) (models.Role, error){
	var role models.Role
	err := this.Orm.QueryTable("role").Filter("name", name).One(&role)
	return role, err
}

func (this AclService) GetRoleById(id int64) (models.Role, error) {
	var role models.Role
	err := this.Orm.QueryTable("role").Filter("id", id).One(&role)
	return role, err
}

func (this AclService) DeleteRole(roleId int64) error {
	role, _ := this.GetRoleById(roleId);
	_, err := this.Orm.Delete(&role);
	return err;
}

// Role resource
func (this *AclService) AddResourceToRole(resourceId int64, roleId int64) (int64, error) {
	var role = new(models.Role)
	role.Id = roleId

	var resource = new(models.Resource)
	resource.Id = resourceId

	roleResource := models.RoleResource{RoleId:roleId, Resource:resource}

	//don't add a resource to a role twice
	if this.Orm.QueryTable(new(models.RoleResource)).Filter("role_id", roleId).Filter("resource_id", resourceId).Exist(){
		return 0, errors.New("The selected resource is already in this role")
	}

	return this.Orm.Insert(&roleResource)
}

func (this *AclService) GetResourceInRole(roleId int64) []*models.Resource {

	var roleResources []models.RoleResource

	this.Orm.QueryTable(new(models.RoleResource)).RelatedSel().Filter("role_id", roleId).All(&roleResources)

	var resources = make([]*models.Resource, len(roleResources))

	for index, roleResource := range roleResources{
		resources[index] = roleResource.Resource
	}
	return resources
}

func (this *AclService) RemoveResourceFromRole(resourceId int64, roleId int64) error {
	var roleResource models.RoleResource
	err := this.Orm.QueryTable(new(models.RoleResource)).Filter("role_id", roleId).Filter("resource_id", resourceId).One(&roleResource)
	if err != nil{
		return errors.New("The specified resource is not in the given role")
	}
	_, err = this.Orm.Delete(&roleResource)
	return err
}

func (this *AclService) GetResourcesByRoleId(roleId int64) []*models.Resource {
	var roleResources []models.RoleResource
	_, err := this.Orm.QueryTable(new(models.RoleResource)).Filter("role_id", roleId).RelatedSel("Resource").All(roleResources)

	resources := make([]*models.Resource, len(roleResources))
	if err != nil{
		return resources
	}

	for index, roleResource := range roleResources{
		resources[index] = roleResource.Resource
	}
	return resources
}

//User role
func (this *AclService) AddUserToRole(userId int64, roleId int64) error {
	var user models.User
	err := this.Orm.QueryTable(new(models.User)).Filter("id", userId).One(&user)
	if err != nil{
		return errors.New("Invalid user Id")
	}
	var role models.Role
	err = this.Orm.QueryTable(new(models.Role)).Filter("id", roleId).One(&role)
	if err != nil{
		return errors.New("Invalid role Id")
	}
	if this.Orm.QueryTable(new(models.UserRole)).Filter("user_id", userId).Filter("role_id", roleId).Exist(){
		return errors.New("The specified user is alreay in the selected role")
	}
	userRole := models.UserRole{Role:&role, User:&user}
	_, err = this.Orm.Insert(&userRole)
	return err
}

func (this *AclService) GetUsersInRole(roleId int64) []*models.User {
	var userRoles []models.UserRole
	var users []*models.User
	_, err := this.Orm.QueryTable(new(models.UserRole)).RelatedSel().Filter("role_id", roleId).All(&userRoles)

	if err != nil{
		return users
	}

	users = make([]*models.User, len(userRoles))
	for index, userRole := range userRoles{
		users[index] = userRole.User
	}
	return users
}

func (this *AclService) RemoveUserFromRole(userId int64, roleId int64) error {
	var userRole models.UserRole
	err := this.Orm.QueryTable(new(models.UserRole)).Filter("role_id", roleId).Filter("user_id", userId).One(&userRole)
	if err != nil{
		return errors.New("The specified user is not in the given role")
	}
	_, err = this.Orm.Delete(&userRole)
	return err
}

func (this *AclService) GetRolesForUser(userId int64) []*models.Role {
	var userRoles []models.UserRole
	var roles []*models.Role

	_, err := this.Orm.QueryTable(new(models.UserRole)).RelatedSel().Filter("user_id", userId).All(&userRoles)
	if err != nil{
		return roles
	}

	roles = make([]*models.Role, len(userRoles))
	for index, userRole := range userRoles{
		roles[index] = userRole.Role
	}
	return roles
}

func (this *AclService) GetResourcesForUser(userId int64) []int64 {
	sql := "SELECT role_resource.resource_id from role_resource INNER JOIN role ON role_resource.role_id = role.id INNER JOIN user_role ON user_role.role_id = role.id" +
		" WHERE user_role.user_id = ?"
	resourceIdParam := []orm.Params{}
	this.Orm.Raw(sql, userId).Values(&resourceIdParam)

	fmt.Printf("Resources %v\n", resourceIdParam)

	ids := make([]int64, len(resourceIdParam))
	for index, param := range resourceIdParam{
		value := param["resource_id"].(string)
		id, err := strconv.ParseInt(value, 10, 64)
		if err != nil{
			panic(err)
		}
		ids[index] = id
	}

	return ids
}