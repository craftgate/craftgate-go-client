package tests

import (
	"context"
	"github.com/craftgate/craftgate-go-client/adapter"
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"testing"
)

var bkmExpressClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func TestBkm_Express_Init(t *testing.T) {
	request := adapter.InitBkmExpressRequest{
		Price:          1.25,
		PaidPrice:      1.25,
		Currency:       craftgate.Currency_TRY,
		PaymentGroup:   craftgate.PaymentGroup_LISTING_OR_SUBSCRIPTION,
		ConversationId: "foo-bar",
		Items: []craftgate.PaymentItem{
			{
				Name:       "Item 1",
				Price:      1,
				ExternalId: "1",
			},
			{
				Name:       "Item 2",
				Price:      0.25,
				ExternalId: "2",
			},
		},
	}
	res, err := bkmExpressClient.BkmExpress.Init(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestBkm_Express_Complete(t *testing.T) {
	request := adapter.CompleteBkmExpressRequest{
		Status:   true,
		Message:  "İşlem Başarılı",
		TicketId: "404308dcfdc163-0545-46d7-8f86-5a11718e56ec",
	}

	res, err := bkmExpressClient.BkmExpress.Complete(context.Background(), request)

	require.NotNil(t, res.AuthCode)
	require.NotNil(t, res.HostReference)
	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestBkm_Express_RetrievePayment(t *testing.T) {
	err := bkmExpressClient.BkmExpress.RetrievePayment(context.Background(), "dcfdc163-0545-46d7-8f86-5a11718e56ec")
	if err != nil {
		t.Fatalf("RetrievePayment failed: %v", err)
	}
}
