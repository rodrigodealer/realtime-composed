package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSuccessfulHealthcheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	res := httptest.NewRecorder()
	client := new(clientMock)
	client.On("Ping").Return(200)

	handler := http.HandlerFunc(HealthcheckHandler(client))
	handler.ServeHTTP(res, req)

	assert.Equal(t, res.Code, 200)
	assert.Equal(t, res.Body.String(), "WORKING")
}

func TestFailedHealthcheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	res := httptest.NewRecorder()
	client := new(clientMock)
	client.On("Ping").Return(500)

	handler := http.HandlerFunc(HealthcheckHandler(client))
	handler.ServeHTTP(res, req)

	assert.Equal(t, res.Code, 500)
	assert.Equal(t, res.Body.String(), "{\"status\":\"FAILED\",\"services\":[{\"name\":\"elasticsearch\",\"state\":\"FAILED\",\"code\":500}]}\n")
}

type clientMock struct {
	mock.Mock
}

func (o clientMock) Ping() int {
	args := o.Called()
	return args.Int(0)
}

func (o clientMock) Connect() {
}
