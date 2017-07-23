package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rodrigodealer/realtime/server"
	"github.com/stretchr/testify/assert"
)

func TestSuccessfulSubscriptionPost(t *testing.T) {
	req, _ := http.NewRequest("POST", "/subscription?u=chris", nil)
	res := httptest.NewRecorder()
	server.Server().ServeHTTP(res, req)

	assert.Equal(t, res.Code, 200)
	assert.Equal(t, res.Body.String(), "Hello, chris")
}
