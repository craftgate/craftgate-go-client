package tests

import (
	"context"
	"github.com/craftgate/craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
	"testing"
	"time"
)

var settlementReportingClient, _ = adapter.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func TestSettlementReporting_SearchPayoutCompletedItems(t *testing.T) {
	request := adapter.SearchPayoutCompletedItemsRequest{
		StartDate: time.Now().AddDate(0, 0, -180),
		EndDate:   time.Now(),
        Page:      0,
        Size:      10,
	}

	res, err := settlementReportingClient.SettlementReporting.SearchPayoutCompletedItems(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlementReporting_SearchPayoutBouncedTransactions(t *testing.T) {
	request := adapter.SearchPayoutBouncedTransactionsRequest{
		StartDate: time.Now().AddDate(0, 0, -180),
		EndDate:   time.Now(),
	}

	res, err := settlementReportingClient.SettlementReporting.SearchPayoutBouncedTransactions(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlementReporting_RetrieveSettlementRow(t *testing.T) {
	res, err := settlementReportingClient.SettlementReporting.RetrieveSettlementRow(context.Background(), 123)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlementReporting_SearchPayoutRows(t *testing.T) {
	request := adapter.SearchPayoutRowRequest{
		FileStatus: adapter.FileStatus_CREATED,
		StartDate:  time.Now().AddDate(0, 0, -180),
		EndDate:    time.Now(),
	}
	res, err := settlementReportingClient.SettlementReporting.SearchPayoutRows(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlementReporting_RetrievePaymentTransactionPayout(t *testing.T) {
	res, err := settlementReportingClient.SettlementReporting.RetrievePaymentTransactionPayout(context.Background(), 123)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlementReporting_SearchPaymentTransactionPayouts(t *testing.T) {
	request := adapter.SearchPaymentTransactionPayoutsRequest{
		PaymentTransactionIds: []int64{1, 2},
	}
	res, err := settlementReportingClient.SettlementReporting.SearchPaymentTransactionPayouts(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
