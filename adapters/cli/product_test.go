package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/matheus-aloisio-lehnen/hexagonal-architecture/adapters/cli"
	mockapplication "github.com/matheus-aloisio-lehnen/hexagonal-architecture/application/mocks"
	"github.com/stretchr/testify/require"
)

func Test_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := "abc"
	productStatus := "enabled"
	productPrice := 25.99
	productName := "Product Test"

	productMock := mockapplication.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mockapplication.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s  with the name %s has been created with the price %f and status %s",
		productId, productName, productPrice, productStatus,
	)

	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	result, err = cli.Run(service, "enable", productId, "", 0)
	resultExpected = fmt.Sprintf("Product %s has been enabled.", productName)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	result, err = cli.Run(service, "disable", productId, "", 0)
	resultExpected = fmt.Sprintf("Product %s has been disabled.", productName)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	result, err = cli.Run(service, "get", productId, "", 0)
	resultExpected = fmt.Sprintf(" Product ID: %s\n Name: %s\n Price: %f\n Status: %s",
		productId, productName, productPrice, productStatus)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

}
