package models

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := new(User)
	result1 := user.CreateUser("user_name", "password")
	if result1 != nil {
		t.Fatal("create user info error")
	}
	result2 := user.CreateUser("user_name", "password")
	if result2 == nil {
		t.Fatal("re create user must error")
	}

}
