package resources

type CreateResourceResource struct {
	Key string `json:"key"`
	ApplicationId int64 `json:"application_id"`
}

type CreateResourcesInputDto struct {
	ApplicationId int64 `json:"application_id"`
	Keys []string `json:"keys"`
}

type CreateResourcesOutputDto struct {
	Key	string		`json:"key"`
	Message	string		`json:"message"`
	Success	bool		`json:"success"`
}

type AddResourcesToRoleInputDto struct {
	RoleId		int64 		`json:"role_id"`
	ResourceIds	[]int64		`json:"resource_ids"`
}