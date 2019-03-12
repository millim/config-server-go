package models

import (
	"config-server-go/common"
	"config-server-go/models/user"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //初始化数据库
)

//MigrateDB 数据库结构初始化
func MigrateDB() {
	db := common.GetDB()
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&Config{})
}

//InitTestDB 测试的数据库初始化
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
