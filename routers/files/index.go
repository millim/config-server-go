package files

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

//DownloadType 下载的类型
type DownloadType struct {
	InitStatus  bool
	ContentType string
	FileSuffix  string
}

func downloadTypes(key string) DownloadType {
	switch key {
	case "json":
		return DownloadType{InitStatus: true, ContentType: "application/json; charset=utf-8", FileSuffix: "json"}
	default:
		return DownloadType{InitStatus: false}
	}

}

func headerText(filename string) string {
	return `attachment; filename="` + filename + `"`
}

//Routes 路由初始化内容
func Routes(route *gin.RouterGroup) {

	route.GET("/files/:name/:config_file", func(c *gin.Context) {

		configFile := c.Param("config_file")
		splitFileName := strings.Split(configFile, ".")
		downloadFileType := splitFileName[len(splitFileName)-1]
		downloadType := downloadTypes(downloadFileType)

		//MarshalJSON
		m := map[string]interface{}{
			"hello": "world",
			"momo":  "oh not",
			"cc":    []int{1, 2, 3},
		}
		b, err := json.MarshalIndent(m, "", "  ")
		if err != nil {
			logrus.Error("format error ------>", err)
		}

		contentType := downloadType.ContentType
		extraHeaders := map[string]string{
			"Content-Disposition": headerText(configFile),
		}

		bLen := int64(len(b))
		c.DataFromReader(http.StatusOK, bLen, contentType, bytes.NewReader(b), extraHeaders)
		//c.String(http.StatusOK, "ok")
	})

}
