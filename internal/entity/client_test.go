package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Nome", "nome@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Nome", client.Name)
	assert.Equal(t, "nome@email.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.Nil(t, client)
	assert.NotNil(t, err)
}

func TestUpdateClient(t *testing.T) {
	client, err := NewClient("Nome", "nome@email.com")
	updatedAt := client.UpdatedAt
	client.Update("Novo Nome", "novoemail@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Novo Nome", client.Name)
	assert.Equal(t, "novoemail@email.com", client.Email)
	assert.NotEqual(t, updatedAt, client.UpdatedAt)
}

func TestClientAddAccount(t *testing.T) {

	client, err := NewClient("Nome", "nome@email.com")

	account := NewAccount(client, 10)

	error := client.AddAccount(account)

	assert.Nil(t, err)
	assert.Nil(t, error)
	assert.Equal(t, 1, len(client.Accounts))
}
