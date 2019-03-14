package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	//sqlite 需要初始化
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

//SetDB set init db drive
func SetDB(_db *gorm.DB) {
	db = _db
}

//GetDB get this db drive
func GetDB() *gorm.DB {
	return db
}

//InitTestDB 测试的数据库初始化
func InitTestDB(dbFile string) *gorm.DB {
	if dbFile == "" {
		dbFile = "../../server_test.db"
	}
	os.Remove(dbFile)
	_db, error := gorm.Open("sqlite3", dbFile)
	if error != nil {
		fmt.Println(dbFile, "-----> error:")
		fmt.Println(error)
	}
	SetDB(_db)
	return _db
}
