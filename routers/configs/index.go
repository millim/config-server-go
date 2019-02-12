package configs

import "github.com/gin-gonic/gin"

func Router(group *gin.RouterGroup) {
	group.POST("/configs", createConfig)
}
