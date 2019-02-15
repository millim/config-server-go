package users

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_createUser(t *testing.T) {
	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"name": "createUserName", "password": "12345678"}`)
	req, _ := http.NewRequest("POST", "/api/users", bytes.NewReader(jsonStr))
	testRoute.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
