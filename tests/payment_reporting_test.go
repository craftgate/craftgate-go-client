package tests

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
	"fmt"
	"testing"
	"time"
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
		MinCreatedDate: adapter.Time{
			Time: time.Now().AddDate(0, 0, -180),
		},
		MaxCreatedDate: adapter.Time{
			Time: time.Now(),
		},
	}
	res, err := paymentReporting.SearchPayments(request)
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
