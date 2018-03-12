package resources

import "github.com/ademuanthony/bas/models"

type (
	UserResource struct {
		Data models.User
	}

	LoginModel struct {
		Username string `json:"username"`
		Password string	`json:"password"`
	}

	LoginResource struct {
		Data LoginModel
	}

	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}

	ChangePasswordModel struct {
		Username string			`json:"username"`
		OldPassword string		`json:"old_password"`
		NewPassword string		`json:"new_password"`
		ConfirmPassword string		`json:"confirm_password"`
	}
)
