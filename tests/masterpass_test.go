package tests

import (
	"context"
	"github.com/craftgate/craftgate-go-client/adapter"
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"testing"
)

var masterpassClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func TestMasterpass_CheckMasterpassUser(t *testing.T) {
	request := adapter.CheckMasterpassUserRequest{
		MasterpassGsmNumber: "903000000000",
	}
	res, err := masterpassClient.Masterpass.CheckMasterpassUser(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestMasterpass_GenerateMasterpassPaymentToken(t *testing.T) {
	request := adapter.MasterpassPaymentTokenGenerateRequest{
		UserId:       "masterpass-user-id",
		Msisdn:       "900000000000",
		BinNumber:    "404308",
		ForceThreeDS: true,
		CreatePayment: craftgate.MasterpassCreatePayment{
			Price:          1.25,
			PaidPrice:      1.25,
			Installment:    1,
			Currency:       craftgate.Currency_TRY,
			PaymentGroup:   craftgate.PaymentGroup_LISTING_OR_SUBSCRIPTION,
			ConversationId: "foo-bar",
			ExternalId:     "115",
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
		},
	}

	res, err := masterpassClient.Masterpass.GenerateMasterpassPaymentToken(context.Background(), request)

	require.NotNil(t, res.Token)
	require.NotNil(t, res.ReferenceId)
	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestMasterpass_CompleteMasterpassPayment(t *testing.T) {
	request := adapter.MasterpassPaymentCompleteRequest{
		ReferenceId: "referenceId",
		Token:       "token",
	}

	res, err := masterpassClient.Masterpass.CompleteMasterpassPayment(context.Background(), request)

	require.NotNil(t, res.Id)
	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestMasterpass_Init3DSMasterpassPayment(t *testing.T) {
	request := adapter.MasterpassPaymentThreeDSInitRequest{
		ReferenceId: "referenceId",
		CallbackUrl: "https://www.your-website.com/craftgate-3DSecure-callback",
	}

	res, err := masterpassClient.Masterpass.Init3DSMasterpassPayment(context.Background(), request)

	require.NotNil(t, res.ReturnUrl)
	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestMasterpass_Complete3DSMasterpassPayment(t *testing.T) {
	request := adapter.MasterpassPaymentThreeDSCompleteRequest{
		PaymentId: 1,
	}

	res, err := masterpassClient.Masterpass.Complete3DSMasterpassPayment(context.Background(), request)

	require.NotNil(t, res.Id)
	if err != nil {
		t.Errorf("Error %s", err)
	}
}
