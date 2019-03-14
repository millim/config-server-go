package configs

import (
	"config-server-go/models/config"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateConfigFrom 传入的配置表单
type CreateConfigFrom struct {
	Env        string            `json:"env" from:"env"`
	Name       string            `json:"name" from:"name"`
	DataJSON   map[string]string `json:"data_json" from:"data_json"`
	DataString string            `json:"data_string" from:"data_string"`
	Cover      bool              `json:"cover" from:"cover"`
}

func createConfig(c *gin.Context) {
	var data CreateConfigFrom
	c.DefaultPostForm("cover", "false")
	c.ShouldBindJSON(&data)

	config, err := buildConfig(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	err = saveNewConfig(&config, data.Cover)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0})
}

func saveNewConfig(c *config.Config, cover bool) error {
	return nil
}

func buildConfig(c *CreateConfigFrom) (config.Config, error) {
	var error error
	cf := new(config.Config)

	cf.Env = c.Env
	cf.Name = c.Name

	if c.DataString != "" {
		cf.MetaData = []byte(c.DataString)
		if json.Valid(cf.MetaData) {
			error = errors.New("data_json format error")
		}
	} else if len(c.DataJSON) != 0 {
		cf.MetaData, error = json.Marshal(c.DataJSON)
	} else {
		error = errors.New("data_json and data_string only one is empty")
	}

	return *cf, error
}
