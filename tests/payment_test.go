package tests

import (
	"context"
	"github.com/craftgate/craftgate-go-client/adapter"
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"testing"
)

var paymentClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func TestPayment_CreatePayment(t *testing.T) {
	request := adapter.CreatePaymentRequest{
		Price:          1.25,
		PaidPrice:      1.25,
		Installment:    1,
		Currency:       craftgate.TRY,
		PaymentGroup:   craftgate.LISTING_OR_SUBSCRIPTION,
		ConversationId: "foo-bar",
		ExternalId:     "115",
		Card: &craftgate.Card{
			CardHolderName: "Card Holder",
			CardNumber:     "4256690000000001",
			ExpireYear:     "2035",
			ExpireMonth:    "11",
			Cvc:            "123",
		},
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

	res, err := paymentClient.Payment.CreatePayment(context.Background(), request)

	require.NotNil(t, res.Id)
	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_CreateApmPayment(t *testing.T) {
	request := adapter.CreateApmPaymentRequest{
		ApmType:        craftgate.ApmTypeCASH_ON_DELIVERY,
		Price:          1.25,
		PaidPrice:      1.25,
		Currency:       craftgate.TRY,
		PaymentGroup:   craftgate.LISTING_OR_SUBSCRIPTION,
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
	}

	res, err := paymentClient.Payment.CreateApmPayment(context.Background(), request)

	require.NotNil(t, res.Id)
	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RetrievePayment(t *testing.T) {
	paymentId := int64(1)
	res, err := paymentClient.Payment.RetrievePayment(context.Background(), paymentId)

	require.Equal(t, *res.Id, paymentId)
	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_Init3DSPayment(t *testing.T) {
	request := adapter.Init3DSPaymentRequest{
		Price:          1.25,
		PaidPrice:      1.25,
		Installment:    1,
		Currency:       craftgate.TRY,
		PaymentGroup:   craftgate.LISTING_OR_SUBSCRIPTION,
		ConversationId: "foo-bar",
		ExternalId:     "115",
		Card: &craftgate.Card{
			CardHolderName: "Card Holder",
			CardNumber:     "4256690000000001",
			ExpireYear:     "2035",
			ExpireMonth:    "11",
			Cvc:            "123",
		},
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
		CallbackUrl: "https://www.your-website.com/callback",
	}

	res, err := paymentClient.Payment.Init3DSPayment(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	require.NotEmpty(t, *res.HtmlContent)
	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_Complete3DSPayment(t *testing.T) {
	request := adapter.Complete3DSPaymentRequest{
		PaymentId: 123,
	}
	res, err := paymentClient.Payment.Complete3DSPayment(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_CreatePreAuthPayment(t *testing.T) {
	request := adapter.CreatePaymentRequest{
		Price:          1.25,
		PaidPrice:      1.25,
		Installment:    1,
		Currency:       craftgate.TRY,
		PaymentGroup:   craftgate.LISTING_OR_SUBSCRIPTION,
		PaymentPhase:   craftgate.PRE_AUTH,
		ConversationId: "foo-bar",
		ExternalId:     "115",
		Card: &craftgate.Card{
			CardHolderName: "Card Holder",
			CardNumber:     "4256690000000001",
			ExpireYear:     "2035",
			ExpireMonth:    "11",
			Cvc:            "123",
		},
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

	res, err := paymentClient.Payment.CreatePayment(context.Background(), request)

	require.NotNil(t, res.Id)
	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_PostAuthPayment(t *testing.T) {
	request := adapter.PostAuthPaymentRequest{
		PaidPrice: 1.25,
	}
	res, err := paymentClient.Payment.PostAuthPayment(context.Background(), 123, request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_InitCheckoutPayment(t *testing.T) {
	request := adapter.InitCheckoutPaymentRequest{
		Price:          1.25,
		PaidPrice:      1.25,
		Currency:       craftgate.TRY,
		PaymentGroup:   craftgate.LISTING_OR_SUBSCRIPTION,
		PaymentPhase:   craftgate.AUTH,
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
		CallbackUrl: "https://www.your-website.com/callback",
	}
	res, err := paymentClient.Payment.InitCheckoutPayment(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RetrieveCheckoutPayment(t *testing.T) {
	res, err := paymentClient.Payment.RetrieveCheckoutPayment(context.Background(), "foo-bar")
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_ExpireCheckoutPayment(t *testing.T) {
	err := paymentClient.Payment.ExpireCheckoutPayment(context.Background(), "foo-bar")

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_CreateDepositPayment(t *testing.T) {
	request := adapter.DepositPaymentRequest{
		Price:          100,
		BuyerMemberId:  1,
		ConversationId: "456d1297-908e-4bd6-a13b-4be31a6e47d5",
		Card: craftgate.Card{
			CardHolderName: "Haluk Demir",
			CardNumber:     "5258640000000001",
			ExpireYear:     "2044",
			ExpireMonth:    "07",
			Cvc:            "000",
		},
	}
	res, err := paymentClient.Payment.CreateDepositPayment(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_Init3DSDepositPayment(t *testing.T) {
	request := adapter.DepositPaymentRequest{
		Price:          100,
		BuyerMemberId:  1,
		CallbackUrl:    "https://www.your-website.com/craftgate-3DSecure-callback",
		ConversationId: "456d1297-908e-4bd6-a13b-4be31a6e47d5",
		Card: craftgate.Card{
			CardHolderName: "Haluk Demir",
			CardNumber:     "5258640000000001",
			ExpireYear:     "2044",
			ExpireMonth:    "07",
			Cvc:            "000",
		},
	}
	res, err := paymentClient.Payment.Init3DSDepositPayment(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_Complete3DSDepositPayment(t *testing.T) {
	request := adapter.Complete3DSPaymentRequest{
		PaymentId: 1,
	}
	res, err := paymentClient.Payment.Complete3DSDepositPayment(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_CreateFundTransferDepositPayment(t *testing.T) {
	request := adapter.CreateFundTransferDepositPaymentRequest{
		Price:          100,
		BuyerMemberId:  1,
		ConversationId: "456d1297-908e-4bd6-a13b-4be31a6e47d5",
	}

	res, err := paymentClient.Payment.CreateFundTransferDepositPayment(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_InitDepositApmPayment(t *testing.T) {
	request := adapter.InitApmDepositPaymentRequest{
		ApmType:        craftgate.ApmTypePAPARA,
		Price:          1.25,
		Currency:       craftgate.TRY,
		BuyerMemberId:  1,
		ConversationId: "foo-bar",
		CallbackUrl:    "https://www.your-website.com/callback",
		ClientIp:       "127.0.0.1",
	}
	res, err := paymentClient.Payment.InitApmDepositPayment(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_InitGarantiPayPayment(t *testing.T) {
	request := adapter.InitGarantiPayPaymentRequest{
		Price:          100,
		PaidPrice:      100,
		Currency:       craftgate.TRY,
		PaymentGroup:   craftgate.LISTING_OR_SUBSCRIPTION,
		ConversationId: "456d1297-908e-4bd6-a13b-4be31a6e47d5",
		CallbackUrl:    "https://www.your-website.com/craftgate-garantipay-callback",
		Items: []craftgate.PaymentItem{
			{
				Name:       "Item 1",
				Price:      60,
				ExternalId: "1",
			},
			{
				Name:       "Item 2",
				Price:      40,
				ExternalId: "2",
			},
		},
		Installments: []adapter.GarantiPayInstallment{
			{
				Number:     2,
				TotalPrice: 120,
			},
			{
				Number:     3,
				TotalPrice: 125,
			},
		},
	}
	res, err := paymentClient.Payment.InitGarantiPayPayment(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_InitApmPayment(t *testing.T) {
	request := adapter.InitApmPaymentRequest{
		ApmType:         craftgate.ApmTypeEDENRED,
		Price:           1.25,
		PaidPrice:       1.25,
		Currency:        craftgate.TRY,
		PaymentGroup:    craftgate.LISTING_OR_SUBSCRIPTION,
		ConversationId:  "foo-bar",
		ApmUserIdentity: "4242424242424242",
		CallbackUrl:     "https://www.your-website.com/callback",
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
	res, err := paymentClient.Payment.InitApmPayment(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_InitKlarnaApmPayment(t *testing.T) {
	request := adapter.InitApmPaymentRequest{
		ApmType:        craftgate.ApmTypeKLARNA,
		Price:          1,
		PaidPrice:      1,
		Currency:       craftgate.USD,
		PaymentGroup:   craftgate.LISTING_OR_SUBSCRIPTION,
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
		AdditionalParams: map[string]string{
			"country": "de",
			"locale":  "en-DE",
		},
	}
	res, err := paymentClient.Payment.InitApmPayment(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_InitAfterpayApmPayment(t *testing.T) {
	request := adapter.InitApmPaymentRequest{
		ApmType:        craftgate.ApmTypeAFTERPAY,
		Price:          1,
		PaidPrice:      1,
		Currency:       craftgate.USD,
		PaymentGroup:   craftgate.LISTING_OR_SUBSCRIPTION,
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
	}
	res, err := paymentClient.Payment.InitApmPayment(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_CompleteApmPayment(t *testing.T) {
	request := adapter.CompleteApmPaymentRequest{
		PaymentId:        123,
		AdditionalParams: map[string]string{"otpCode": "123456"},
	}
	res, err := paymentClient.Payment.CompleteApmPayment(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RetrieveLoyalties(t *testing.T) {
	request := adapter.RetrieveLoyaltiesRequest{
		CardNumber:  "4043080000000003",
		ExpireYear:  "2044",
		ExpireMonth: "07",
		Cvc:         "000",
	}
	res, err := paymentClient.Payment.RetrieveLoyalties(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RefundPaymentTransaction(t *testing.T) {
	request := adapter.RefundPaymentTransactionRequest{
		PaymentTransactionId:  1,
		ConversationId:        "456d1297-908e-4bd6-a13b-4be31a6e47d5",
		RefundPrice:           20,
		RefundDestinationType: craftgate.RefundDestinationTypePROVIDER,
	}

	res, err := paymentClient.Payment.RefundPaymentTransaction(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RetrievePaymentTransactionRefund(t *testing.T) {
	res, err := paymentClient.Payment.RetrievePaymentTransactionRefund(context.Background(), 123)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RefundPayment(t *testing.T) {
	request := adapter.RefundPaymentRequest{
		PaymentId:             1,
		RefundDestinationType: craftgate.RefundDestinationTypePROVIDER,
	}
	res, err := paymentClient.Payment.RefundPayment(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RetrievePaymentRefund(t *testing.T) {
	res, err := paymentClient.Payment.RetrievePaymentRefund(context.Background(), 123)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_StoreCard(t *testing.T) {
	request := adapter.StoreCardRequest{
		CardHolderName: "Haluk Demir",
		CardNumber:     "5258640000000001",
		ExpireYear:     "2044",
		ExpireMonth:    "07",
		CardAlias:      "My Card",
	}
	res, err := paymentClient.Payment.StoreCard(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_UpdateStoredCard(t *testing.T) {
	request := adapter.UpdateStoredCardRequest{
		CardUserKey: "6bcbac4b-6460-418d-b060-2d9896c08156",
		CardToken:   "aa57f470-7423-449e-87b7-afb1fba151fb",
		ExpireYear:  "2050",
		ExpireMonth: "12",
	}
	res, err := paymentClient.Payment.UpdateStoredCard(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_DeleteStoredCard(t *testing.T) {
	request := adapter.DeleteStoredCardRequest{
		CardUserKey: "d94018bb-baa9-4418-84f8-760942f669af",
		CardToken:   "76efcbfd-6bef-4b85-bc85-f0beb0baf664",
	}
	err := paymentClient.Payment.DeleteStoredCard(context.Background(), request)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_SearchStoredCards(t *testing.T) {
	request := adapter.SearchStoredCardsRequest{
		CardAlias:       "My YKB Card",
		CardBankName:    "YAPI VE KREDI BANKASI A.S.",
		CardBrand:       "World",
		CardType:        craftgate.CREDIT_CARD,
		CardAssociation: craftgate.MASTER_CARD,
	}

	res, err := paymentClient.Payment.SearchStoredCards(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_UpdatePaymentTransaction(t *testing.T) {
	request := adapter.UpdatePaymentTransactionRequest{
		SubMerchantMemberId:    1,
		SubMerchantMemberPrice: 0.60,
	}
	res, err := paymentClient.Payment.UpdatePaymentTransaction(context.Background(), 123, request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_ApprovePaymentTransactions(t *testing.T) {
	request := adapter.PaymentTransactionsApprovalRequest{
		PaymentTransactionIds: []int64{1, 2},
		IsTransactional:       true,
	}
	res, err := paymentClient.Payment.ApprovePaymentTransactions(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_DisapprovePaymentTransactions(t *testing.T) {
	request := adapter.PaymentTransactionsApprovalRequest{
		PaymentTransactionIds: []int64{1, 2},
		IsTransactional:       true,
	}
	res, err := paymentClient.Payment.DisapprovePaymentTransactions(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_CheckMasterpassUser(t *testing.T) {
	request := adapter.CheckMasterpassUserRequest{
		MasterpassGsmNumber: "903000000000",
	}
	res, err := paymentClient.Payment.CheckMasterpassUser(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_Verify3DSCallback(t *testing.T) {
	merchantThreeDsCallbackKey := "merchantThreeDsCallbackKeySndbox"

	params := make(map[string]string)
	params["hash"] = "1d3fa1e51fe7c350185c5a7f8c3ff513a991367b08c16a56f4ab9abeb738a1e1"
	params["paymentId"] = "5"
	params["conversationData"] = "conversation-data"
	params["conversationId"] = "conversation-id"
	params["status"] = "SUCCESS"
	params["completeStatus"] = "WAITING"

	is3DSecureCallbackVerified := paymentClient.Payment.Is3DSecureCallbackVerified(merchantThreeDsCallbackKey, params)

	require.True(t, is3DSecureCallbackVerified)
}

func TestPayment_NotVerify3DSCallback(t *testing.T) {
	merchantThreeDsCallbackKey := "merchantThreeDsCallbackKeySndbox"

	params := make(map[string]string)
	params["hash"] = "39427942bcaasjaduqabzhdancaASasdhbcxjancakjscace82"
	params["paymentId"] = "5"
	params["conversationData"] = "conversation-data"
	params["conversationId"] = "conversation-id"
	params["status"] = "SUCCESS"
	params["completeStatus"] = "WAITING"

	is3DSecureCallbackVerified := paymentClient.Payment.Is3DSecureCallbackVerified(merchantThreeDsCallbackKey, params)

	require.False(t, is3DSecureCallbackVerified)
}
