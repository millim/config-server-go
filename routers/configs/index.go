package configs

import "github.com/gin-gonic/gin"

//Routes 路由初始化内容
func Routes(group *gin.RouterGroup) {
	group.POST("/configs", createConfig)
}
