package tests

import (
	"context"
	"craftgate-go-client/adapter"
	craftgate "craftgate-go-client/adapter"
	"testing"
)

var installmentClient, _ = craftgate.New("sandbox-YEhueLgomBjqsnvBlWVVuFsVhlvJlMHE", "sandbox-tBdcdKVGmGupzfaWcULcwDLMoglZZvTz", "https://sandbox-api.craftgate.io")

func Test_RetrieveBinNumber(t *testing.T) {
	res, err := installmentClient.Installment.RetrieveBinNumber(context.Background(), "487074")

	if err != nil || *res.BinNumber != "487074" {
		t.Errorf("Error %s", err)
	}
}

func Test_SearchInstallments(t *testing.T) {
	request := adapter.SearchInstallmentsRequest{
		BinNumber: "487074",
		Price:     100,
		Currency:  craftgate.Currency(craftgate.TRY),
	}
	res, err := installmentClient.Installment.SearchInstallments(context.Background(), request)

	if err != nil && *res.Items[0].BinNumber != "487074" {
		t.Errorf("Error %s", err)
	}
}
