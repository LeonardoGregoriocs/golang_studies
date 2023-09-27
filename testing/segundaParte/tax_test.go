package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculaTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculaTax(0)
	assert.Error(t, err, "amout must be greater than 0")
	assert.Equal(t, 0.0, tax)
}
