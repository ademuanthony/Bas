package resources

type ResponseResource struct {
	StatusCode int `json:"status"`
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
	Token string `json:"token"`
}