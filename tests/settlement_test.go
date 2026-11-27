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
