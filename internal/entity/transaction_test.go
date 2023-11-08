package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewTransaction(t *testing.T) {
	client1, erroC1 := NewClient("c1", "c1@email.com")
	accountC1 := NewAccount(client1, 1000.0)
	client1.AddAccount(accountC1)

	client2, erroC2 := NewClient("c2", "c2@email.com")
	accountC2 := NewAccount(client2, 1000.0)
	client2.AddAccount(accountC2)

	transaction, errorTrans := NewTransaction(accountC1, accountC2, 500.0)

	assert.Nil(t, erroC1)
	assert.Nil(t, erroC2)
	assert.Nil(t, errorTrans)

	assert.NotNil(t, transaction)
	assert.Equal(t, 500.0, accountC1.Balance)
	assert.Equal(t, 500.0, client1.Accounts[0].Balance)
	assert.Equal(t, client2.Accounts[0].Balance, 1500.0)
}

func TestCreateNewTransactionWithInsuficientFounds(t *testing.T) {
	client1, erroC1 := NewClient("c1", "c1@email.com")
	accountC1 := NewAccount(client1, 1000.0)
	client1.AddAccount(accountC1)
	assert.Nil(t, erroC1)

	client2, erroC2 := NewClient("c2", "c2@email.com")
	accountC2 := NewAccount(client2, 1000.0)
	client2.AddAccount(accountC2)
	assert.Nil(t, erroC2)

	transaction, errorTrans := NewTransaction(accountC1, accountC2, 2000.0)
	assert.NotNil(t, errorTrans)
	assert.Nil(t, transaction)
	assert.Equal(t, "insuficient funds", errorTrans.Error())
}
