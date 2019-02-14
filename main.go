package main

import (
	"config-server-go/common"
	"config-server-go/models"
	"config-server-go/routers"
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"os"
)

func main() {
	dbName := "server_dev.db"

	env := os.Getenv("GIN_MODE")
	if env == "release" {
		dbName = "server_pro.db"
	}
	db, error := gorm.Open("sqlite3", "./"+dbName)
	if error != nil {
		fmt.Println("database error:")
		fmt.Println(error)
	}
	common.SetDB(db)
	defer db.Close()

	models.MigrateDB()

	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong1",
		})
	})

	routers.InitRoutes(server)

	fmt.Println("pid is --> ", os.Getpid())
	srv := &http.Server{
		Addr:    ":3000",
		Handler: server,
	}

	gracehttp.Serve(srv)

	defer fmt.Println("server is close！")
}
