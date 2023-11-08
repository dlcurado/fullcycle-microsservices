package mocks

import (
	"github.com.br/dlcurado/fc-ws-wallet/internal/entity"
	"github.com/stretchr/testify/mock"
)

// Criando um mock do Gateway
type AccountGatewayMock struct {
	mock.Mock
}

// Criando um mock das funcoes do Mock Gateway - GET
func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

// Criando um mock das funcoes do Mock Gateway - SAVE
func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) UpdateBalance(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}
