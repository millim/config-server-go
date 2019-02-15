package configs

import "github.com/gin-gonic/gin"

func Routes(group *gin.RouterGroup) {
	group.POST("/configs", createConfig)
}
