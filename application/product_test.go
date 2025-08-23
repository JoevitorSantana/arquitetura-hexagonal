package application_test

import (
	"testing"

	"github.com/codeedu/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Test Product"
	product.Status = application.DISABLED
	product.Price = 10.0
	err := product.Enable()
	assert.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	assert.Equal(t, "product price must be greater than zero to enable it", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Test Product"
	product.Status = application.ENABLED
	product.Price = 0.0

	err := product.Disable()
	assert.Nil(t, err)

	product.Price = 10.0
	err = product.Disable()
	assert.Equal(t, "product cannot be disabled while it has a price", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Test Product"
	product.Status = application.DISABLED
	product.Price = 10.0

	isValid, err := product.IsValid()
	assert.Nil(t, err)
	assert.Equal(t, true, isValid)

	product.Status = "INVALID"
	isValid, err = product.IsValid()
	assert.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	assert.Equal(t, "the price must be greater or equal to zero", err.Error())
}
