package db

import (
	"fmt"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//var db *gorm.DB

func TestMain(m *testing.M) {
	m.Run()
	if db != nil {
		defer db.Close()
	}
	os.Exit(0)
}

func TestSetDB(t *testing.T) {
	dbName := "server_test.db"
	_db, error := gorm.Open("sqlite3", "../"+dbName)
	if error != nil {
		t.Fatal(error)
	}
	if db != nil {
		t.Fatal("db must is dont init")
	}
	SetDB(_db)
	if db == nil {
		t.Fatal("db must is init()")
	}
}

func TestGetDB(t *testing.T) {
	testDB := GetDB()
	fmt.Println(testDB)
}
