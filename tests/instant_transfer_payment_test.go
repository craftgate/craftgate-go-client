package tests

import (
	"context"
	"testing"

	"github.com/craftgate/craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

var instantTransferPaymentClient, _ = adapter.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func TestRetrieveActiveBanks(t *testing.T) {
	res, err := instantTransferPaymentClient.Payment.RetrieveActiveBanks(context.Background())
	require.NotEmpty(t, res.Items)
	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestInitInstantTransferAPMPayment(t *testing.T) {
	additionalParams := make(map[string]string)
	additionalParams["bankCode"] = "0"

	request := adapter.InitApmPaymentRequest{
		ApmType:        adapter.ApmType_INSTANT_TRANSFER,
		Price:          1,
		PaidPrice:      1,
		Currency:       adapter.Currency_TRY,
		PaymentGroup:   adapter.PaymentGroup_LISTING_OR_SUBSCRIPTION,
		ConversationId: "foo-bar",
		CallbackUrl:    "https://www.your-website.com/callback",
		Items: []adapter.PaymentItem{
			{
				Name:  "Item 1",
				Price: 0.6,
			},
			{
				Name:  "Item 2",
				Price: 0.4,
			},
		},
		AdditionalParams: additionalParams,
	}
	res, err := instantTransferPaymentClient.Payment.InitApmPayment(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
