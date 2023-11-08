package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	client, err := NewClient("Nome", "nome@email.com")
	account := NewAccount(client, 0)

	assert.Nil(t, err)
	assert.NotNil(t, client)

	assert.NotNil(t, account)

	assert.Equal(t, client.ID, account.Client.ID)
	assert.Equal(t, 0.0, account.Balance)
}

func TestCreateNewAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil, 0)
	assert.Nil(t, account)
}

func TestAccountCredit(t *testing.T) {
	client, err := NewClient("Nome", "nome@email.com")
	account := NewAccount(client, 0)

	assert.Nil(t, err)
	assert.NotNil(t, client)

	assert.NotNil(t, account)

	assert.Equal(t, client.ID, account.Client.ID)
	assert.Equal(t, 0.0, account.Balance)

	account.Credit(10)
	assert.Equal(t, 10.0, account.Balance)
}

func TestAccountDebit(t *testing.T) {
	client, err := NewClient("Nome", "nome@email.com")
	account := NewAccount(client, 10)

	assert.Nil(t, err)
	assert.NotNil(t, client)

	assert.NotNil(t, account)

	assert.Equal(t, client.ID, account.Client.ID)
	assert.Equal(t, 10.0, account.Balance)

	account.Debit(2)
	assert.Equal(t, 8.0, account.Balance)
}
