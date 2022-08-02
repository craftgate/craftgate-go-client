package tests

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
	"fmt"
	"testing"
	"time"
)

var settlementReporting = adapter.SettlementReporting{
	Opts: model.RequestOptions{
		BaseURL:   "https://sandbox-api.craftgate.io",
		ApiKey:    "sandbox-YEhueLgomBjqsnvBlWVVuFsVhlvJlMHE",
		SecretKey: "sandbox-tBdcdKVGmGupzfaWcULcwDLMoglZZvTz",
	},
}

func TestSettlementReporting_SearchPayoutCompletedTransactions(t *testing.T) {
	request := adapter.SearchPayoutCompletedTransactionsRequest{
		SettlementFileId: 0,
		SettlementType:   "",
		StartDate:        time.Now().AddDate(0, 0, -180),
		EndDate:          time.Now(),
	}

	res, err := settlementReporting.SearchPayoutCompletedTransactions(request)
	fmt.Println(res)

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
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlementReporting_RetrievePayoutDetails(t *testing.T) {
	res, err := settlementReporting.RetrievePayoutDetails(adapter.RetrievePayoutDetailsRequest{123})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
