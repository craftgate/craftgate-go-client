package tests

import (
	"context"
	"github.com/craftgate/craftgate-go-client/adapter"
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

var settlementClient, _ = craftgate.New("api-key", "secret-key", "https://sandbox-api.craftgate.io", make(map[string]string))

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
		AccountType:       craftgate.PayoutAccountType_WISE,
		ExternalAccountId: "wiseRecipientId",
		Currency:          craftgate.Currency_USD,
		AccountOwner:      craftgate.AccountOwner_MERCHANT,
	}

	res, err := settlementClient.Settlement.CreatePayoutAccount(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlement_CreateSubMerchantPayoutAccount(t *testing.T) {
	request := adapter.CreatePayoutAccountRequest{
		AccountType:         craftgate.PayoutAccountType_WISE,
		ExternalAccountId:   "wiseRecipientId",
		Currency:            craftgate.Currency_EUR,
		AccountOwner:        craftgate.AccountOwner_SUB_MERCHANT_MEMBER,
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
		AccountType:       craftgate.PayoutAccountType_WISE,
		ExternalAccountId: "wiseRecipientId2",
	}

	err := settlementClient.Settlement.UpdatePayoutAccount(context.Background(), 18, request)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlement_DeletePayoutAccount(t *testing.T) {
	err := settlementClient.Settlement.DeletePayoutAccount(context.Background(), 18)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func TestSettlement_SearchPayoutAccounts(t *testing.T) {
	request := adapter.SearchPayoutAccountRequest{
		Currency:     craftgate.Currency_USD,
		AccountOwner: craftgate.AccountOwner_MERCHANT,
	}

	res, err := settlementClient.Settlement.SearchPayoutAccounts(context.Background(), request)
	_, _ = spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
