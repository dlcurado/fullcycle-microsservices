package mocks

import (
	"github.com.br/dlcurado/fc-ws-wallet/internal/entity"
	"github.com/stretchr/testify/mock"
)

// Criando um mock do Gateway
type ClientGatewayMock struct {
	mock.Mock
}

// Criando um mock das funcoes do Mock Gateway - GET
func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

// Criando um mock das funcoes do Mock Gateway - SAVE
func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}
