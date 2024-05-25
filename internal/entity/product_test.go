package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	expectedName := "Produto 1"
	expectedPrice := 10.50
	p, err := NewProduct(expectedName, expectedPrice)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, expectedName, p.Name)
	assert.Equal(t, expectedPrice, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	expectedPrice := 10.50
	p, err := NewProduct("", expectedPrice)

	assert.Nil(t, p)
	assert.ErrorIs(t, err, ErrNameIsRequired)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Produto 1", 0)

	assert.Nil(t, p)
	assert.ErrorIs(t, err, ErrPriceIsRequired)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Produto 1", -1)

	assert.Nil(t, p)
	assert.ErrorIs(t, err, ErrInvalidPrice)
}
