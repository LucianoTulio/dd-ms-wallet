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
