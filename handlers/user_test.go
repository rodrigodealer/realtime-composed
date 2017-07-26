package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserGet(t *testing.T) {
	req, _ := http.NewRequest("GET", "/user?u=chris", nil)
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(UserHandler)

	handler.ServeHTTP(res, req)

	assert.Equal(t, res.Body.String(), "Hello, chris")
}
