package tests

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
	"fmt"
	"testing"
)

var installment = adapter.Installment{
	Opts: model.RequestOptions{
		BaseURL:   "https://sandbox-api.craftgate.io",
		ApiKey:    "sandbox-SpqVrfuINfhbFtDEWBqQTCAhIzTEOedj",
		SecretKey: "sandbox-aJGxugIvDEdmgUYFByWAyNCrgaEpYWOw",
	},
}

func Test_RetrieveBinNumber(t *testing.T) {
	res, err := installment.RetrieveBinNumber("487074")
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_SearchInstallments(t *testing.T) {
	request := adapter.SearchInstallmentRequest{
		BinNumber: "487074",
		Price:     100,
		Currency:  model.Currency(model.TRY),
	}
	res, err := installment.SearchInstallments(request)
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
