package models

import (
	"config-server-go/common"
)

func MigrateDB() {
	db := common.GetDB()
	db.AutoMigrate(&User{})
}
