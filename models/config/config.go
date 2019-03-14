package config

import "github.com/jinzhu/gorm"

//Config 数据配置文件
type Config struct {
	gorm.Model
	Env      string `gorm:"default:'default'"`
	Name     string `gorm:"not null;index"`
	MetaData []byte
}
