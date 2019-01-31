package models

import (
	"config-server-go/common"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null;unique"`
	Password string `gorm:"not"`
}

/*
Create user info, userName  String
*/
func (u *User) CreateUser(userName string, password string) error {
	s := sha256.New()
	s.Write([]byte(password))

	u.Name = userName
	u.Password = string(s.Sum([]byte(nil)))

	db := common.GetDB()
	fmt.Println("create db ===>", db)
	result := db.Create(&u)

	if result.Error != nil {
		return errors.New("User name is exists!")
	}
	return nil
}

/*
User login, need userName and password
*/
func (u User) LoginUser(userName string, password string) error {
	return nil
}
