package tests

import (
	"context"
	"testing"

	"github.com/craftgate/craftgate-go-client/adapter"
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

var compayPaymentClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func TestRetrieveActiveBanks(t *testing.T) {
	res, err := compayPaymentClient.Payment.RetrieveActiveBanks(context.Background())
	require.NotEmpty(t, res.Items)
	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestInitCompayAPMPayment(t *testing.T) {
	additionalParams := make(map[string]string)
	additionalParams["bankCode"] = "0"
    additionalParams["shopUrl"] = "your-website.com"
    additionalParams["receiptDescription"] = "Your receipt description"

	request := adapter.InitApmPaymentRequest{
		ApmType:        craftgate.ApmType_COMPAY,
		Price:          1,
		PaidPrice:      1,
		Currency:       craftgate.Currency_TRY,
		PaymentGroup:   craftgate.PaymentGroup_LISTING_OR_SUBSCRIPTION,
		ConversationId: "foo-bar",
		CallbackUrl:    "https://www.your-website.com/callback",
		Items: []craftgate.PaymentItem{
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
	res, err := compayPaymentClient.Payment.InitApmPayment(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
