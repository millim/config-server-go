package models

import "github.com/jinzhu/gorm"

type Config struct {
	gorm.Model
	Env      string `gorm:"default:'default'"`
	Name     string `gorm:"not null;index"`
	MetaData []byte
}
