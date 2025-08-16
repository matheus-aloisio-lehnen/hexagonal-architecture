package application_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/matheus-aloisio-lehnen/hexagonal-architecture/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10
	err := product.Enable()
	require.Nil(t, err)
	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than 0", err.Error())
}

func TestProduct_Disabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0
	err := product.Disable()
	require.Nil(t, err)
	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be equal 0", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10
	product.ID = uuid.NewString()

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater than 0", err.Error())
}
