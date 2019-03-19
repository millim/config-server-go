package main

import (
	"config-server-go/common"
	"config-server-go/common/db"
	"config-server-go/models/migrate"
	"config-server-go/routers"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	_ "config-server-go/extend/log"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var fp = common.FlagParams

func main() {
	dbName := "server_dev.db"

	env := os.Getenv("GIN_MODE")
	if env == "release" {
		dbName = "server_pro.db"
	}

	if fp.DBPath == "" {
		fp.DBPath = "./" + dbName
	}

	linkDB, error := gorm.Open("sqlite3", fp.DBPath)
	if error != nil {
		fmt.Println("database error:")
		fmt.Println(error)
	}
	db.SetDB(linkDB)
	defer linkDB.Close()

	migrate.Run()

	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong1",
		})
	})

	routers.InitRoutes(server)
	logrus.Info("pid is -->", os.Getpid())
	logrus.Infof("server start is at %s:%s ", fp.Host, fp.Port)
	addr := fmt.Sprintf("%s:%s", fp.Host, fp.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: server,
	}

	gracehttp.Serve(srv)
	defer logrus.Info("server is close！")
}
