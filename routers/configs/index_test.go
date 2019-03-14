package configs

import (
	"config-server-go/common/db"
	"config-server-go/models/migrate"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var testRoute *gin.Engine
var testDB *gorm.DB

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	testRoute = gin.Default()

	prefixGroup := testRoute.Group("/api")
	Routes(prefixGroup)

	testDB = db.InitTestDB("../../server_test.db")
	migrate.Run()
	run := m.Run()
	testDB.Close()
	os.Exit(run)
}
