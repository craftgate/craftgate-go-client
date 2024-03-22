package tests

import (
	"context"
	"github.com/craftgate/craftgate-go-client/adapter"
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
	"github.com/stretchr/testify/require"
    "github.com/davecgh/go-spew/spew"
    "testing"
)

var juzdanClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func Test_InitJuzdanPayment(t *testing.T) {
	request := adapter.InitJuzdanPaymentRequest{
		Price:          1.25,
		PaidPrice:      1.25,
		Currency:       craftgate.Currency_TRY,
		PaymentGroup:   craftgate.PaymentGroup_LISTING_OR_SUBSCRIPTION,
		ConversationId: "foo-bar",
		ExternalId:     "115",
		CallbackUrl:    "www.test.com",
		PaymentPhase:   craftgate.PaymentPhase_AUTH,
		PaymentChannel: "test",
		BankOrderId:    "test",
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
		ClientType:     craftgate.ClientType_W,
		LoanCampaignId: 1,
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
