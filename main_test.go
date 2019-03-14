package main_test

import (
	"config-server-go/common/db"
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestMain(m *testing.M) {
	dbName := "server_test.db"
	linkDB, error := gorm.Open("sqlite3", "./"+dbName)
	if error != nil {
		fmt.Println("database error:")
		fmt.Println(error)
	}
	db.SetDB(linkDB)
	gin.SetMode(gin.TestMode)
	run := m.Run()
	linkDB.Close()
	os.Exit(run)
}

func TestRun(t *testing.T) {
	linkDB := db.GetDB()
	if linkDB == nil {
		t.Fatal("database is init error")
	}
}
