package data

import (
	"github.com/ademuanthony/Bas/models"
	"github.com/ademuanthony/Bas/common"
	"fmt"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"github.com/astaxie/beego/orm"
)

type UserRepository struct {
	Db orm.Ormer
}

func (this UserRepository) CreateUser(user models.User) (int64, error) {
	var existingUser models.User
	err := common.Orm.Where("user_name = ?", user.Username).Find(&existingUser)
	if err == nil{
		return 0, fmt.Errorf("The selected %s have been taken", "Username")
	}
	err = common.Orm.Where("email = ?", user.Email).Find(&existingUser)
	if err == nil{
		return 0, fmt.Errorf("The selected %s have been taken", "Email")
	}
	err = common.Orm.Save(&user)
	return user.Id, err
}

func (this UserRepository) Login(username, password string) (models.User, error) {
	var user models.User
	var err error
	err = common.Orm.Where("user_name = ?", username).Find(&user)
	if err != nil{
		return user, errors.New("Invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil{
		return user, errors.New("Invalid credentials")
	}
	return user, nil
}