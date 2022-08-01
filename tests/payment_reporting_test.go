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

func Test_SearchPayment(t *testing.T) {
	request := adapter.SearchPaymentsRequest{
		Page: 0, Size: 10,
		PaymentType:   model.PaymentType(model.CARD_PAYMENT),
		PaymentStatus: model.PaymentStatus(model.SUCCESS),
		Currency:      model.Currency(model.TRY),
		//MinCreatedDate: adapter.CraftgateTime{
		//	Time: time.Now().AddDate(0, 0, -180),
		//},
		//MaxCreatedDate: adapter.CraftgateTime{
		//	Time: time.Now(),
		//},
	}
	res, err := paymentReporting.SearchPayments(request)
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_SearchPaymentRefunds(t *testing.T) {
	request := adapter.SearchPaymentRefundsRequest{
		Page: 0, Size: 10,
		Currency: model.Currency(model.TRY),
		//MinCreatedDate: adapter.CraftgateTime{
		//	Time: time.Now().AddDate(0, 0, -180),
		//},
		//MaxCreatedDate: adapter.CraftgateTime{
		//	Time: time.Now(),
		//},
	}
	res, err := paymentReporting.SearchPaymentRefunds(request)
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_SearchPaymentTransactionRefunds(t *testing.T) {
	request := adapter.SearchPaymentTransactionRefundsRequest{
		Page: 0, Size: 10,
		Currency: model.Currency(model.TRY),
		//MinCreatedDate: adapter.CraftgateTime{
		//	Time: time.Now().AddDate(0, 0, -180),
		//},
		//MaxCreatedDate: adapter.CraftgateTime{
		//	Time: time.Now(),
		//},
	}
	res, err := paymentReporting.SearchPaymentTransactionRefunds(request)
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_RetrievePayment(t *testing.T) {
	res, err := paymentReporting.RetrievePayment(118185)
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_RetrievePaymentTransactions(t *testing.T) {
	res, err := paymentReporting.RetrievePaymentTransactions(118185)
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_RetrievePaymentRefunds(t *testing.T) {
	res, err := paymentReporting.RetrievePaymentRefunds(118185)
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_RetrievePaymentTransactionRefunds(t *testing.T) {
	res, err := paymentReporting.RetrievePaymentTransactionRefunds(118185, 101131)
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
