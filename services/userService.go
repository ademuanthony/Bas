package services

import (
	"errors"
	"fmt"
	"bitbucket.org/superfluxteam/pmsserver/models"
	"github.com/astaxie/beego/orm"
	"time"
	"gopkg.in/hlandau/passlib.v1"
)

type UserService struct {
	Orm orm.Ormer
}

func (this *UserService) CreateUser(user models.User) (int64, error) {
	if user.Username == ""{
		return 0, fmt.Errorf("%s is required", "Username")
	}
	exists := this.Orm.QueryTable("user").Filter("username__exact", user.Username).Exist()
	if exists {
		return 0, fmt.Errorf("The selected %s have been taken", "Username")
	}

	if user.Email == ""{
		return 0, fmt.Errorf("%s is required", "Email")
	}
	exists = this.Orm.QueryTable("user").Filter("email__exact", user.Email).Exist()
	if exists {
		return 0, fmt.Errorf("The selected %s have been taken", "Email")
	}
	//generate password hash

	/*hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("NUM: ERR: %v\n", err)
	}
	*/
	hashedPassword, err := passlib.Hash(user.Password)
	if err != nil {
		fmt.Printf("NUM: ERR: %v\n", err)
	}

	user.CreatedDate = time.Now()
	user.UpdatedDate = time.Now()
	user.PasswordHash = string(hashedPassword)
	id, err := this.Orm.Insert(&user)
	return id, err
}

func (this *UserService) GetUserById(id int64) (models.User, error) {
	var user models.User
	err := this.Orm.QueryTable(new(models.User)).Filter("id", id).One(&user)
	return user, err
}

func (this *UserService) UpdateUser(user models.User) error {
	_, err := this.Orm.Update(&user)
	return err
}

func (this *UserService) Login(username, password string) (models.User, error) {
	var user  models.User
	err := this.Orm.QueryTable(new(models.User)).Filter("username", username).One(&user)

	if err != nil {
		err := this.Orm.QueryTable(new(models.User)).Filter("email", username).One(&user)

		if err != nil {
			return user, errors.New("Invalid credentials")
		}

	}

	fmt.Printf("%v\n", username)

	newHash, err := passlib.Verify(password, user.PasswordHash)

	//err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		//panic(err)
		return models.User{}, errors.New("Invalid credentials")
	}

	if newHash != ""{
		user.PasswordHash = newHash
		this.Orm.Update(&user)
	}

	return user, nil
}

func (this *UserService) ChangePassword(userId int64, oldPassword, newPassword string) error {
	user := models.User{Id:userId}
	err := this.Orm.QueryTable(new(models.User)).Filter("id", userId).One(&user)
	if err != nil{
		return errors.New("User not found")
	}

	_, err = passlib.Verify(oldPassword, user.PasswordHash)

	if err != nil {
		return errors.New("Invalid credentials")
	}

	hashedPassword, err := passlib.Hash(newPassword) // bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("NUM: ERR: %v\n", err)
	}

	user.PasswordHash = string(hashedPassword)
	this.Orm.Update(&user)
	return nil
}

func (this *UserService) ChangePasswordForId(userId int64, newPassword string) error {
	user := models.User{Id:userId}
	err := this.Orm.QueryTable(new(models.User)).Filter("id", userId).One(&user)
	if err != nil{
		return errors.New("User not found")
	}

	hashedPassword, err := passlib.Hash(newPassword) // bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("NUM: ERR: %v\n", err)
	}

	user.PasswordHash = string(hashedPassword)
	this.Orm.Update(&user)
	return nil
}