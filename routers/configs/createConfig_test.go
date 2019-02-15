package configs

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_createConfig(t *testing.T) {
	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"env": "default", "name": "momomo"}`)
	req, _ := http.NewRequest("POST", "/api/configs", bytes.NewReader(jsonStr))
	testRoute.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
