package handlers

import (
	"errors"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/rodrigodealer/realtime/models"
	"github.com/rodrigodealer/realtime/services"
	"github.com/stretchr/testify/mock"
)

type clientMock struct {
	mock.Mock
}

func (o clientMock) Ping() int {
	args := o.Called()
	return args.Int(0)
}

func (o clientMock) Connect() {
}

func (o clientMock) GetUser(index string, ID string) (models.FacebookUser, error) {
	args := o.Called()
	var facebookUser models.FacebookUser
	return facebookUser, args.Error(1)
}

func (o clientMock) IndexUser(index string, user *models.FacebookUser, span opentracing.Span) {
}

type fbClientMock struct {
	mock.Mock
}

func (o fbClientMock) GetRequest(url string) (services.FbResponse, error) {
	var response = services.FbResponse{Code: 200, Body: "bla"}
	return response, errors.New("can't work with 42")
}
