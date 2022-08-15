package tests

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
	"github.com/davecgh/go-spew/spew"
	"testing"
	"time"
)

var settlementReporting = adapter.SettlementReporting{
	Opts: model.RequestOptions{
		BaseURL:   "https://sandbox-api.craftgate.io",
		ApiKey:    "api-key",
		SecretKey: "secret-key",
	},
}

func TestSettlementReporting_SearchPayoutCompletedTransactions(t *testing.T) {
	request := adapter.SearchPayoutCompletedTransactionsRequest{
		StartDate: time.Now().AddDate(0, 0, -180),
		EndDate:   time.Now(),
	}

	res, err := settlementReporting.SearchPayoutCompletedTransactions(request)
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlementReporting_SearchPayoutBouncedTransactions(t *testing.T) {
	request := adapter.SearchPayoutBouncedTransactionsRequest{
		StartDate: time.Now().AddDate(0, 0, -180),
		EndDate:   time.Now(),
	}

	res, err := settlementReporting.SearchPayoutBouncedTransactions(request)
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlementReporting_RetrievePayoutDetails(t *testing.T) {
	res, err := settlementReporting.RetrievePayoutDetails(adapter.RetrievePayoutDetailsRequest{PayoutDetailId: 49})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
