package tests

import (
	"context"
	"github.com/craftgate/craftgate-go-client/adapter"
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
	"github.com/stretchr/testify/require"
	"testing"
)

var juzdanClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func Test_InitJuzdanPayment(t *testing.T) {
	request := adapter.InitJuzdanPaymentRequest{
		price:          1.25,
		paidPrice:      1.25,
		currency:       craftgate.Currency_TRY,
		paymentGroup:   craftgate.PaymentGroup_LISTING_OR_SUBSCRIPTION,
		conversationId: "foo-bar",
		externalId:     "115",
		callbackUrl:    "www.test.com",
		paymentPhase:   craftgate.PaymentPhase_AUTH,
		paymentChannel: "test",
		buyerMemberId:  1,
		bankOrderId:    "test",
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
		clientType:     craftgate.ClientType.W,
		loanCampaignId: 1,
	}

	res, err := juzdanClient.Juzdan.InitJuzdanPayment(context.Background(), request)

	require.NotNil(t, res.JuzdanQrUrl)
	require.NotNil(t, res.ReferenceId)
	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func Test_RetrieveJuzdanPayment(t *testing.T) {
	res, err := juzdanClient.Juzdan.RetrieveJuzdanPayment(context.Background(), "test-reference-id")
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
