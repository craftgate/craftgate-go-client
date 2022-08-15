package tests

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

var payment = adapter.Payment{
	Opts: model.RequestOptions{
		BaseURL:   "https://sandbox-api.craftgate.io",
		ApiKey:    "api-key",
		SecretKey: "secret-key",
	},
}

func TestPayment_CreatePayment(t *testing.T) {
	request := adapter.CreatePaymentRequest{
		Price:          1.25,
		PaidPrice:      1.25,
		Installment:    1,
		Currency:       model.Currency(model.TRY),
		PaymentGroup:   model.PaymentGroup(model.LISTING_OR_SUBSCRIPTION),
		ConversationId: "foo-bar",
		ExternalId:     "115",
		Card: model.Card{
			CardHolderName: "Card Holder",
			CardNumber:     "4256690000000001",
			ExpireYear:     "2035",
			ExpireMonth:    "11",
			Cvc:            "123",
		},
		Items: []model.PaymentItem{
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

	res, err := payment.CreatePayment(request)
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RetrievePayment(t *testing.T) {
	res, err := payment.RetrievePayment(123)
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_Init3DSPayment(t *testing.T) {
	request := adapter.Init3DSPaymentRequest{
		Price:          1.25,
		PaidPrice:      1.25,
		Installment:    1,
		Currency:       model.Currency(model.TRY),
		PaymentGroup:   model.PaymentGroup(model.LISTING_OR_SUBSCRIPTION),
		ConversationId: "foo-bar",
		ExternalId:     "115",
		Card: model.Card{
			CardHolderName: "Card Holder",
			CardNumber:     "4256690000000001",
			ExpireYear:     "2035",
			ExpireMonth:    "11",
			Cvc:            "123",
		},
		Items: []model.PaymentItem{
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
		CallbackUrl: "www.callback.url",
	}

	res, err := payment.Init3DSPayment(request)
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_Complete3DSPayment(t *testing.T) {
	res, err := payment.Complete3DSPayment(adapter.Complete3DSPaymentRequest{
		PaymentId: 123,
	})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_PostAuthPayment(t *testing.T) {
	res, err := payment.PostAuthPayment(123, adapter.PostAuthPaymentRequest{
		PaidPrice: 15.3,
	})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_InitCheckoutPayment(t *testing.T) {
	request := adapter.InitCheckoutPaymentRequest{
		Price:          1.25,
		PaidPrice:      1.25,
		BuyerMemberId:  1,
		Currency:       model.TRY,
		PaymentGroup:   model.LISTING_OR_SUBSCRIPTION,
		PaymentPhase:   model.AUTH,
		ConversationId: "foo-bar",
		ExternalId:     "115",
		Items: []model.PaymentItem{
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
		CallbackUrl: "www.callback.url",
	}
	res, err := payment.InitCheckoutPayment(request)
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RetrieveCheckoutPayment(t *testing.T) {
	res, err := payment.RetrieveCheckoutPayment("foo-bar")
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_CreateDepositPayment(t *testing.T) {
	res, err := payment.CreateDepositPayment(adapter.DepositPaymentRequest{
		Price:          100,
		BuyerMemberId:  1,
		ConversationId: "456d1297-908e-4bd6-a13b-4be31a6e47d5",
		Card: model.Card{
			CardHolderName: "Haluk Demir",
			CardNumber:     "5258640000000001",
			ExpireYear:     "2044",
			ExpireMonth:    "07",
			Cvc:            "000",
		},
	})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_Init3DSDepositPayment(t *testing.T) {
	res, err := payment.Init3DSDepositPayment(adapter.DepositPaymentRequest{
		Price:          100,
		BuyerMemberId:  1,
		CallbackUrl:    "https://www.your-website.com/craftgate-3DSecure-callback",
		ConversationId: "456d1297-908e-4bd6-a13b-4be31a6e47d5",
		Card: model.Card{
			CardHolderName: "Haluk Demir",
			CardNumber:     "5258640000000001",
			ExpireYear:     "2044",
			ExpireMonth:    "07",
			Cvc:            "000",
		},
	})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_Complete3DSDepositPayment(t *testing.T) {
	res, err := payment.Complete3DSDepositPayment(adapter.Complete3DSPaymentRequest{
		PaymentId: 1,
	})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_CreateFundTransferDepositPayment(t *testing.T) {
	err := payment.CreateFundTransferDepositPayment(adapter.CreateFundTransferDepositPaymentRequest{
		Price:          100,
		BuyerMemberId:  1,
		ConversationId: "456d1297-908e-4bd6-a13b-4be31a6e47d5",
	})

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_InitGarantiPayPayment(t *testing.T) {
	res, err := payment.InitGarantiPayPayment(adapter.InitGarantiPayPaymentRequest{
		Price:          100,
		PaidPrice:      100,
		Currency:       model.TRY,
		PaymentGroup:   model.LISTING_OR_SUBSCRIPTION,
		ConversationId: "456d1297-908e-4bd6-a13b-4be31a6e47d5",
		CallbackUrl:    "https://www.your-website.com/craftgate-garantipay-callback",
		Items: []model.PaymentItem{
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
	})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RetrieveLoyalties(t *testing.T) {
	res, err := payment.RetrieveLoyalties(adapter.RetrieveLoyaltiesRequest{
		CardNumber:  "4043080000000003",
		ExpireYear:  "2044",
		ExpireMonth: "07",
		Cvc:         "000",
	})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RefundPaymentTransaction(t *testing.T) {
	res, err := payment.RefundPaymentTransaction(adapter.RefundPaymentTransactionRequest{
		PaymentTransactionId:  1,
		ConversationId:        "456d1297-908e-4bd6-a13b-4be31a6e47d5",
		RefundPrice:           20,
		RefundDestinationType: model.RefundDestinationTypeCARD,
	})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RetrievePaymentTransactionRefund(t *testing.T) {
	res, err := payment.RetrievePaymentTransactionRefund(123)
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RefundPayment(t *testing.T) {
	res, err := payment.RefundPayment(adapter.RefundPaymentRequest{
		PaymentId:             1,
		RefundDestinationType: model.RefundDestinationTypeCARD,
	})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RetrievePaymentRefund(t *testing.T) {
	res, err := payment.RetrievePaymentRefund(123)
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_StoreCard(t *testing.T) {
	res, err := payment.StoreCard(adapter.StoreCardRequest{
		CardHolderName: "Haluk Demir",
		CardNumber:     "5258640000000001",
		ExpireYear:     "2044",
		ExpireMonth:    "07",
		CardAlias:      "My Card",
	})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_UpdateStoredCard(t *testing.T) {
	res, err := payment.UpdateStoredCard(adapter.UpdateStoredCardRequest{
		CardUserKey: "fac377f2-ab15-4696-88d2-5e71b27ec378",
		CardToken:   "11a078c4-3c32-4796-90b1-51ee5517a212",
		ExpireYear:  "2044",
		ExpireMonth: "07",
	})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_DeleteStoredCard(t *testing.T) {
	err := payment.DeleteStoredCard(adapter.DeleteStoredCardRequest{
		CardUserKey: "fac377f2-ab15-4696-88d2-5e71b27ec378",
		CardToken:   "11a078c4-3c32-4796-90b1-51ee5517a212",
	})

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_SearchStoredCards(t *testing.T) {
	res, err := payment.SearchStoredCards(adapter.SearchStoredCardsRequest{
		CardAlias:       "My YKB Card",
		CardBankName:    "YAPI VE KREDI BANKASI A.S.",
		CardBrand:       "World",
		CardType:        model.CREDIT_CARD,
		CardAssociation: model.MASTER_CARD,
		CardUserKey:     "c115ecdf-0afc-4d83-8a1b-719c2af19cbd",
		CardToken:       "d9b19d1a-243c-43dc-a498-add08162df72",
	})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_ApprovePaymentTransactions(t *testing.T) {
	res, err := payment.ApprovePaymentTransactions(adapter.PaymentTransactionsApprovalRequest{
		PaymentTransactionIds: []int64{1, 2},
		IsTransactional:       true,
	})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_DisapprovePaymentTransactions(t *testing.T) {
	res, err := payment.DisapprovePaymentTransactions(adapter.PaymentTransactionsApprovalRequest{
		PaymentTransactionIds: []int64{1, 2},
		IsTransactional:       true,
	})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_CheckMasterpassUser(t *testing.T) {
	res, err := payment.CheckMasterpassUser(adapter.CheckMasterpassUserRequest{
		MasterpassGsmNumber: "903000000000",
	})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
