package tests

import (
	"context"
	"github.com/craftgate/craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

var settlementClient, _ = adapter.New("api-key", "secret-key", "https://sandbox-api.craftgate.io")

func TestSettlement_CreateInstantWalletSettlement(t *testing.T) {
	request := adapter.CreateInstantWalletSettlementRequest{ExcludedSubMerchantMemberIds: []int64{1, 2}}

	res, err := settlementClient.Settlement.CreateInstantWalletSettlement(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlement_CreateMerchantPayoutAccount(t *testing.T) {
	request := adapter.CreatePayoutAccountRequest{
		AccountType:       adapter.PayoutAccountType_WISE,
		ExternalAccountId: "wiseRecipientId",
		Currency:          adapter.Currency_USD,
		AccountOwner:      adapter.AccountOwner_MERCHANT,
	}

	res, err := settlementClient.Settlement.CreatePayoutAccount(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlement_CreateSubMerchantPayoutAccount(t *testing.T) {
	request := adapter.CreatePayoutAccountRequest{
		AccountType:         adapter.PayoutAccountType_WISE,
		ExternalAccountId:   "wiseRecipientId",
		Currency:            adapter.Currency_EUR,
		AccountOwner:        adapter.AccountOwner_SUB_MERCHANT_MEMBER,
		SubMerchantMemberId: 1,
	}

	res, err := settlementClient.Settlement.CreatePayoutAccount(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlement_UpdatePayoutAccount(t *testing.T) {
	request := adapter.UpdatePayoutAccountRequest{
		AccountType:       adapter.PayoutAccountType_WISE,
		ExternalAccountId: "wiseRecipientId2",
	}

	err := settlementClient.Settlement.UpdatePayoutAccount(context.Background(), 18, request)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlement_DeletePayoutAccount(t *testing.T) {
	err := settlementClient.Settlement.DeletePayoutAccount(context.Background(),
		adapter.DeletePayoutAccountRequest{Id: 18})

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlement_SearchPayoutAccounts(t *testing.T) {
	request := adapter.SearchPayoutAccountRequest{
		Currency:     adapter.Currency_USD,
		AccountOwner: adapter.AccountOwner_MERCHANT,
	}

	res, err := settlementClient.Settlement.SearchPayoutAccounts(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlement_ApprovePaymentTransactions(t *testing.T) {
	request := adapter.PaymentTransactionsApprovalRequest{
		PaymentTransactionIds: []int64{1, 2},
		IsTransactional:       true,
	}
	res, err := settlementClient.Settlement.ApprovePaymentTransactions(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlement_DisapprovePaymentTransactions(t *testing.T) {
	request := adapter.PaymentTransactionsApprovalRequest{
		PaymentTransactionIds: []int64{1, 2},
		IsTransactional:       true,
	}
	res, err := settlementClient.Settlement.DisapprovePaymentTransactions(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlement_UpdatePaymentTransaction(t *testing.T) {
	request := adapter.UpdatePaymentTransactionRequest{
		SubMerchantMemberId:    1,
		SubMerchantMemberPrice: 0.60,
	}
	res, err := settlementClient.Settlement.UpdatePaymentTransaction(context.Background(), 123, request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
