package resources

import "github.com/ademuanthony/Bas/models"

type (
	UserResource struct {
		Data models.User
	}

	LoginModel struct {
		Username, Password string
	}

	LoginResource struct {
		Data LoginModel
	}

	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
)
