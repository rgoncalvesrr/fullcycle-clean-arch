package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	expectedName := "Usuario X"
	expectedEMail := "usuario@dominio.com"
	expectedPassword := "123456"

	u, e := NewUser(expectedName, expectedEMail, expectedPassword)

	assert.Nil(t, e)
	assert.NotNil(t, u)
	assert.NotEmpty(t, u.ID)
	assert.Equal(t, expectedName, u.Name)
	assert.Equal(t, expectedEMail, u.EMail)
	assert.NotEmpty(t, u.Password)
}

func TestUser_ValidatePassword(t *testing.T) {
	expectedName := "Usuario X"
	expectedEMail := "usuario@dominio.com"
	expectedPassword := "123456"

	u, e := NewUser(expectedName, expectedEMail, expectedPassword)

	assert.Nil(t, e)
	assert.NotNil(t, u)
	assert.True(t, u.ValidatePassword(expectedPassword))
	assert.False(t, u.ValidatePassword("SenhaIncorreta"))
	assert.NotEqual(t, expectedPassword, u.Password)
}
