package routers

import (
	"config-server-go/routers/files"
	"config-server-go/routers/users"
	"github.com/gin-gonic/gin"
)

func InitRoutes(route *gin.Engine) {

	prefixUsersRoute := route.Group("/api")

	users.Routes(prefixUsersRoute)
	files.Routes(prefixUsersRoute)

}
