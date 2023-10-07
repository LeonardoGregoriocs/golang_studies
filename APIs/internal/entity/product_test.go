package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProcuct(t *testing.T) {
	p, err := NewProduct("Product 1", 10)
	assert.Nil(t, err)                   // Valido que o err é nil
	assert.NotNil(t, p)                  // Valido que minha váriavel P não é nil
	assert.NotEmpty(t, p.ID)             // Valido que o ID do produto não é vazio
	assert.Equal(t, "Product 1", p.Name) // Valido o nome
	assert.Equal(t, 10, p.Price)         // Valido o preço
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Product 1", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Product 1", -10)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
