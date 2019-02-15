package configs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateConfigFrom struct {
	Env        string            `json:"env" from:"env"`
	Name       string            `json:"name" from:"name"`
	DataJson   map[string]string `json:"data_json" from:"data_json"`
	DataString string            `json:"string" from:"string"`
	Cover      bool              `json:"cover" from:"cover"`
}

func createConfig(c *gin.Context) {
	var data CreateConfigFrom
	c.ShouldBindJSON(&data)
	fmt.Println(data)
	c.JSON(http.StatusOK, gin.H{"code": 0})
}
