package main

import (
	"config-server-go/common/db"
	"config-server-go/models/migrate"
	"config-server-go/routers"
	"fmt"
	"net/http"
	"os"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	dbName := "server_dev.db"

	env := os.Getenv("GIN_MODE")
	if env == "release" {
		dbName = "server_pro.db"
	}
	linkDB, error := gorm.Open("sqlite3", "./"+dbName)
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

	fmt.Println("pid is --> ", os.Getpid())
	srv := &http.Server{
		Addr:    ":3000",
		Handler: server,
	}

	gracehttp.Serve(srv)

	defer fmt.Println("server is close！")
}
