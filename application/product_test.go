package application_test

import (
	"testing"

	"github.com/luisescx/goHexagonalArchitecture/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.ID = "1"
	product.Name = "P1"
	product.Status = application.DISABLED
	product.Price = 10.0

	err := product.Enable()
	require.Nil(t, err)
	
	product.Price = 10.0
	err = product.Enable()
	require.EqualError(t, err, "THE PRICE MUST BE GREATER THAN ZERO TO ENABLE THE PRODUCT")
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.ID = "1"
	product.Name = "P1"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10.0
	err = product.Disable()
	require.EqualError(t, err, "THE PRICE MUST BE ZERO IN ORDER TO HAVE THE PRODUCT DISABLED")
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "P1"
	product.Price = 10
	product.Status = application.DISABLED

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.EqualError(t, err, "THE STATUS MUST BE ENABLED OR DISABLED")

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)
	
	product.Price = -10.0
	_, err = product.IsValid()
	require.EqualError(t, err, "THE PRICE MUST GREATER OR EQUAL TO ZERO")
}

func TestProduct_GetId(t *testing.T) {
	product := application.Product{}
	ID := uuid.NewV4().String()
	product.ID = ID
	product.Name = "P1"
	product.Price = 10.0
	product.Status = application.DISABLED

	pID := product.GetID()

	require.Equal(t, ID, pID)
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "P1"
	product.Price = 10.0
	product.Status = application.DISABLED

	productName := product.GetName()

	require.Equal(t, "P1", productName)
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "P1"
	product.Price = 10.0
	product.Status = application.DISABLED

	productStatus := product.GetStatus()

	require.Equal(t, application.DISABLED, productStatus)
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "P1"
	product.Price = 10.0
	product.Status = application.DISABLED

	productPrice := product.GetPrice()

	require.Equal(t, 10.0, productPrice)
}