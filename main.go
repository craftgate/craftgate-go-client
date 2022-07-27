package main

import (
	"craftgate-go-client/adapter"
	"fmt"
)

func main() {
	baseURL := "https://sandbox-api.craftgate.io"
	apiKey := "sandbox-YEhueLgomBjqsnvBlWVVuFsVhlvJlMHE"
	secretKey := "sandbox-tBdcdKVGmGupzfaWcULcwDLMoglZZvTz"
	Craftgate := CraftgateClient(apiKey, secretKey, baseURL)

	for i := 0; i < 10; i++ {
		res, _ := Craftgate.Installment.SearchInstallments(adapter.SearchInstallmentRequest{BinNumber: "487074", Price: 100.00})
		fmt.Println(res)
	}

	//resPaymentSearch, _ := Craftgate.PaymentReporting.SearchPayments(adapter.SearchPaymentsRequest{Currency: model.TRY})
	//fmt.Println(resPaymentSearch)
}
