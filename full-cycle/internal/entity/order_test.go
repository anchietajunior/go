package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorIfIDsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "ID is required")
}

func TestIfItGetsAnErrorIfPriceIsBlank(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "Price must be greater than zero")
}

func TestIfItGetsAnErrorIfTaxIsBlank(t *testing.T) {
	order := Order{ID: "123", Price: 10.0}
	assert.Error(t, order.Validate(), "Invalid Tax")
}

func TestFinalPrice(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 1.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, order.ID, "123")
	assert.Equal(t, order.Price, 10.0)
	assert.Equal(t, order.Tax, 1.0)
	order.CalculateFinalPrice()
	assert.Equal(t, order.FinalPrice, 11.0)
}
