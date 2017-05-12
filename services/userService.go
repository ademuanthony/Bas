package services

import (
	"errors"
	"fmt"
	"github.com/ademuanthony/Bas/models"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService struct {
	Orm orm.Ormer
}

func (this *UserService) CreateUser(user models.User) (int64, error) {
	exists := this.Orm.QueryTable("user").Filter("username__exact", user.Username).Exist()
	if exists {
		return 0, fmt.Errorf("The selected %s have been taken", "Username")
	}
	exists = this.Orm.QueryTable("user").Filter("email__exact", user.Email).Exist()
	if exists {
		return 0, fmt.Errorf("The selected %s have been taken", "Email")
	}
	//generate password hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("NUM: ERR: %v\n", err)
	}
	user.CreatedDate = time.Now()
	user.UpdatedDate = time.Now()
	user.PasswordHash = string(hashedPassword)
	id, err := this.Orm.Insert(&user)
	return id, err
}

func (this *UserService) Login(username, password string) (models.User, error) {
	user := models.User{Username: username}
	err := this.Orm.QueryTable(new(models.User)).Filter("username", username).One(&user)

	if err != nil {
		return user, errors.New("Invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return user, errors.New("Invalid credentials")
	}
	return user, nil
}

func (this *UserService) ChangePassword(username, newPassword string) error {
	user := models.User{Username:username}
	err := this.Orm.QueryTable(new(models.User)).Filter("username", username).One(&user)
	if err != nil{
		return errors.New("User not found")
	}

	/*if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil{
		return errors.New("Invalid credentials")
	}*/
	haspPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil{
		return errors.New("Unknown error has occured")
	}
	user.PasswordHash = string(haspPassword)
	this.Orm.Update(&user)
	return nil
}
