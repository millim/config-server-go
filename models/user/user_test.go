package user

import (
	"config-server-go/common/db"
	"os"

	"testing"
)

func TestMain(m *testing.M) {
	linkDB := db.InitTestDB("../../server_test.db")
	linkDB.AutoMigrate(&User{})
	defer linkDB.Close()
	os.Exit(m.Run())
}

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
