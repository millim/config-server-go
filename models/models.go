package models

import (
	"config-server-go/common"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

func MigrateDB() {
	db := common.GetDB()
	db.AutoMigrate(&User{})
}

func InitTestDB(dbFile string) *gorm.DB {
	if dbFile == "" {
		dbFile = "../server_test.db"
	}
	os.Remove(dbFile)
	_db, error := gorm.Open("sqlite3", dbFile)
	if error != nil {
		fmt.Println(dbFile, "-----> error:")
		fmt.Println(error)
	}
	common.SetDB(_db)
	MigrateDB()
	return _db
}
