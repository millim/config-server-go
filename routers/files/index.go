package files

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"net/http"
	"strings"
)

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
		json := jsoniter.ConfigCompatibleWithStandardLibrary
		b, err := json.MarshalIndent(m, "", "  ")
		if err != nil {
			fmt.Println("format error ------>", err)
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
