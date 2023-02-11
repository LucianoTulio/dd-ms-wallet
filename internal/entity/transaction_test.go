package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, err := NewClient("Joao Lasagna", "lasagna.bolonhesa@lasa.com")
	account1 := NewAccount(client1)
	client2, err := NewClient("Junior Zoio Buneca", "jrzoio@buneca.com")
	account2 := NewAccount(client2)

	account1.Credit(10000)
	account2.Credit(10000)

	transaction, err := NewTransaction(account1, account2, 100)

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 10100.0, account2.Balance)
	assert.Equal(t, 9900.0, account1.Balance)

}

func TestCreateTransactionInsufficientBalance(t *testing.T) {
	client1, err := NewClient("Joao Lasagna", "lasagna.bolonhesa@lasa.com")
	account1 := NewAccount(client1)
	client2, err := NewClient("Junior Zoio Buneca", "jrzoio@buneca.com")
	account2 := NewAccount(client2)

	account1.Credit(10000)
	account2.Credit(10000)

	transaction, err := NewTransaction(account1, account2, 10100)

	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient funds")
	assert.Nil(t, transaction)
	assert.Equal(t, 10000.0, account2.Balance)
	assert.Equal(t, 10000.0, account1.Balance)

}
