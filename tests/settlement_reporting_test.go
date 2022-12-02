package tests

import (
    "context"
    "github.com/craftgate/craftgate-go-client/adapter"
    craftgate "github.com/craftgate/craftgate-go-client/adapter"
    "github.com/davecgh/go-spew/spew"
    "testing"
    "time"
)

var settlementReportingClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func TestSettlementReporting_SearchPayoutCompletedTransactions(t *testing.T) {
    request := adapter.SearchPayoutCompletedTransactionsRequest{
        StartDate: time.Now().AddDate(0, 0, -180),
        EndDate:   time.Now(),
    }

    res, err := settlementReportingClient.SettlementReporting.SearchPayoutCompletedTransactions(context.Background(), request)
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

func TestSettlementReporting_RetrievePayoutDetails(t *testing.T) {
    res, err := settlementReportingClient.SettlementReporting.RetrievePayoutDetails(context.Background(), 123)
    _, _ = spew.Printf("%#v\n", res)

    if err != nil {
        t.Errorf("Error %s", err)
    }
}
