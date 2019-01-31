package users

import (
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	route.POST("/users", createUser)
}
