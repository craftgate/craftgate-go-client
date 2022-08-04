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
		ApiKey:    "sandbox-YEhueLgomBjqsnvBlWVVuFsVhlvJlMHE",
		SecretKey: "sandbox-tBdcdKVGmGupzfaWcULcwDLMoglZZvTz",
	},
}

func TestPayment_CreatePayment(t *testing.T) {
	request := adapter.CreatePaymentRequest{
		Price:          1.25,
		PaidPrice:      1.25,
		Installment:    1,
		Currency:       model.Currency(model.TRY),
		PaymentGroup:   model.PaymentGroup(model.LISTING_OR_SUBSCRIPTION),
		ConversationId: "abcdef",
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
		ConversationId: "abcdef",
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
	res, err := payment.Complete3DSPayment(adapter.Complete3DSPaymentRequest{})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_PostAuthPayment(t *testing.T) {
	res, err := payment.PostAuthPayment(123, adapter.PostAuthPaymentRequest{})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_InitCheckoutPayment(t *testing.T) {
	res, err := payment.InitCheckoutPayment(adapter.InitCheckoutPaymentRequest{})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RetrieveCheckoutPayment(t *testing.T) {
	res, err := payment.RetrieveCheckoutPayment("asdfg")
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_CreateDepositPayment(t *testing.T) {
	res, err := payment.CreateDepositPayment(adapter.DepositPaymentRequest{})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_Init3DSDepositPayment(t *testing.T) {
	res, err := payment.Init3DSDepositPayment(adapter.DepositPaymentRequest{})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_Complete3DSDepositPayment(t *testing.T) {
	res, err := payment.Complete3DSDepositPayment(adapter.Complete3DSPaymentRequest{})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_CreateFundTransferDepositPayment(t *testing.T) {
	err := payment.CreateFundTransferDepositPayment(adapter.CreateFundTransferDepositPaymentRequest{})

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_InitGarantiPayPayment(t *testing.T) {
	res, err := payment.InitGarantiPayPayment(adapter.InitGarantiPayPaymentRequest{})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RetrieveLoyalties(t *testing.T) {
	res, err := payment.RetrieveLoyalties(adapter.RetrieveLoyaltiesRequest{})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_RefundPaymentTransaction(t *testing.T) {
	res, err := payment.RefundPaymentTransaction(adapter.RefundPaymentTransactionRequest{})
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
	res, err := payment.RefundPayment(adapter.RefundPaymentRequest{})
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
	res, err := payment.StoreCard(adapter.StoreCardRequest{})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_UpdateStoredCard(t *testing.T) {
	res, err := payment.UpdateStoredCard(adapter.UpdateStoredCardRequest{})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_DeleteStoredCard(t *testing.T) {
	err := payment.DeleteStoredCard(adapter.DeleteStoredCardRequest{})

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_SearchStoredCards(t *testing.T) {
	res, err := payment.SearchStoredCards(adapter.SearchStoredCardsRequest{})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_ApprovePaymentTransactions(t *testing.T) {
	res, err := payment.ApprovePaymentTransactions(adapter.PaymentTransactionsApprovalRequest{})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_DisapprovePaymentTransactions(t *testing.T) {
	res, err := payment.DisapprovePaymentTransactions(adapter.PaymentTransactionsApprovalRequest{})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestPayment_CheckMasterpassUser(t *testing.T) {
	res, err := payment.CheckMasterpassUser(adapter.CheckMasterpassUserRequest{})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
