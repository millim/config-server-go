package configs

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateConfigFrom 传入的配置表单
type CreateConfigFrom struct {
	Env        string            `json:"env" from:"env"`
	Name       string            `json:"name" from:"name"`
	DataJSON   map[string]string `json:"data_json" from:"data_json"`
	DataString string            `json:"string" from:"string"`
	Cover      bool              `json:"cover" from:"cover"`
}

func createConfig(c *gin.Context) {
	var data CreateConfigFrom
	c.ShouldBindJSON(&data)
	fmt.Println(data)
	c.JSON(http.StatusOK, gin.H{"code": 0})
}
