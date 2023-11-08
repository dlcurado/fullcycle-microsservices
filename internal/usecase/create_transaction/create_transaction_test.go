package create_transaction

import (
	"context"
	"testing"

	"github.com.br/dlcurado/fc-ws-wallet/internal/entity"
	"github.com.br/dlcurado/fc-ws-wallet/internal/event"
	"github.com.br/dlcurado/fc-ws-wallet/internal/usecase/mocks"
	"github.com.br/dlcurado/fc-ws-wallet/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTrasactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("Client 1", "client1@email.com")
	account1 := entity.NewAccount(client1, 1000)
	client2, _ := entity.NewClient("Client 2", "client2@email.com")
	account2 := entity.NewAccount(client2, 1000)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	// mockAccount := &mocks.AccountGatewayMock{}
	// mockAccount.On("FindByID", account1.ID).Return(account1, nil)
	// mockAccount.On("FindByID", account2.ID).Return(account2, nil)

	// mockTransaction := &mocks.TransactionGatewayMock{}
	// mockTransaction.On("Create", mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDto{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	dispatcher := events.NewEventDispatcher()
	event := event.NewTransactionCreated()
	ctx := context.Background()

	uc := NewCreateTransactionUseCase(mockUow, dispatcher, event)
	output, err := uc.Execute(ctx, inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	// mockAccount.AssertExpectations(t)
	// mockTransaction.AssertExpectations(t)

	// mockAccount.AssertNumberOfCalls(t, "FindByID", 2)
	// mockTransaction.AssertNumberOfCalls(t, "Create", 1)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
