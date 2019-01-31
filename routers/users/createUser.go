package users

import (
	"config-server-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserFrom struct {
	Name     string `form:"name" json:"name"`
	Password string `json:"password"`
}

func createUser(c *gin.Context) {
	var data CreateUserFrom
	c.ShouldBindJSON(&data)
	user := new(models.User)
	err := user.CreateUser(data.Name, data.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"name": user.Name, "id": user.ID})
	}

}
