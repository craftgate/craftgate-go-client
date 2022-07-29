package tests

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
	"fmt"
	"testing"
)

var paymentReporting = adapter.PaymentReporting{
	Opts: model.RequestOptions{
		BaseURL:   "https://sandbox-api.craftgate.io",
		ApiKey:    "sandbox-YEhueLgomBjqsnvBlWVVuFsVhlvJlMHE",
		SecretKey: "sandbox-tBdcdKVGmGupzfaWcULcwDLMoglZZvTz",
	},
}

// test function
func Test_SearchPayment(t *testing.T) {
	res, err := paymentReporting.SearchPayments(adapter.SearchPaymentsRequest{})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
