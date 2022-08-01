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

	res, _ := Craftgate.Installment.RetrieveBinNumber(adapter.RetrieveBinNumberRequest{BinNumber: "552096"})
	fmt.Println(res)
	//
	//res, _ := Craftgate.Installment.SearchInstallments(adapter.SearchInstallmentRequest{BinNumber: "552096", Price: 100.00})
	//fmt.Println(res)

	//res, _ := Craftgate.Installment.SearchInstallments(adapter.SearchInstallmentRequest{BinNumber: "552096"})
	//fmt.Println(res)

	//res, _ := Craftgate.Wallet.RetrieveMemberWallet(adapter.RetrieveMemberWalletRequest{MemberId: 66988})
	//fmt.Println(res)

	//resPaymentSearch, _ := Craftgate.PaymentReporting.SearchPayments(adapter.SearchPaymentsRequest{Currency: model.TRY})
	//fmt.Println(resPaymentSearch)
	//
	//resPaymentRetrieve, _ := Craftgate.PaymentReporting.RetrievePayment(115376)
	//fmt.Println(resPaymentRetrieve)

	//resPaymentTransactionRetrieve, _ := Craftgate.PaymentReporting.RetrievePaymentTransactions(115376)
	//fmt.Println(resPaymentTransactionRetrieve)

	//resPaymentRetrieveRefunds, _ := Craftgate.PaymentReporting.RetrievePaymentRefunds(115376)
	//fmt.Println(resPaymentRetrieveRefunds)

	//resPaymentTransactionRefunds, _ := Craftgate.PaymentReporting.RetrievePaymentTransactionRefunds(115482, 99732)
	//fmt.Println(resPaymentTransactionRefunds)

	//resPaymentRefunds, _ := Craftgate.PaymentReporting.SearchPaymentRefunds(adapter.SearchPaymentRefundsRequest{})
	//fmt.Println(resPaymentRefunds)

	//resPaymentTransactionRefunds, _ := Craftgate.PaymentReporting.SearchPaymentTransactionRefunds(adapter.SearchPaymentTransactionRefundsRequest{})
	//fmt.Println(resPaymentTransactionRefunds)
}
