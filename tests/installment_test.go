package tests

import (
	"context"
	"github.com/craftgate/craftgate-go-client/adapter"
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
	"github.com/stretchr/testify/require"
	"testing"
)

var installmentClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func Test_RetrieveBinNumber(t *testing.T) {
	res, err := installmentClient.Installment.RetrieveBinNumber(context.Background(), "487074")

	require.Equal(t, *res.BinNumber, "487074")
	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_SearchInstallments(t *testing.T) {
	request := adapter.SearchInstallmentsRequest{
		BinNumber: "487074",
		Price:     100,
		Currency:  craftgate.Currency_TRY,
	}
	res, err := installmentClient.Installment.SearchInstallments(context.Background(), request)

	require.Equal(t, *res.Items[0].BinNumber, "487074")
	if err != nil {
		t.Errorf("Error %s", err)
	}
}
