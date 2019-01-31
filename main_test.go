package main_test

import (
	"config-server-go/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	dbName := "server_test.db"
	db, error := gorm.Open("sqlite3", "./"+dbName)
	if error != nil {
		fmt.Println("database error:")
		fmt.Println(error)
	}
	common.SetDB(db)
	gin.SetMode(gin.TestMode)
	run := m.Run()
	db.Close()
	os.Exit(run)
}

func TestRun(t *testing.T) {
	db := common.GetDB()
	if db == nil {
		t.Fatal("database is init error")
	}
}
