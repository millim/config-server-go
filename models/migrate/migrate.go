package migrate

import (
	"config-server-go/common/db"
	"config-server-go/models/config"
	"config-server-go/models/user"

	"github.com/jinzhu/gorm"
)

var linkDB *gorm.DB

//Run 执行数据库结构调整
func Run() {

	linkDB = db.GetDB()
	linkDB.AutoMigrate(&user.User{})
	linkDB.AutoMigrate(&config.Config{})
}
