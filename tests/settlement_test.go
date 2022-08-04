package tests

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

var settlement = adapter.Settlement{
	Opts: model.RequestOptions{
		BaseURL:   "https://sandbox-api.craftgate.io",
		ApiKey:    "sandbox-bnqfCZGyogzVmQKuiHPvwilBKDAmYvoB",
		SecretKey: "sandbox-okZEjZlBlteIPARYChRHewtPgKgHAoXO",
	},
}

func TestSettlement_CreateInstantWalletSettlement(t *testing.T) {
	res, err := settlement.CreateInstantWalletSettlement(adapter.CreateInstantWalletSettlementRequest{ExcludedSubMerchantMemberIds: []int64{1, 2}})
	spew.Printf("%#v\n", res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
