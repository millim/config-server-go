package user

import (
	"config-server-go/common/db"
	"crypto/sha256"
	"errors"

	"github.com/jinzhu/gorm"
)

//User 用户信息
type User struct {
	gorm.Model
	Name     string `gorm:"not null;unique"`
	Password string `gorm:"not"`
}

//CreateUser Create user info, userName  String
func (u *User) CreateUser(userName string, password string) error {
	s := sha256.New()
	s.Write([]byte(password))

	u.Name = userName
	u.Password = string(s.Sum([]byte(nil)))

	linkDB := db.GetDB()

	result := linkDB.Create(&u)

	if result.Error != nil {
		return errors.New("user name is exists")
	}
	return nil
}

//LoginUser User login, need userName and password
func (u User) LoginUser(userName string, password string) error {
	return nil
}
