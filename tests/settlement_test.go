package tests

import (
	"craftgate-go-client/adapter"
	"craftgate-go-client/model"
	"fmt"
	"testing"
)

var settlement = adapter.Settlement{
	Opts: model.RequestOptions{
		BaseURL:   "https://sandbox-api.craftgate.io",
		ApiKey:    "sandbox-YEhueLgomBjqsnvBlWVVuFsVhlvJlMHE",
		SecretKey: "sandbox-tBdcdKVGmGupzfaWcULcwDLMoglZZvTz",
	},
}

func TestSettlement_CreateInstantWalletSettlement(t *testing.T) {
	res, err := settlement.CreateInstantWalletSettlement(adapter.CreateInstantWalletSettlementRequest{ExcludedSubMerchantMemberIds: []int64{1, 2, 3}})
	fmt.Println(res)

	if err != nil {
		t.Errorf("Error %s", err)
	}
}
