package tests

import (
	"context"
	"github.com/craftgate/craftgate-go-client/adapter"
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

var bankAccountTrackingClient, _ = craftgate.New("api-key", "secret-key", "http://localhost:8000")

func TestBankAccountTracking_SearchBankAccountTrackingRecords(t *testing.T) {
	request := adapter.SearchBankAccountTrackingRecordRequest{
		Page:     0,
		Size:     10,
		Currency: craftgate.Currency("TRY"),
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
