package configs

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createConfigNotData(t *testing.T) {
	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"env": "default", "name": "momomo","data_string":"","data_json":{},"cover": false}`)
	req, _ := http.NewRequest("POST", "/api/configs", bytes.NewReader(jsonStr))
	testRoute.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, true, strings.Contains(string(body), "error"))
	assert.Equal(t, 200, w.Code)
}

func Test_createConfig(t *testing.T) {
	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"env": "default", "name": "momomo","data_string":"","data_json":{"database_name":"test","database_host":"127.0.0.1"},"cover": false}`)
	req, _ := http.NewRequest("POST", "/api/configs", bytes.NewReader(jsonStr))
	testRoute.ServeHTTP(w, req)
	t.Logf("body value ----> %v", w.Body)
	assert.Equal(t, 200, w.Code)
}
