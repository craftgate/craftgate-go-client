package tests

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

var installment = adapter.Installment{
	Opts: model.RequestOptions{
		BaseURL:   "https://sandbox-api.craftgate.io",
		ApiKey:    "api-key",
		SecretKey: "secret-key",
	},
}

func Test_RetrieveBinNumber(t *testing.T) {
	res, err := installment.RetrieveBinNumber("487074")
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_SearchInstallments(t *testing.T) {
	request := adapter.SearchInstallmentsRequest{
		BinNumber: "487074",
		Price:     100,
		Currency:  model.Currency(model.TRY),
	}
	res, err := installment.SearchInstallments(request)
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
