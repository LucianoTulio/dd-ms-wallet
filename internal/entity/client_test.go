package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {

	client, err := NewClient("Joao Lasagna", "lasagna.bolonhesa@lasa.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Joao Lasagna", client.Name)
	assert.Equal(t, "lasagna.bolonhesa@lasa.com", client.Email)

}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Joao Lasagna", "lasagna.bolonhesa@lasa.com")
	err := client.Update("Joao Picanha", "picanha@pc.com")
	assert.Nil(t, err)
	assert.Equal(t, "Joao Picanha", client.Name)
	assert.Equal(t, "picanha@pc.com", client.Email)
}

func TestUpdateClientWithArgsAreInvalid(t *testing.T) {
	client, _ := NewClient("Joao Lasagna", "lasagna.bolonhesa@lasa.com")
	err := client.Update("", "picanha@pc.com")
	assert.Error(t, err, "name is required")

}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("Joao Lasagna", "lasagna.bolonhesa@lasa.com")
	account := NewAccount(client)
	err := client.AddAccount(account)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
