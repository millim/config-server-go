package models

import (
	"config-server-go/common"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	dbFile := "../server_test.db"
	os.Remove(dbFile)
	db, error := gorm.Open("sqlite3", dbFile)
	if error != nil {
		fmt.Println("database error:")
		fmt.Println(error)
	}
	common.SetDB(db)
	MigrateDB()
	run := m.Run()
	db.Close()
	os.Exit(run)
}
