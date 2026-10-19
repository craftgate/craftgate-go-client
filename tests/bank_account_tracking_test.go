package tests

import (
	"context"
	"github.com/craftgate/craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

var bankAccountTrackingClient, _ = adapter.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func TestBankAccountTracking_SearchBankAccountTrackingRecords(t *testing.T) {
	request := adapter.SearchBankAccountTrackingRecordRequest{
		Page:     0,
		Size:     10,
		Currency: adapter.Currency_TRY,
	}

	res, err := bankAccountTrackingClient.BankAccountTracking.SearchRecords(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestBankAccountTracking_RetrieveRecord(t *testing.T) {
	res, err := bankAccountTrackingClient.BankAccountTracking.RetrieveRecords(context.Background(), 1)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
