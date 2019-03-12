package users

import (
	"github.com/gin-gonic/gin"
)

//Routes 路由初始化内容
func Routes(route *gin.RouterGroup) {
	route.POST("/users", createUser)
}
