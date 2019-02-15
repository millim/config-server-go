package users

import (
	"config-server-go/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"os"
	"testing"
)

var testRoute *gin.Engine
var testDB *gorm.DB

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	testRoute = gin.Default()

	prefixGroup := testRoute.Group("/api")
	Routes(prefixGroup)

	testDB = models.InitTestDB("../../server_test.db")

	run := m.Run()
	testDB.Close()
	os.Exit(run)
}
