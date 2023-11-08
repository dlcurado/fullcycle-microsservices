package create_client

import (
	"testing"

	"github.com.br/dlcurado/fc-ws-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Criando um mock do Gateway
type ClientGatewayMock struct {
	mock.Mock
}

// Criando um mock das funcoes do Mock Gateway - GET
func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(0)
}

// Criando um mock das funcoes do Mock Gateway - SAVE
func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)

	uc := NewCreateClientUseCase(m)

	output, err := uc.Execute(
		CreateClientInputDTO{
			Name:  "client 1",
			Email: "client1@email.com",
		})

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "client 1", output.Name)
	assert.Equal(t, "client1@email.com", output.Email)

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
