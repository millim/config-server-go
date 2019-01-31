package files

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(route *gin.RouterGroup) {
	route.GET("/files/:name/:config", func(c *gin.Context) {
		fmt.Println("name =>", c.Param("name"))
		fmt.Println("config =>", c.Param("config"))
		fmt.Println("type =>", c.Param("type"))
		c.Stream()
		c.String(http.StatusOK, "ok")
	})

}
