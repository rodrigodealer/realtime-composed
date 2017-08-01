package handlers

import (
	"errors"

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

type fbClientMock struct {
	mock.Mock
}

func (o fbClientMock) GetRequest(url string) (services.FbResponse, error) {
	var response = services.FbResponse{Code: 200, Body: "bla"}
	return response, errors.New("can't work with 42")
}
