package common

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetDB(_db *gorm.DB) {
	db = _db
}

func GetDB() *gorm.DB {
	return db
}
