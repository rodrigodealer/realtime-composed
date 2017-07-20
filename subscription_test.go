package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rodrigodealer/realtime/server"
	"github.com/stretchr/testify/assert"
)

func TestSuccessfulSubscriptionGet(t *testing.T) {
	req, _ := http.NewRequest("GET", "/subscription?hub.verify_token=moi&hub.challenge=bla", nil)
	res := httptest.NewRecorder()
	server.Server().ServeHTTP(res, req)

	assert.Equal(t, res.Code, 200)
	assert.Equal(t, res.Body.String(), "bla")
}

func TestFailedSubscriptionGet(t *testing.T) {
	req, _ := http.NewRequest("GET", "/subscription?hub.verify_token=moa&hub.challenge=bla", nil)
	res := httptest.NewRecorder()
	server.Server().ServeHTTP(res, req)

	assert.Equal(t, res.Body.String(), "")
	assert.Equal(t, res.Code, 401)
}
