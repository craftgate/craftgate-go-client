package main

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
	"fmt"
)

func main() {
	baseURL := "https://sandbox-api.craftgate.io"
	apiKey := "sandbox-YEhueLgomBjqsnvBlWVVuFsVhlvJlMHE"
	secretKey := "sandbox-tBdcdKVGmGupzfaWcULcwDLMoglZZvTz"
	Craftgate := CraftgateClient(apiKey, secretKey, baseURL)

	//res, _ := Craftgate.Installment.RetrieveBinNumber(adapter.RetrieveBinNumberRequest{BinNumber: "552096"})
	//fmt.Println(res)
	//
	//res, _ := Craftgate.Installment.SearchInstallments(adapter.SearchInstallmentRequest{BinNumber: "552096", Price: 100.00})
	//fmt.Println(res)

	resPaymentSearch, _ := Craftgate.PaymentReporting.SearchPayments(adapter.SearchPaymentsRequest{Currency: model.TRY})
	fmt.Println(resPaymentSearch)

	resPaymentRetrieve, _ := Craftgate.PaymentReporting.RetrievePayment(115376)
	fmt.Println(resPaymentRetrieve)
}
