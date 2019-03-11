package common

import (
	"github.com/jinzhu/gorm"
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
